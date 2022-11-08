package controllers

import (
	"backend/api/v1/validations"
	"backend/internal/app/DTOs"
	"backend/internal/app/models"
	"backend/internal/app/utils"
	"fmt"
	"github.com/Squwid/go-randomizer"
	"github.com/gin-gonic/gin"
	"net/http"
)

func requestCreateLoginCustomer(c *gin.Context) {
	dto, err := validations.RequestCreateLoginCustomer(c)
	if err != nil {
		return
	}
	customer, err := models.NewMainManager().FindCustomerByMobile(c, dto.Mobile)
	if err != nil {
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
		_, err = models.NewMainManager().UpdateCustomer(c, DTOs.CreateUpdateCustomer{
			Mobile:        dto.Mobile,
			VerifyCode:    randCode,
			LastSendSMSAt: lastSendSMSAt,
		})
		if err != nil {
			return
		}
	} else {
		err = models.NewMainManager().CreateCodeCustomer(c, dto, randCode)
		if err != nil {
			return
		}
	}
	shop, err := models.NewMainManager().FindShopByID(c, dto.ShopID, 0)
	if err != nil {
		return
	}
	text := fmt.Sprintf("%s %s : %s", "کد تایید", shop.Name, randCode)
	err = utils.SendSMS(c, dto.Mobile, text, true)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "کد تایید برای شما ارسال گردید",
	})
	return
}

func verifyCreateLoginCustomer(c *gin.Context) {
	dto, err := validations.CreateUpdateCustomer(c)
	if err != nil {
		return
	}

	customer, err := models.NewMainManager().FindCustomerByMobileAndVerifyCode(c, dto.Mobile, dto.VerifyCode)
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
	customerNew, err := models.NewMainManager().UpdateCustomer(c, updateDto)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":  "اطلاعات شما با موفقیت ثبت شد",
		"customer": customerNew,
	})
}
