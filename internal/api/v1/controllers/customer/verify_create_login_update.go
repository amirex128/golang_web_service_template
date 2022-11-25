package customer

import (
	"github.com/amirex128/selloora_backend/internal/DTOs"
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/amirex128/selloora_backend/internal/validations"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

// VerifyCreateLoginUpdateCustomer
// @Summary دریافت اطلاعات مشتری و ویرایش و ثبت اطلاعات جدید در سبد خرید
// @description مشتریان با دریافت کد تایید میتوانند اطلاعات خود را وارد کنند و یا بروزرسانی نماییند همچنین دریافت کنند اطلاعات قبلی را
// @Tags customer
// @Router       /user/customer/verify [post]
// @Param message body DTOs.CreateUpdateCustomer true "ورودی"
func VerifyCreateLoginUpdateCustomer(c *gin.Context) {
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
		"message": "اطلاعات شما با موفقیت ثبت شد",
		"data":    customerNew,
	})
}
