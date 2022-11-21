package controllers

import (
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils"
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
// @Param	title			 body   string	true "عنوان"
// @Param	province_id		 body   string	true "شناسه استان"
// @Param	city_id			 body   string	true "شناسه شهر"
// @Param	address			 body   string	true "ادرس کامل"
// @Param	postal_code		 body   string	true "کد پستی"
// @Param	mobile			 body   string	true "موبایل"
// @Param	full_name		 body   string	true "نام گیرنده"
// @Param	lat	        	 body   string	false "عرض جغذافیایی"
// @Param	long			 body   string	false "طول جقرافیایی"
func CreateAddress(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "createAddress", "request")
	defer span.End()
	dto, err := validations.CreateAddress(c)
	if err != nil {
		return
	}
	_, err = models.NewMysqlManager(ctx).CreateAddress(c, ctx, dto)
	if err != nil {
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
// @Param	id			 body   string	true "شناسه آدرس"
// @Param	title		 body   string	false "عنوان"
// @Param	province_id	 body   string	false "شناسه استان"
// @Param	city_id		 body   string	false "شناسه شهر"
// @Param	address		 body   string	false "ادرس کامل"
// @Param	postal_code	 body   string	false "کد پستی"
// @Param	mobile		 body   string	false "موبایل"
// @Param	full_name	 body   string	false "نام گیرنده"
// @Param	lat	         body   string	false "عرض جغذافیایی"
// @Param	long		 body   string	false "طول جقرافیایی"
func UpdateAddress(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "updateAddress", "request")
	defer span.End()
	dto, err := validations.UpdateAddress(c)
	if err != nil {
		return
	}
	err = models.NewMysqlManager(ctx).UpdateAddress(c, ctx, dto)
	if err != nil {
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
// @Param	id			 path   string	true "شناسه آدرس"
func DeleteAddress(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "deleteAddress", "request")
	defer span.End()

	addressID := utils.StringToUint64(c.Param("id"))
	err := models.NewMysqlManager(ctx).DeleteAddress(c, ctx, addressID)
	if err != nil {
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
	span, ctx := apm.StartSpan(c.Request.Context(), "indexAddress", "request")
	defer span.End()

	dto, err := validations.IndexAddress(c)
	addresses, err := models.NewMysqlManager(ctx).GetAllAddressWithPagination(c, ctx, dto)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"addresses": addresses,
	})
}
