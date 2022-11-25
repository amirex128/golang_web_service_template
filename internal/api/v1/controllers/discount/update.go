package discount

import (
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/amirex128/selloora_backend/internal/validations"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

// UpdateDiscount
// @Summary ویرایش تخفیف
// @description ویرایش تخفیف
// @Tags discount
// @Router       /user/discount/update/{id} [post]
// @Param	Authorization	 header string	true "Authentication"
// @Param	id			 path   string	false "شناسه تخفیف" SchemaExample(1)
// @Param	message	 body   DTOs.UpdateDiscount  	true "ورودی"
func UpdateDiscount(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:updateDiscount", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	dto, err := validations.UpdateDiscount(c)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	err = models.NewMysqlManager(c).UpdateDiscount(dto)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "تخفیف با موفقیت بروزرسانی شد",
	})
}
