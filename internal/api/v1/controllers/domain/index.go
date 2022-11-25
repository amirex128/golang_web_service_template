package domain

import (
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/amirex128/selloora_backend/internal/validations"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

// IndexDomain
// @Summary لیست دامنه ها
// @description لیست دامنه ها
// @Tags domain
// @Router       /user/domain [get]
// @Param	Authorization	 header string	true "Authentication"
// @Param	search			 query   string	false "متن جستجو"
// @Param	page			 query   string	false "شماره صفحه"
// @Param	page_size		 query   string	false "تعداد صفحه"
// @Param	sort			 query   string	false "مرتب سازی براساس desc/asc"
func IndexDomain(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:indexDomain", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	dto, err := validations.IndexDomain(c)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	domains, err := models.NewMysqlManager(c).GetAllDomainWithPagination(dto)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": domains,
	})
}
