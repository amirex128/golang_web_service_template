package discount

import (
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/amirex128/selloora_backend/internal/validations"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

// CreateDiscount
// @Summary ایجاد تخفیف
// @description ایجاد یک تخفیف بر روی یک محصصول یا چند محصول
// @Tags discount
// @Router       /user/discount/create [post]
// @Param	Authorization	 header string	true "Authentication"
// @Param	message	 body   DTOs.CreateDiscount  	true "ورودی"
func CreateDiscount(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:createDiscount", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	dto, err := validations.CreateDiscount(c)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	discount, err := models.NewMysqlManager(c).CreateDiscount(dto)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "تخفیف با موفقیت ایجاد شد",
		"data":    discount,
	})

}
