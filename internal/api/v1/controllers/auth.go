package controllers

import (
	"fmt"
	"github.com/Squwid/go-randomizer"
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/amirex128/selloora_backend/internal/validations"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm"
	"net/http"
)

// RegisterLogin
// @Summary ثبت نام و ورود
// @description بعد از وارد کردن شماره همراه در صورت ثبت نام شما و تنظیم شدن پسورد بر روی اکانت شما باید پسورد خود را وارد نمایید برای ورود و در غیر این صورت باید کد تائید ارسال شده را وارد نماید تا توکن را دریافت نمایید
// @Tags auth
// @Router       /login/register [post]
// @Param	Authorization	header string	true "Authentication"
// @Param	message	body DTOs.RequestLoginRegister 	true "ورودی"
func RegisterLogin(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:registerLogin", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	dto, err := validations.RequestLoginRegister(c)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	user, err := models.NewMysqlManager(c).FindUserByMobile(dto.Mobile)
	if err == nil {
		if user.Password != "" {
			c.JSON(http.StatusOK, gin.H{
				"message":      "لطفا گذرواژه خود را برای ورود وارد نمایید",
				"is_register":  true,
				"has_password": true,
				"verify_code":  false,
			})
			return
		}
		dif := utils.DifferentWithNow(user.LastSendSMSAt)
		var randCode string
		var lastSendSMSAt string
		if dif < 7200 && dif > 0 {
			randCode = user.VerifyCode
			lastSendSMSAt = user.LastSendSMSAt
		} else {
			randCode = fmt.Sprintf("%d", randomizer.Number(1000, 9999))
			lastSendSMSAt = utils.NowTime()
		}

		newUser := models.User{
			ID:            user.ID,
			Mobile:        user.Mobile,
			VerifyCode:    randCode,
			LastSendSMSAt: lastSendSMSAt,
		}
		err = models.NewMysqlManager(c).UpdateUser(&newUser)
		if err != nil {
			return
		}

		text := fmt.Sprintf("%s : %s \n %s", "کد ورود به سامانه سلورا", randCode, "سلورا دستیار فروش شما")

		err := utils.SendSMS(c, ctx, user.Mobile, text, true)
		if err != nil {
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message":      "کد تایید به شماره همراه شما ارسال گردید",
			"is_register":  true,
			"has_password": false,
			"verify_code":  true,
		})
		return
	}

	randCode := fmt.Sprintf("%d", randomizer.Number(1000, 9999))
	text := fmt.Sprintf("%s : %s \n %s", "کد ورود به سامانه سلورا", randCode, "سلورا دستیار فروش شما")
	err = utils.SendSMS(c, ctx, dto.Mobile, text, true)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}

	newUser := models.User{
		Mobile:        dto.Mobile,
		VerifyCode:    randCode,
		LastSendSMSAt: utils.NowTime(),
	}

	err = models.NewMysqlManager(c).CreateUser(&newUser)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":      "ثبت نام شما با موفقیت انجام و کد تایید برای شما ارسال گردید",
		"is_register":  false,
		"has_password": false,
		"verify_code":  true,
	})
}

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
	userID := models.GetUser(c)
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

// ForgetPassword
// @Summary فراموشی رمز عبور
// @description در صورتی که کاربر پسورد خود را فراموش کرده باشد با استفاده از این وبسرویس درخواست کد تائید میتواند برای شماره همراه خود بدهد و بعد از دریافت کد تائید میتواند به وبسرویس ورود برود و در آنجا پسورد خود را با استفاده از توکن دریافتی عوض کند
// @Tags auth
// @Router       /forget [post]
// @Param	Authorization	header string	true "Authentication"
// @Param	message	body DTOs.ChangePassword 	true "ورودی"
func ForgetPassword(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:ForgetPassword", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	dto, err := validations.ForgetPassword(c)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	user, err := models.NewMysqlManager(c).FindUserByMobile(dto.Mobile)
	if err == nil {
		dif := utils.DifferentWithNow(user.LastSendSMSAt)
		var randCode string
		var lastSendSMSAt string
		if dif < 7200 && dif > 0 {
			randCode = user.VerifyCode
			lastSendSMSAt = user.LastSendSMSAt
		} else {
			randCode = fmt.Sprintf("%d", randomizer.Number(1000, 9999))
			lastSendSMSAt = utils.NowTime()
		}

		newUser := models.User{
			ID:            user.ID,
			Mobile:        user.Mobile,
			VerifyCode:    randCode,
			LastSendSMSAt: lastSendSMSAt,
		}
		err = models.NewMysqlManager(c).UpdateUser(&newUser)
		if err != nil {
			return
		}

		text := fmt.Sprintf("%s : %s \n %s", "کد ورود به سامانه سلورا", randCode, "سلورا دستیار فروش شما")

		err := utils.SendSMS(c, ctx, user.Mobile, text, true)
		if err != nil {
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message":      "کد تایید به شماره همراه شما ارسال گردید",
			"is_register":  true,
			"has_password": true,
			"verify_code":  true,
		})
		return
	}
}
