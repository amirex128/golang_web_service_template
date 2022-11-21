package controllers

import (
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/amirex128/selloora_backend/internal/validations"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

// CreateDomain
// @Summary ایجاد دامنه
// @description یک دامنه یا یک ساب دامنه کاربر میتواند اضافه نمایید تا سایت ایجاد شده خود را بر بستر آن دامنه مشاهده نماید
// @Tags domain
// @Router       /user/domain/create [post]
// @Param	Authorization	 header string	true "Authentication"
// @Param	message	 body   DTOs.CreateDomain  	true "ورودی"
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

// DeleteDomain
// @Summary حذف دامنه
// @description باحذف دامنه امکان دسترسی از این دامنه بر روی سایت کاربر گرفته میشود
// @Tags domain
// @Router       /user/domain/delete/{id} [post]
// @Param	Authorization	 header string	true "Authentication"
// @Param	id			 path   string	true "شناسه دامنه" SchemaExample(1)
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

// IndexDomain
// @Summary لیست دامنه ها
// @description لیست دامنه ها
// @Tags domain
// @Router       /user/domain [get]
// @Param	Authorization	 header string	true "Authentication"
// @Param	search			 query   string	false "متن جستجو"
// @Param	page			 query   string	false "شماره صفحه"
// @Param	page_size		 query   string	false "تعداد صفحه"
// @Param	sort			 query   string	false "مرتب سازی براساس desc/asc"
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
