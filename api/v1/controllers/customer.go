package controllers

import (
	"backend/api/v1/validations"
	"backend/internal/app/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// createCustomer ثبت نام مشتری
func createCustomer(c *gin.Context) {
	dto, err := validations.CreateCustomer(c)
	if err != nil {
		return
	}
	err = models.NewMainManager().CreateCustomer(c, dto)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "اطلاعات شما با موفقیت ثبت شد",
	})
	return
}

// loginCustomer ارسال شماره موبایل برای ورود و در صورت عدم وجود شماره موبایل پیامک ارسال خواهد شد
func loginCustomer(c *gin.Context) {
	dto, err := validations.LoginCustomer(c)
	if err != nil {
		return
	}
	_, err = models.NewMainManager().FindCustomerByMobile(dto.Mobile)
	if err != nil {
		//TODO ارسال پیامک به مشتری
		c.JSON(http.StatusOK, gin.H{
			"message":   "شماره موبایل یا رمز عبور اشتباه است",
			"is_exists": false,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   "کد تایید یا پسورد خود را وارد نمایید",
		"is_exists": true,
	})
}

// verifyCustomer ارسال پسورد یا کد تایید برای ورود به حساب
func verifyCustomer(c *gin.Context) {
	dto, err := validations.VerifyCustomer(c)
	if err != nil {
		return
	}
	customer, err := models.NewMainManager().FindCustomerByMobileAndVerifyCode(dto.Mobile, dto.VerifyCode)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "رمز عبور یا کد تایید اشتباه است",
			"status":  false,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "شما با موفقیت وارد شدید",
		"status":   true,
		"customer": customer,
	})

}

// updateCustomer ویرایش اطلاعات مشتری
func updateCustomer(c *gin.Context) {
	dto, err := validations.UpdateCustomer(c)
	if err != nil {
		return
	}
	err = models.NewMainManager().UpdateCustomer(dto)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "اطلاعات شما با موفقیت بروز شد",
		})
		return
	}
}
