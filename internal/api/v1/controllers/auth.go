package controllers

import (
	"fmt"
	"github.com/Squwid/go-randomizer"
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/amirex128/selloora_backend/internal/validations"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm"
	"net/http"
)

func RegisterLogin(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "registerLogin", "request")
	defer span.End()
	dto, err := validations.RequestLoginRegister(c)
	if err != nil {
		return
	}
	user, err := models.NewMysqlManager(ctx).FindUserByMobile(ctx, dto.Mobile)
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
		err = models.NewMysqlManager(ctx).UpdateUser(c, ctx, &newUser)
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
		return
	}

	newUser := models.User{
		Mobile:        dto.Mobile,
		VerifyCode:    randCode,
		LastSendSMSAt: utils.NowTime(),
	}

	err = models.NewMysqlManager(ctx).CreateUser(c, ctx, &newUser)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":      "ثبت نام شما با موفقیت انجام و کد تایید برای شما ارسال گردید",
		"is_register":  false,
		"has_password": false,
		"verify_code":  true,
	})
}

func ChangePassword(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "changePassword", "request")
	defer span.End()
	dto, err := validations.ChangePassword(c)
	if err != nil {
		return
	}
	userID := models.GetUser(c)
	err = models.NewMysqlManager(ctx).UpdateUser(c, ctx, &models.User{
		ID:       *userID,
		Password: models.GeneratePasswordHash(dto.Password),
	})
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "رمز عبور شما با موفقیت تغییر یافت",
		"changed": true,
	})
}

//func ForgetPassword(c *gin.Context) {
//	span, ctx := apm.StartSpan(c.Request.Context(), "ForgetPassword", "request")
//	defer span.End()
//	dto, err := validations.ForgetPassword(c)
//	if err != nil {
//		return
//	}
//	user, err := models.NewMysqlManager(ctx).FindUserByMobile(ctx, dto.Mobile)
//
//}
