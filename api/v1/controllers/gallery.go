package controllers

import (
	"backend/api/v1/validations"
	"backend/internal/app/DTOs"
	"backend/internal/app/models"
	"backend/internal/app/utils"
	"errors"
	"github.com/chai2010/webp"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"image"
	"image/jpeg"
	"image/png"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func createGallery(c *gin.Context) {
	dto, err := validations.CreateGallery(c)
	if err != nil {
		return
	}
	userID := utils.GetUser(c)

	galleryAddress, userDir, err := createDirectory(c, userID)
	if err != nil {
		return
	}
	relativePath, info, err := uploadImage(c, userDir, galleryAddress, dto)
	if err != nil {
		return
	}
	gallery := &models.Gallery{
		UserID:    userID,
		OwnerID:   dto.OwnerID,
		OwnerType: dto.OwnerType,
		Size:      float64(info.Size() / 1024),
		Path:      relativePath,
		Width:     dto.Width,
		Height:    dto.Height,
	}
	_, err = models.NewMainManager().UploadImage(c, gallery)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "تصویر با موفقیت آپلود شد",
		"gallery": gallery,
	})
}

func createDirectory(c *gin.Context, userID uint64) (string, string, error) {
	abs, _ := filepath.Abs("../../assets")
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
	relativePath := filepath.Join("/assets", galleryAddress, imageName)

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
