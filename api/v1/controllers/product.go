package controllers

import (
	"backend/api/v1/validations"
	"backend/internal/app/models"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func indexProduct(c *gin.Context) {

	dto, err := validations.IndexProduct(c)
	if err != nil {
		return
	}
	products, err := models.NewMainManager().IndexProduct(dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "خطا در دریافت اطلاعات"})
	}
	c.JSON(http.StatusOK, products)
	return
}

func createProduct(c *gin.Context) {
	userID := uint64(jwt.ExtractClaims(c)["id"].(float64))

	dto, err := validations.CreateProduct(c)
	if err != nil {
		return
	}

	images, err := uploadImage(c, dto.Images, userID)
	if err != nil {
		return
	}

	dto.ImagePath = images
	err = models.NewMainManager().CreateProduct(dto, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "خطا در ایجاد محصول"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "محصول با موفقیت ایجاد شد",
	})
	return
}

func uploadImage(c *gin.Context, images []*multipart.FileHeader, userID uint64) ([]string, error) {
	abs, _ := filepath.Abs("../../assets")
	var imagesPath []string
	userDir := filepath.Join(abs, "user_"+strconv.FormatUint(userID, 10), "images")
	if err := os.MkdirAll(userDir, os.ModePerm); err != nil {
		log.Fatal(err)
	}
	for i := range images {
		path := filepath.Join(userDir, uuid.NewString()+"."+strings.Split(images[i].Header["Content-Type"][0], "/")[1])
		err := c.SaveUploadedFile(images[i], path)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "خطا در آپلود تصایر"})
			return nil, err
		}
		imagesPath = append(imagesPath, path)
	}
	return imagesPath, nil
}

func updateProduct(c *gin.Context) {

}

func deleteProduct(c *gin.Context) {

}

func showProduct(c *gin.Context) {

}
