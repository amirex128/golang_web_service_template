package auth

import (
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/amirex128/selloora_backend/internal/validations"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

// ChangePassword
// @Summary تغییر پسورد
// @description درصورتی که کاربر ثبت نام نکرده بود و کد تائید را وارد نمود یک توکن دریافت میکند که میتواند با استفاده از آن توکن پسورد خود را تغییر دهد
// @Tags auth
// @Router       /profile/change-password [post]
// @Param	Authorization	header string	true "Authentication"
// @Param	message	body DTOs.ChangePassword 	true "ورودی"
func ChangePassword(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:changePassword", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	dto, err := validations.ChangePassword(c)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	userID := utils.GetUserID(c)
	err = models.NewMysqlManager(c).UpdateUser(&models.User{
		ID:       *userID,
		Password: models.GeneratePasswordHash(dto.Password),
	})
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "رمز عبور شما با موفقیت تغییر یافت",
		"changed": true,
	})
}
