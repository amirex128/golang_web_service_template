package post

import (
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/amirex128/selloora_backend/internal/validations"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

// IndexPost
// @Summary لیست مقالات
// @description فروشگاه ها میتوانند برای خود مقاله بسازند تا در سئو بهتر باشند
// @Tags post
// @Router       /user/post [get]
// @Param	Authorization	 header string	true "Authentication"
// @Param	search			 query   string	false "متن جستجو"
// @Param	page			 query   string	false "شماره صفحه"
// @Param	page_size		 query   string	false "تعداد صفحه"
// @Param	sort			 query   string	false "مرتب سازی براساس desc/asc"
func IndexPost(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:indexPost", "request")
	c.Request.WithContext(ctx)
	defer span.End()

	dto, err := validations.IndexPost(c)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	posts, err := models.NewMysqlManager(c).GetAllPostWithPagination(dto)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})

}
