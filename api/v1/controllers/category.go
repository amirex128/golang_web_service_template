package controllers

import (
	"backend/api/v1/validations"
	"backend/internal/app/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func indexCategory(c *gin.Context) {
	dto, err := validations.IndexCategory(c)
	if err != nil {
		return
	}
	categories, err := models.NewMainManager().GetAllCategoryWithPagination(c, dto)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"categories": categories,
	})

}
