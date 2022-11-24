package controllers

import (
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
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
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:createDomain", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	dto, err := validations.CreateDomain(c)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	err = models.NewMysqlManager(c).CreateDomain(dto)
	if err != nil {
		errorx.ResponseErrorx(c, err)
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
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:deleteDomain", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	domainID := utils.StringToUint64(c.Param("id"))
	err := models.NewMysqlManager(c).DeleteDomain(domainID)
	if err != nil {
		errorx.ResponseErrorx(c, err)
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
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:indexDomain", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	dto, err := validations.IndexDomain(c)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	domains, err := models.NewMysqlManager(c).GetAllDomainWithPagination(dto)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"domains": domains,
	})
}

// ShowDomain
// @Summary نمایش دامنه
// @description هر فروشگاه برای خود میتواند به تعداد دلخواه دامنه در موقعیت های مختلف مثل بالای صفحه و پایین صفحه ایجاد نماید
// @Tags post
// @Router       /user/domain/show/{id} [get]
// @Param	Authorization	 header string	true "Authentication"
// @Param	id			 path   string	true "شناسه " SchemaExample(1)
func ShowDomain(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:showDomain", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	domainID := c.Param("id")
	domain, err := models.NewMysqlManager(c).FindDomainByID(utils.StringToUint64(domainID))
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"domain": domain,
	})
}
