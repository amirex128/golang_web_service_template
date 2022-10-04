package controllers

import (
	"backend/api/v1/validations"
	"backend/internal/app/models"
	"backend/internal/app/utils"
	"fmt"
	"github.com/Squwid/go-randomizer"
	"github.com/gin-gonic/gin"
	"net/http"
)

func register(c *gin.Context) {
	dto, err := validations.RequestLoginRegister(c)
	if err != nil {
		return
	}
	user, err := models.NewMainManager().FindUserByMobile(dto.Mobile)
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
		err = models.NewMainManager().UpdateUser(c, &newUser)
		if err != nil {
			return
		}

		text := fmt.Sprintf("%s : %s", "کد ورود سامانه سلورا :", randCode)
		err := utils.SendSMS(c, user.Mobile, text, true)
		if err != nil {
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "کد تایید به شماره همراه شما ارسال گردید",
		})
	}

	randCode := fmt.Sprintf("%d", randomizer.Number(1000, 9999))
	text := fmt.Sprintf("%s : %s", "کد ورود سامانه سلورا :", randCode)
	err = utils.SendSMS(c, dto.Mobile, text, true)
	if err != nil {
		return
	}

	newUser := models.User{
		Mobile:        dto.Mobile,
		VerifyCode:    randCode,
		LastSendSMSAt: utils.NowTime(),
	}

	err = models.NewMainManager().CreateUser(c, &newUser)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ثبت نام شما با موفقیت انجام و کد تایید برای شما ارسال گردید",
	})
}
