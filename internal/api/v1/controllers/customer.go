package controllers

import (
	"fmt"
	"github.com/Squwid/go-randomizer"
	"github.com/amirex128/selloora_backend/internal/DTOs"
	"github.com/amirex128/selloora_backend/internal/models"
	utils2 "github.com/amirex128/selloora_backend/internal/utils"
	"github.com/amirex128/selloora_backend/internal/validations"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

func RequestCreateLoginCustomer(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "requestCreateLoginCustomer", "request")
	defer span.End()
	dto, err := validations.RequestCreateLoginCustomer(c)
	if err != nil {
		return
	}
	customer, err := models.NewMysqlManager(ctx).FindCustomerByMobile(c, ctx, dto.Mobile)
	if err != nil {
		return
	}
	dif := utils2.DifferentWithNow(customer.LastSendSMSAt)
	var randCode string
	var lastSendSMSAt string
	if dif < 7200 && dif > 0 {
		randCode = customer.VerifyCode
		lastSendSMSAt = customer.LastSendSMSAt
	} else {
		randCode = fmt.Sprintf("%d", randomizer.Number(1000, 9999))
		lastSendSMSAt = utils2.NowTime()
	}
	if customer.ID > 0 {
		_, err = models.NewMysqlManager(ctx).UpdateCustomer(c, nil, DTOs.CreateUpdateCustomer{
			Mobile:        dto.Mobile,
			VerifyCode:    randCode,
			LastSendSMSAt: lastSendSMSAt,
		})
		if err != nil {
			return
		}
	} else {
		err = models.NewMysqlManager(ctx).CreateCodeCustomer(c, ctx, dto, randCode)
		if err != nil {
			return
		}
	}
	shop, err := models.NewMysqlManager(ctx).FindShopByID(c, ctx, dto.ShopID)
	if err != nil {
		return
	}
	text := fmt.Sprintf("%s %s : %s", "کد تایید", shop.Name, randCode)
	err = utils2.SendSMS(c, ctx, dto.Mobile, text, true)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "کد تایید برای شما ارسال گردید",
	})
	return
}

func VerifyCreateLoginCustomer(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "verifyCreateLoginCustomer", "request")
	defer span.End()
	dto, err := validations.CreateUpdateCustomer(c)
	if err != nil {
		return
	}

	customer, err := models.NewMysqlManager(ctx).FindCustomerByMobileAndVerifyCode(c, ctx, dto.Mobile, dto.VerifyCode)
	if err != nil {
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
	customerNew, err := models.NewMysqlManager(ctx).UpdateCustomer(c, ctx, updateDto)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":  "اطلاعات شما با موفقیت ثبت شد",
		"customer": customerNew,
	})
}
