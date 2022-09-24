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
	"strconv"
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
func UploadImages(c *gin.Context, images []*multipart.FileHeader, userID uint64) ([]string, error) {
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
func RemoveImages(images []string) {
	for i := range images {
		err := os.Remove(images[i])
		if err != nil {
			continue
		}
	}
}
