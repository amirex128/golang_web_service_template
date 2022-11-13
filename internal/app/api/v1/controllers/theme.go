package controllers

import (
	"backend/internal/app/models"
	"backend/internal/app/validations"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

func IndexTheme(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "indexTheme", "request")
	defer span.End()
	dto, err := validations.IndexTheme(c)
	if err != nil {
		return
	}
	pages, err := models.NewMainManager().GetAllThemeWithPagination(c, ctx, dto)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"pages": pages,
	})
}
