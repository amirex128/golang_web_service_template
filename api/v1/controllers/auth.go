package controllers

import (
	"backend/api/v1/validations"
	"backend/internal/app/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func register(c *gin.Context) {
	login, err := validations.Register(c)
	if err != nil {
		return
	}
	user := &models.User{
		Mobile:   login.Mobile,
		Password: login.Password,
	}
	err = models.NewMainManager().CreateUser(c, user)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ثبت نام با موفقیت انجام شد",
	})
}
func forget(c *gin.Context) {

}
