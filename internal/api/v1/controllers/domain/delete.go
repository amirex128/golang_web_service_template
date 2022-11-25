package domain

import (
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

// DeleteDomain
// @Summary حذف دامنه
// @description باحذف دامنه امکان دسترسی از این دامنه بر روی سایت کاربر گرفته میشود
// @Tags domain
// @Router       /user/domain/delete/{id} [post]
// @Param	Authorization	 header string	true "Authentication"
// @Param	id			 path   string	true "شناسه دامنه" SchemaExample(1)
func DeleteDomain(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:deleteDomain", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	domainID := utils.StringToUint64(c.Param("id"))
	err := models.NewMysqlManager(c).DeleteDomain(domainID)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "دامنه با موفقیت حذف شد",
	})
}
