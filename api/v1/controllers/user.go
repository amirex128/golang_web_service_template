package controllers

import (
	"backend/api/v1/validations"
	"backend/internal/app/models"
	"backend/internal/app/utils"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

func updateProfile(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "registerLogin", "request")
	defer span.End()
	dto, err := validations.UpdateUser(c)
	if err != nil {
		return
	}

	id := utils.StringToUint64(c.Param("id"))
	userID := models.GetUser(c)

	if userID == id {
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
	err = models.NewMainManager().UpdateUser(c, ctx, &user)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "اطلاعات کاربری با موفقیت بروزرسانی شد",
	})

}
