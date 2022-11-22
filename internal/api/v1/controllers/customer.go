package controllers

import (
	"fmt"
	"github.com/Squwid/go-randomizer"
	"github.com/amirex128/selloora_backend/internal/DTOs"
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/amirex128/selloora_backend/internal/validations"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

// RequestCreateLoginCustomer
// @Summary ثبت نام یا ورود کاربر با ارسال کد
// @description مشتریان میتوانند با ارسال شماره موبایل و کد تایید به سیستم وارد شوند یا اگر کاربری با این شماره موبایل وجود نداشته باشد ثبت نام کنند
// @Tags customer
// @Router       /user/customer/login/register [post]
// @Param message body DTOs.RequestCreateLoginCustomer true "ورودی"
func RequestCreateLoginCustomer(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:requestCreateLoginCustomer", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	dto, err := validations.RequestCreateLoginCustomer(c)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	customer, err := models.NewMysqlManager(c).FindCustomerByMobile(dto.Mobile)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	dif := utils.DifferentWithNow(customer.LastSendSMSAt)
	var randCode string
	var lastSendSMSAt string
	if dif < 7200 && dif > 0 {
		randCode = customer.VerifyCode
		lastSendSMSAt = customer.LastSendSMSAt
	} else {
		randCode = fmt.Sprintf("%d", randomizer.Number(1000, 9999))
		lastSendSMSAt = utils.NowTime()
	}
	if customer.ID > 0 {
		_, err = models.NewMysqlManager(c).UpdateCustomer(DTOs.CreateUpdateCustomer{
			Mobile:        dto.Mobile,
			VerifyCode:    randCode,
			LastSendSMSAt: lastSendSMSAt,
		})
		if err != nil {
			return
		}
	} else {
		err = models.NewMysqlManager(c).CreateCodeCustomer(dto, randCode)
		if err != nil {
			return
		}
	}
	shop, err := models.NewMysqlManager(c).FindShopByID(dto.ShopID)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	text := fmt.Sprintf("%s %s : %s", "کد تایید", shop.Name, randCode)
	err = utils.SendSMS(c, ctx, dto.Mobile, text, true)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "کد تایید برای شما ارسال گردید",
	})
	return
}

// VerifyCreateLoginCustomer
// @Summary دریافت اطلاعات مشتری و ویرایش و ثبت اطلاعات جدید در سبد خرید
// @description مشتریان با دریافت کد تایید میتوانند اطلاعات خود را وارد کنند و یا بروزرسانی نماییند همچنین دریافت کنند اطلاعات قبلی را
// @Tags customer
// @Router       /user/customer/verify [post]
// @Param message body DTOs.CreateUpdateCustomer true "ورودی"
func VerifyCreateLoginCustomer(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:verifyCreateLoginCustomer", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	dto, err := validations.CreateUpdateCustomer(c)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}

	customer, err := models.NewMysqlManager(c).FindCustomerByMobileAndVerifyCode(dto.Mobile, dto.VerifyCode)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	updateDto := DTOs.CreateUpdateCustomer{
		Mobile:     customer.Mobile,
		FullName:   dto.FullName,
		ProvinceID: dto.ProvinceID,
		CityID:     dto.CityID,
		Address:    dto.Address,
		PostalCode: dto.PostalCode,
	}
	customerNew, err := models.NewMysqlManager(c).UpdateCustomer(updateDto)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":  "اطلاعات شما با موفقیت ثبت شد",
		"customer": customerNew,
	})
}
