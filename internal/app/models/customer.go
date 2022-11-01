package models

import (
	"backend/internal/app/DTOs"
	"backend/internal/app/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Customer struct {
	ID            uint64 `gorm:"primary_key;auto_increment" json:"id"`
	FullName      string `json:"full_name"`
	Mobile        string `json:"mobile"`
	ProvinceID    uint32 `json:"province_id"`
	CityID        uint32 `json:"city_id"`
	Address       string `json:"address"`
	PostalCode    uint64 `json:"postal_code"`
	VerifyCode    string `json:"verify_code"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
	LastSendSMSAt string `json:"last_send_sms_at"`
}

func initCustomer(manager *MysqlManager) {
	manager.GetConn().AutoMigrate(&Customer{})
	manager.CreateCustomer(&gin.Context{}, DTOs.CreateUpdateCustomer{
		ShopID:        1,
		Mobile:        "09123456789",
		VerifyCode:    "1234",
		FullName:      "محمد محمدی",
		ProvinceID:    1,
		CityID:        1,
		Address:       "تهران",
		PostalCode:    1234567890,
		LastSendSMSAt: "2020-01-01 00:00:00",
	})
}
func (m *MysqlManager) FindCustomerById(c *gin.Context, customerID uint64) (Customer, error) {

	customer := Customer{}
	err := m.GetConn().Where("id = ?", customerID).First(&customer).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "مشتری یافت نشد",
			"error":   err.Error(),
			"type":    "model",
		})
		return customer, err
	}
	return customer, nil
}
func (m *MysqlManager) FindCustomerByMobile(c *gin.Context, mobile string) (Customer, error) {
	customer := Customer{}
	err := m.GetConn().Where("mobile = ?", mobile).Find(&customer).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "شماره موبایل تکراری یا اشتباه میباشد",
			"error":   err.Error(),
			"type":    "model",
		})
		return customer, err
	}
	return customer, nil
}

func (m *MysqlManager) FindCustomerByMobileAndVerifyCode(c *gin.Context, mobile, verifyCode string) (Customer, error) {

	customer := Customer{}
	err := m.GetConn().Where("mobile = ?", mobile).Where("verify_code = ?", verifyCode).First(&customer).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "رمز عبور یا کد تایید اشتباه است",
			"error":   err.Error(),
			"type":    "model",
		})
		return customer, err
	}
	return customer, nil
}
func (m *MysqlManager) CreateCustomer(c *gin.Context, dto DTOs.CreateUpdateCustomer) error {
	customer := Customer{
		FullName:   dto.FullName,
		Mobile:     dto.Mobile,
		ProvinceID: dto.ProvinceID,
		CityID:     dto.CityID,
		Address:    dto.Address,
		PostalCode: dto.PostalCode,
		VerifyCode: dto.VerifyCode,
	}
	err := m.GetConn().Create(&customer).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "مشکلی در ثبت نام شما پیش آمده است",
			"error":   err.Error(),
			"type":    "model",
		})
		return err
	}
	return nil
}
func (m *MysqlManager) CreateCodeCustomer(c *gin.Context, dto DTOs.RequestCreateLoginCustomer, encryptPassword string) error {
	rowsAffected := m.GetConn().Where("mobile = ?", dto.Mobile).First(&Customer{}).RowsAffected
	if rowsAffected > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "شماره موبایل قبلا ثبت شده است",
		})
		return errors.New("mobile failed")
	}

	customer := Customer{
		Mobile:        dto.Mobile,
		VerifyCode:    encryptPassword,
		LastSendSMSAt: utils.NowTime(),
	}
	err := m.GetConn().Create(&customer).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "مشکلی در ثبت نام شما پیش آمده است",
			"error":   err.Error(),
			"type":    "model",
		})
		return err
	}
	return nil
}
func (m *MysqlManager) UpdateCustomer(c *gin.Context, dto DTOs.CreateUpdateCustomer) (Customer, error) {
	customer, err := m.FindCustomerByMobile(c, dto.Mobile)
	if err != nil {
		return customer, err
	}

	if dto.FullName != "" {
		customer.FullName = dto.FullName
	}
	if dto.ProvinceID != 0 {
		customer.ProvinceID = dto.ProvinceID
	}
	if dto.CityID != 0 {
		customer.CityID = dto.CityID
	}
	if dto.Address != "" {
		customer.Address = dto.Address
	}
	if dto.PostalCode != 0 {
		customer.PostalCode = dto.PostalCode
	}
	if dto.VerifyCode != "" {
		customer.VerifyCode = dto.VerifyCode
	}
	err = m.GetConn().Save(&customer).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "مشکلی در ویرایش اطلاعات شما پیش آمده است",
			"error":   err.Error(),
			"type":    "model",
		})
		return customer, err
	}
	return customer, nil
}
