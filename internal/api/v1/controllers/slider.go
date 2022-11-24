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

// CreateSlider
// @Summary ایجاد اسلایدر
// @description هر فروشگاه برای خود میتواند به تعداد دلخواه اسلایدر در موقعیت های مختلف مثل بالای صفحه و پایین صفحه ایجاد نماید
// @Tags slider
// @Router       /user/slider/create [post]
// @Param	Authorization	 header string	true "Authentication"
// @Param	message	 body   DTOs.CreateSlider  	true "ورودی"
func CreateSlider(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:createSlider", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	dto, err := validations.CreateSlider(c)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	slider, err := models.NewMysqlManager(c).CreateSlider(dto)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "اسلایدر با موفقیت ایجاد شد",
		"data":    slider,
	})
}

// UpdateSlider
// @Summary ویرایش اسلایدر
// @description هر فروشگاه برای خود میتواند به تعداد دلخواه اسلایدر در موقعیت های مختلف مثل بالای صفحه و پایین صفحه ایجاد نماید
// @Tags slider
// @Router       /user/slider/update/{id} [post]
// @Param	Authorization	 header string	true "Authentication"
// @Param	id			 path   string	false "شناسه " SchemaExample(1)
// @Param	message	 body   DTOs.UpdateSlider  	true "ورودی"
func UpdateSlider(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:updateSlider", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	dto, err := validations.UpdateSlider(c)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	err = models.NewMysqlManager(c).UpdateSlider(dto)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "اسلایدر با موفقیت ویرایش شد",
	})
}

// DeleteSlider
// @Summary حذف اسلایدر
// @description هر فروشگاه برای خود میتواند به تعداد دلخواه اسلایدر در موقعیت های مختلف مثل بالای صفحه و پایین صفحه ایجاد نماید
// @Tags slider
// @Router       /user/slider/delete/{id} [post]
// @Param	Authorization	 header string	true "Authentication"
// @Param	id			 path   string	true "شناسه اسلایدر" SchemaExample(1)
func DeleteSlider(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:deleteSlider", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	sliderID := utils.StringToUint64(c.Param("id"))
	err := models.NewMysqlManager(c).DeleteSlider(sliderID)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "اسلایدر با موفقیت حذف شد",
	})
}

// IndexSlider
// @Summary لیست اسلایدر ها
// @description هر فروشگاه برای خود میتواند به تعداد دلخواه اسلایدر در موقعیت های مختلف مثل بالای صفحه و پایین صفحه ایجاد نماید
// @Tags slider
// @Router       /user/slider [get]
// @Param	Authorization	 header string	true "Authentication"
// @Param	search			 query   string	false "متن جستجو"
// @Param	page			 query   string	false "شماره صفحه"
// @Param	page_size		 query   string	false "تعداد صفحه"
// @Param	sort			 query   string	false "مرتب سازی براساس desc/asc"
func IndexSlider(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:indexSlider", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	dto, err := validations.IndexSlider(c)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	sliders, err := models.NewMysqlManager(c).GetAllSliderWithPagination(dto)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"sliders": sliders,
	})
}

// ShowSlider
// @Summary نمایش اسلایدر
// @description هر فروشگاه برای خود میتواند به تعداد دلخواه اسلایدر در موقعیت های مختلف مثل بالای صفحه و پایین صفحه ایجاد نماید
// @Tags post
// @Router       /user/slider/show/{id} [get]
// @Param	Authorization	 header string	true "Authentication"
// @Param	id			 path   string	true "شناسه " SchemaExample(1)
func ShowSlider(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:showSlider", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	sliderID := c.Param("id")
	slider, err := models.NewMysqlManager(c).FindSliderByID(utils.StringToUint64(sliderID))
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"slider": slider,
	})
}
