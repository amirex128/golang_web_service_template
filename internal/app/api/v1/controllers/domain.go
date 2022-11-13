package controllers

import (
	"backend/internal/app/models"
	"backend/internal/app/utils"
	"backend/internal/app/validations"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

func CreateDomain(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "createDomain", "request")
	defer span.End()
	dto, err := validations.CreateDomain(c)
	if err != nil {
		return
	}
	err = models.NewMainManager().CreateDomain(c, ctx, dto)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "دامنه با موفقیت ایجاد شد",
	})
}

func DeleteDomain(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "deleteDomain", "request")
	defer span.End()
	domainID := utils.StringToUint64(c.Param("id"))
	err := models.NewMainManager().DeleteDomain(c, ctx, domainID)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "دامنه با موفقیت حذف شد",
	})
}
func IndexDomain(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "indexDomain", "request")
	defer span.End()
	dto, err := validations.IndexDomain(c)
	if err != nil {
		return
	}
	domains, err := models.NewMainManager().GetAllDomainWithPagination(c, ctx, dto)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"domains": domains,
	})
}
