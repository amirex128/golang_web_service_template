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

// CreateAddress
// @Summary ایجاد آدرس جدید
// @description کاربران میتوانند برای خود لیستی از ادرس های مختلف ایجاد کنند تا هر بار به راحتی مشخص کنند محصول خود را میخواند از کدام ادرس ارسال نمایید
// @Tags address
// @Router       /user/address/create [post]
// @Param	Authorization	 header string	false "Authentication"
// @Param message body DTOs.CreateAddress true "ورودی"
func CreateAddress(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:createAddress", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	dto, err := validations.CreateAddress(c)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	err = models.NewMysqlManager(c).CreateAddress(dto)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "آدرس با موفقیت ایجاد شد",
	})
}

// UpdateAddress
// @Summary ویرایش آدرس
// @description کاربران میتوانند برای خود لیستی از ادرس های مختلف ایجاد کنند تا هر بار به راحتی مشخص کنند محصول خود را میخواند از کدام ادرس ارسال نمایید
// @Tags address
// @Router       /user/address/update [post]
// @Param	Authorization	 header string	true "Authentication"
// @Param message body DTOs.UpdateAddress true "ورودی"
func UpdateAddress(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:updateAddress", "request")
	c.Request.WithContext(ctx)
	defer span.End()

	dto, err := validations.UpdateAddress(c)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	err = models.NewMysqlManager(c).UpdateAddress(dto)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "آدرس با موفقیت ویرایش شد",
	})
}

// DeleteAddress
// @Summary حذف آدرس
// @description کاربران میتوانند برای خود لیستی از ادرس های مختلف ایجاد کنند تا هر بار به راحتی مشخص کنند محصول خود را میخواند از کدام ادرس ارسال نمایید
// @Tags address
// @Router       /user/address/delete/{id} [post]
// @Param	Authorization	 header string	true "Authentication"
// @Param	id			 path   string	true "شناسه آدرس" SchemaExample(1)
func DeleteAddress(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:deleteAddress", "request")
	c.Request.WithContext(ctx)
	defer span.End()

	addressID := utils.StringToUint64(c.Param("id"))
	err := models.NewMysqlManager(c).DeleteAddress(addressID)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "آدرس با موفقیت حذف شد",
	})
}

// IndexAddress
// @Summary لیست آدرس
// @description کاربران میتوانند برای خود لیستی از ادرس های مختلف ایجاد کنند تا هر بار به راحتی مشخص کنند محصول خود را میخواند از کدام ادرس ارسال نمایید
// @Tags address
// @Router       /user/address/list [get]
// @Param	Authorization	 header string	true "Authentication"
// @Param	search			 query   string	false "متن جستجو"
// @Param	page			 query   string	false "شماره صفحه"
// @Param	page_size		 query   string	false "تعداد صفحه"
// @Param	sort			 query   string	false "مرتب سازی براساس desc/asc"
func IndexAddress(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:indexAddress", "request")
	c.Request.WithContext(ctx)
	defer span.End()

	dto, err := validations.IndexAddress(c)
	addresses, err := models.NewMysqlManager(c).GetAllAddressWithPagination(dto)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"addresses": addresses,
	})
}
