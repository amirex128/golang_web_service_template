package controllers

import (
	"backend/internal/app/models"
	"backend/internal/app/validations"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

func CreateDomain(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "addDomain", "request")
	defer span.End()
	dto, err := validations.CreateDomain(c)
	if err != nil {
		return
	}
	err = models.NewMainManager().CreateDomain(c, ctx, dto)
	if err != nil {
		return
	}
	// add server domain to nginx config file and reload nginx

	c.JSON(http.StatusOK, gin.H{
		"message": "دامنه با موفقیت اضافه شد",
	})
}
