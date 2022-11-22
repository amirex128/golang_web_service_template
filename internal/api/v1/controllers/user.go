package controllers

import (
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/validations"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

// UpdateProfile
// @Summary ویرایش اطلاعات پروفایل کاربر
// @description کاربر میتواند اطلاعات پروفایل خود را ویرایش نماید
// @Tags user
// @Router       /user/profile/update [post]
// @Param	Authorization	 header string	true "Authentication"
// @Param	message	 body   DTOs.UpdateUser  	true "ورودی"
func UpdateProfile(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "registerLogin", "request")
	defer span.End()
	dto, err := validations.UpdateUser(c)
	if err != nil {
		return
	}

	userID := models.GetUser(c)

	if *userID == dto.ID {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "شما اجازه دسترسی به این صفحه را ندارید",
		})
		return
	}
	var user models.User
	if dto.Firstname != "" {
		user.Firstname = dto.Firstname
	}
	if dto.Lastname != "" {
		user.Lastname = dto.Lastname
	}
	if dto.Email != "" {
		user.Email = dto.Email
	}
	if dto.Mobile != "" {
		user.Mobile = dto.Mobile
	}
	if dto.CartNumber != "" {
		user.CartNumber = dto.CartNumber
	}
	if dto.Shaba != "" {
		user.Shaba = dto.Shaba
	}
	if dto.Password != "" {
		user.Password = dto.Password
	}
	err = models.NewMysqlManager(ctx).UpdateUser(c, ctx, &user)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "اطلاعات کاربری با موفقیت بروزرسانی شد",
	})

}
