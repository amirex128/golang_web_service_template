package controllers

import (
	"backend/internal/app/DTOs"
	"backend/internal/app/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func register(c *gin.Context) {
	login := new(DTOs.Login)
	err := c.BindJSON(&login)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "مقادیر ارسال شده نا درست میباشد",
			"error":   err.Error(),
		})
		return
	}
	err = validate.Struct(login)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "مقادیر ارسال شده نا درست میباشد",
				"error":   err.Error(),
			})
			return
		}
		var errors []gin.H
		for _, err := range err.(validator.ValidationErrors) {
			if err.StructField() == "Mobile" {
				errors = append(errors, gin.H{
					"message": "شماره موبایل نامعتبر میباشد",
				})
			}
			if err.StructField() == "Password" {
				if err.Tag() == "min" {
					errors = append(errors, gin.H{
						"message": "رمز عبور باید حداقل 8 کاراکتر باشد",
					})
				}
				if err.Tag() == "max" {
					errors = append(errors, gin.H{
						"message": "رمز عبور باید حداکثر 20 کاراکتر باشد",
					})
				}
				if err.Tag() == "required" {
					errors = append(errors, gin.H{
						"message": "رمز عبور باید وارد شود",
					})
				}

			}

			c.JSON(http.StatusBadRequest, errors)
			return
		}
	}
	user := &models.User{}
	err = models.NewMainManager().CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "خطایی در فرایند ثبت نام شما رخ داده است لطفا مجدد تلاش نمایید",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, login)
}
func login(c *gin.Context) {
	login := new(DTOs.Login)
	err := c.BindJSON(&login)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "مقادیر ارسال شده نا درست میباشد",
			"error":   err.Error(),
		})
		return
	}
	err = validate.Struct(login)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "مقادیر ارسال شده نا درست میباشد",
				"error":   err.Error(),
			})
			return
		}
		var errors []gin.H
		for _, err := range err.(validator.ValidationErrors) {
			if err.StructField() == "Mobile" {
				errors = append(errors, gin.H{
					"message": "شماره موبایل نامعتبر میباشد",
				})
			}
			if err.StructField() == "Password" {
				if err.Tag() == "min" {
					errors = append(errors, gin.H{
						"message": "رمز عبور باید حداقل 8 کاراکتر باشد",
					})
				}
				if err.Tag() == "max" {
					errors = append(errors, gin.H{
						"message": "رمز عبور باید حداکثر 20 کاراکتر باشد",
					})
				}
				if err.Tag() == "required" {
					errors = append(errors, gin.H{
						"message": "رمز عبور باید وارد شود",
					})
				}

			}

			c.JSON(http.StatusBadRequest, errors)
			return
		}
	}
	user := &models.User{}
	err = models.NewMainManager().CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "خطایی در فرایند ثبت نام شما رخ داده است لطفا مجدد تلاش نمایید",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, login)
}
