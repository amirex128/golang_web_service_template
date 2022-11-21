package controllers

import (
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/amirex128/selloora_backend/internal/validations"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

func CreateSlider(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "createSlider", "request")
	defer span.End()
	dto, err := validations.CreateSlider(c)
	if err != nil {
		return
	}
	err = models.NewMysqlManager(ctx).CreateSlider(c, ctx, dto)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "اسلایدر با موفقیت ایجاد شد",
	})
}
func UpdateSlider(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "updateSlider", "request")
	defer span.End()
	dto, err := validations.UpdateSlider(c)
	if err != nil {
		return
	}
	err = models.NewMysqlManager(ctx).UpdateSlider(c, ctx, dto)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "اسلایدر با موفقیت ویرایش شد",
	})
}
func DeleteSlider(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "deleteSlider", "request")
	defer span.End()
	sliderID := utils.StringToUint64(c.Param("id"))
	err := models.NewMysqlManager(ctx).DeleteSlider(c, ctx, sliderID)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "اسلایدر با موفقیت حذف شد",
	})
}
func IndexSlider(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "indexSlider", "request")
	defer span.End()
	dto, err := validations.IndexSlider(c)
	if err != nil {
		return
	}
	sliders, err := models.NewMysqlManager(ctx).GetAllSliderWithPagination(c, ctx, dto)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"sliders": sliders,
	})
}
