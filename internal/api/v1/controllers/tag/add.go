package tag

import (
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/amirex128/selloora_backend/internal/validations"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

// AddTag
// @Summary افزودن یک تگ به محصول یا مقاله
// @description مقالات یا محصولات برای سئو بیشتر میتوانند بر روی خود تگ قرار دهند تا برای کلمات کلیدی مورد نظر بهترین نتیجه را به دست آورند
// @Tags tag
// @Router       /user/tag/add [post]
// @Param	Authorization	 header string	true "Authentication"
// @Param	message	 body   DTOs.CreateTag  	true "ورودی"
func AddTag(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:addTag", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	dto, err := validations.AddTag(c)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	err = models.NewMysqlManager(c).AddTag(dto)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "تگ با موفقیت به پست اضافه شد",
	})
}
