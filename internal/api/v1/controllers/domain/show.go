package domain

import (
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

// ShowDomain
// @Summary نمایش دامنه
// @description هر فروشگاه برای خود میتواند به تعداد دلخواه دامنه در موقعیت های مختلف مثل بالای صفحه و پایین صفحه ایجاد نماید
// @Tags domain
// @Router       /user/domain/show/{id} [get]
// @Param	Authorization	 header string	true "Authentication"
// @Param	id			 path   string	true "شناسه " SchemaExample(1)
func ShowDomain(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:showDomain", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	domainID := c.Param("id")
	domain, err := models.NewMysqlManager(c).FindDomainByID(utils.StringToUint64(domainID))
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": domain,
	})
}
