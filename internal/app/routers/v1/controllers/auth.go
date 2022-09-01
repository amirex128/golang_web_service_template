package controllers

import (
	"backend/internal/app/models"
	"backend/internal/app/routers/v1/validations"
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
	errRes := models.NewMainManager().CreateUser(user)
	if errRes != "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": errRes,
		})
		return
	}

	c.JSON(http.StatusOK, login)
}
