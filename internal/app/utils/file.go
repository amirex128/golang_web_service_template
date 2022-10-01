package utils

import (
	"encoding/csv"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func ReadCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}
func UploadMultiImage(c *gin.Context, images []*multipart.FileHeader, dest string) ([]string, error) {
	abs, _ := filepath.Abs("../../assets")
	var imagesPath []string
	userDir := filepath.Join(abs, dest, "images")
	if err := os.MkdirAll(userDir, os.ModePerm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "خطا در آپلود تصویر"})
		return nil, err
	}
	for i := range images {
		relativePath := uuid.NewString() + "." + strings.Split(images[i].Header["Content-Type"][0], "/")[1]
		path := filepath.Join(userDir, relativePath)
		relativePath = filepath.Join("/assets", dest, "images", relativePath)
		err := c.SaveUploadedFile(images[i], path)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "خطا در آپلود تصایر"})
			return nil, err
		}
		imagesPath = append(imagesPath, relativePath)
	}
	return imagesPath, nil
}
func UploadImage(c *gin.Context, image *multipart.FileHeader, dest string) (string, error) {
	abs, _ := filepath.Abs("../../assets")
	userDir := filepath.Join(abs, dest, "images")
	if err := os.MkdirAll(userDir, os.ModePerm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "خطا در آپلود تصویر"})
		return "", err
	}
	relativePath := uuid.NewString() + "." + strings.Split(image.Header["Content-Type"][0], "/")[1]
	path := filepath.Join(userDir, relativePath)
	relativePath = filepath.Join("/assets", dest, "images", relativePath)
	err := c.SaveUploadedFile(image, path)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "خطا در آپلود تصایر"})
		return "", err
	}

	return relativePath, nil
}
func RemoveImages(images []string) {
	for i := range images {
		err := os.Remove(images[i])
		if err != nil {
			continue
		}
	}
}
