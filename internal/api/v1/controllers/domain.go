package controllers

import (
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/amirex128/selloora_backend/internal/validations"
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
	err = models.NewMysqlManager(ctx).CreateDomain(c, ctx, dto)
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
	err := models.NewMysqlManager(ctx).DeleteDomain(c, ctx, domainID)
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
	domains, err := models.NewMysqlManager(ctx).GetAllDomainWithPagination(c, ctx, dto)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"domains": domains,
	})
}
