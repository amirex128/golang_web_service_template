package tag

import (
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/amirex128/selloora_backend/internal/validations"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

// CreateTag
// @Summary ایجاد تگ
// @description مقالات یا محصولات برای سئو بیشتر میتوانند بر روی خود تگ قرار دهند تا برای کلمات کلیدی مورد نظر بهترین نتیجه را به دست آورند
// @Tags tag
// @Router       /user/tag/create [post]
// @Param	Authorization	 header string	true "Authentication"
// @Param	message	 body   DTOs.CreateTag  	true "ورودی"
func CreateTag(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:createTag", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	dto, err := validations.CreateTag(c)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	tag, err := models.NewMysqlManager(c).CreateTag(dto)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "تگ با موفقیت ایجاد شد",
		"data":    tag,
	})

}
