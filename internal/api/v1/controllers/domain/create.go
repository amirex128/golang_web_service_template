package domain

import (
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/amirex128/selloora_backend/internal/validations"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

// CreateDomain
// @Summary ایجاد دامنه
// @description یک دامنه یا یک ساب دامنه کاربر میتواند اضافه نمایید تا سایت ایجاد شده خود را بر بستر آن دامنه مشاهده نماید
// @Tags domain
// @Router       /user/domain/create [post]
// @Param	Authorization	 header string	true "Authentication"
// @Param	message	 body   DTOs.CreateDomain  	true "ورودی"
func CreateDomain(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:createDomain", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	dto, err := validations.CreateDomain(c)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	domain, err := models.NewMysqlManager(c).CreateDomain(dto)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "دامنه با موفقیت ایجاد شد",
		"data":    domain,
	})
}
