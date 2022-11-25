package tag

import (
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/amirex128/selloora_backend/internal/validations"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

// IndexTag
// @Summary لیست تگ ها
// @description مقالات یا محصولات برای سئو بیشتر میتوانند بر روی خود تگ قرار دهند تا برای کلمات کلیدی مورد نظر بهترین نتیجه را به دست آورند
// @Tags tag
// @Router       /user/tag [get]
// @Param	Authorization	 header string	true "Authentication"
// @Param	search			 query   string	false "متن جستجو"
// @Param	page			 query   string	false "شماره صفحه"
// @Param	page_size		 query   string	false "تعداد صفحه"
// @Param	sort			 query   string	false "مرتب سازی براساس desc/asc"
func IndexTag(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:indexTag", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	dto, err := validations.IndexTag(c)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	pagination, err := models.NewMysqlManager(c).GetAllTagsWithPagination(dto)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"tags": pagination,
	})
}
