package controllers

import (
	"backend/api/v1/validations"
	"backend/internal/app/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func requestCustomer(c *gin.Context) {
	login, err := validations.Register(c)
	if err != nil {
		return
	}
	user := &models.User{
		Mobile:   login.Mobile,
		Password: login.Password,
	}
	errRes := models.NewMainManager().CreateUser(user)
	if errRes != "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": errRes,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ثبت نام با موفقیت انجام شد",
	})
}
func verifyCustomer(c *gin.Context) {
	login, err := validations.Register(c)
	if err != nil {
		return
	}
	user := &models.User{
		Mobile:   login.Mobile,
		Password: login.Password,
	}
	errRes := models.NewMainManager().CreateUser(user)
	if errRes != "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": errRes,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ثبت نام با موفقیت انجام شد",
	})
}
func updateCustomer(c *gin.Context) {
	login, err := validations.Register(c)
	if err != nil {
		return
	}
	user := &models.User{
		Mobile:   login.Mobile,
		Password: login.Password,
	}
	errRes := models.NewMainManager().CreateUser(user)
	if errRes != "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": errRes,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ثبت نام با موفقیت انجام شد",
	})
}
