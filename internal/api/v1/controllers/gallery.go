package controllers

import (
	"errors"
	"github.com/amirex128/selloora_backend/internal/DTOs"
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/amirex128/selloora_backend/internal/validations"
	"github.com/chai2010/webp"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.elastic.co/apm/v2"
	"image"
	"image/jpeg"
	"image/png"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// CreateGallery
// @Summary آپلود تصویر
// @description با آپلود یک تصویر میتوانید شناسه آن را در بخش های مختلف استفاده نمایید و در آینده بر اساس همین شناسه تصویر را حذف نمایید همچنین تمامی تصاویر به فرمت وب پی تبدیل میشوند
// @Tags gallery
// @Router       /user/gallery/create [post]
// @Param	Authorization	 header string	true "Authentication"
// @Param	message	 body   DTOs.CreateGallery  	true "ورودی"
func CreateGallery(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "createGallery", "request")
	defer span.End()
	dto, err := validations.CreateGallery(c)
	if err != nil {
		return
	}
	userID := models.GetUser(c)

	galleryAddress, userDir, err := createDirectory(c, *userID)
	if err != nil {
		return
	}
	relativePath, info, err := uploadImage(c, userDir, galleryAddress, dto)
	if err != nil {
		return
	}
	size := float64(info.Size() / 1024)
	if size > 1024 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "image size is too large", "message": "حجم تصویر بیشتر از 1 مگابایت است"})
	}
	gallery := &models.Gallery{
		UserID:   *userID,
		Size:     size,
		Path:     relativePath,
		Width:    dto.Width,
		Height:   dto.Height,
		MimeType: "image/webp",
	}
	_, err = models.NewMysqlManager(ctx).UploadImage(c, ctx, gallery)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "تصویر با موفقیت آپلود شد",
		"gallery": gallery,
	})
}

// DeleteGallery
// @Summary حذف گالری
// @description حذف یک تصویر از گالری
// @Tags gallery
// @Router       /user/gallery/delete/{id} [post]
// @Param	Authorization	 header string	true "Authentication"
// @Param	id			 path   string	true "شناسه گالری" SchemaExample(1)
func DeleteGallery(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "deleteGallery", "request")
	defer span.End()
	galleryID := c.Param("id")
	gallery, err := models.NewMysqlManager(ctx).FindGalleryByID(c, ctx, utils.StringToUint64(galleryID))
	if err != nil {
		return
	}
	abs, _ := filepath.Abs("./")
	path := filepath.Join(abs, gallery.Path)
	if err := os.Remove(path); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "خطا در حذف تصویر",
			"error":   err.Error(),
		})
		return
	}
	err = models.NewMysqlManager(ctx).DeleteGallery(c, ctx, utils.StringToUint64(galleryID))
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "تصویر با موفقیت حذف شد",
	})
}

func createDirectory(c *gin.Context, userID uint64) (string, string, error) {
	abs, _ := filepath.Abs("./public")
	galleryAddress := "gallery/user_" + utils.Uint64ToString(userID)
	userDir := filepath.Join(abs, galleryAddress)
	if err := os.MkdirAll(userDir, os.ModePerm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "خطا در آپلود تصویر",
			"error":   err.Error(),
		})
		return "", "", err
	}
	return galleryAddress, userDir, nil
}

func uploadImage(c *gin.Context, userDir string, galleryAddress string, dto DTOs.CreateGallery) (string, os.FileInfo, error) {
	contentType, ok := dto.File.Header["Content-Type"]
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "خطا در آپلود تصویر",
			"error":   "content type not found",
		})
		return "", nil, errors.New("content type not found")
	}

	imageType := strings.Split(contentType[0], "/")[1]
	imageName := uuid.NewString() + "." + "webp"
	fullPath := filepath.Join(userDir, imageName)
	relativePath := filepath.Join("/public", galleryAddress, imageName)

	open, err := dto.File.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "خطا در آپلود تصایر"})
		return "", nil, err
	}

	file, err := os.Create(fullPath)
	defer file.Close()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "خطا در آپلود تصایر"})
		return "", nil, err
	}

	var imgDecode image.Image
	if imageType == "png" {
		imgDecode, err = png.Decode(open)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "فرمت تصویر باید png یا jpg یا jpeg باشد"})
			return "", nil, err
		}
	} else if imageType == "jpeg" || imageType == "jpg" {
		imgDecode, err = jpeg.Decode(open)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "فرمت تصویر باید png یا jpg یا jpeg باشد"})
			return "", nil, err
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "image type not supported", "message": "فرمت پشتیانی نمیشود"})
		return "", nil, errors.New("")
	}

	err = webp.Encode(file, imgDecode, &webp.Options{Lossless: true})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "خطا در آپلود تصایر"})
		return "", nil, err
	}

	stat, err := file.Stat()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "خطا در آپلود تصایر"})
		return "", nil, err
	}

	return relativePath, stat, nil
}
