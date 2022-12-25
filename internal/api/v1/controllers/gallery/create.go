package gallery

import (
	"github.com/amirex128/selloora_backend/internal/DTOs"
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
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
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:createGallery", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	dto, err := validations.CreateGallery(c)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	userID := models.GetUserID(c)

	galleryAddress, userDir, err := createDirectory(c, *userID)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	relativePath, info, err := uploadImage(userDir, galleryAddress, dto)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	size := float64(info.Size() / 1024)
	if size > 1024 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "image size is too large", "message": "حجم تصویر بیشتر از 1 مگابایت است"})
	}
	gallery := &models.Gallery{
		UserID:   userID,
		Size:     size,
		Path:     relativePath,
		Width:    dto.Width,
		Height:   dto.Height,
		MimeType: "image/webp",
	}
	gallery, err = models.NewMysqlManager(c).UploadImage(gallery)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "تصویر با موفقیت آپلود شد",
		"data":    gallery,
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

func uploadImage(userDir string, galleryAddress string, dto DTOs.CreateGallery) (string, os.FileInfo, error) {
	contentType, ok := dto.File.Header["Content-Type"]
	if !ok {
		return "", nil, errorx.New("خطا در آپلود تصویر", "request", nil)
	}

	imageType := strings.Split(contentType[0], "/")[1]
	imageName := uuid.NewString() + "." + "webp"
	fullPath := filepath.Join(userDir, imageName)
	relativePath := filepath.Join("/public", galleryAddress, imageName)

	open, err := dto.File.Open()
	if err != nil {
		return "", nil, errorx.New("خطا در آپلود تصایر", "request", err)
	}

	file, err := os.Create(fullPath)
	defer file.Close()
	if err != nil {
		return "", nil, errorx.New("خطا در آپلود تصایر", "request", err)
	}

	var imgDecode image.Image
	if imageType == "png" {
		imgDecode, err = png.Decode(open)
		if err != nil {
			return "", nil, errorx.New("فرمت تصویر باید png یا jpg یا jpeg باشد", "request", nil)
		}
	} else if imageType == "jpeg" || imageType == "jpg" {
		imgDecode, err = jpeg.Decode(open)
		if err != nil {
			return "", nil, errorx.New("فرمت تصویر باید png یا jpg یا jpeg باشد", "request", nil)
		}
	} else {
		return "", nil, errorx.New("فرمت پشتیبانی نمیشود", "request", nil)
	}

	err = webp.Encode(file, imgDecode, &webp.Options{Lossless: true})
	if err != nil {
		return "", nil, errorx.New("خطا در آپلود تصایر", "request", nil)
	}

	stat, err := file.Stat()
	if err != nil {
		return "", nil, errorx.New("خطا در آپلود تصایر", "request", nil)
	}

	return relativePath, stat, nil
}
