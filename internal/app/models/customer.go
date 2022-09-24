package models

import (
	"backend/internal/app/DTOs"
	"backend/internal/app/utils"
	"encoding/gob"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

type Customer struct {
	ID         uint64 `json:"id"`
	FullName   string `json:"full_name"`
	Mobile     string `json:"mobile"`
	ProvinceID uint32 `json:"province_id"`
	CityID     uint32 `json:"city_id"`
	Address    string `json:"address"`
	PostalCode string `json:"postal_code"`
	VerifyCode string `json:"verify_code"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}
type CustomerArr []Customer

func (s CustomerArr) Len() int {
	return len(s)
}
func (s CustomerArr) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s CustomerArr) Less(i, j int) bool {
	return s[i].ID < s[j].ID
}

func (c *Customer) Encode(iw io.Writer) error {
	return gob.NewEncoder(iw).Encode(c)
}

func (c *Customer) Decode(ir io.Reader) error {
	return gob.NewDecoder(ir).Decode(c)
}
func initCustomer(manager *MysqlManager) {
	manager.GetConn().AutoMigrate(&Customer{})
}
func (m *MysqlManager) FindCustomerById(c *gin.Context, customerID uint64) (Customer, error) {

	customer := Customer{}
	err := m.GetConn().Where("id = ?", customerID).First(&customer).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "مشتری یافت نشد"})
		return customer, err
	}
	return customer, nil
}
func (m *MysqlManager) FindCustomerByMobile(mobile string) (Customer, error) {
	customer := Customer{}
	err := m.GetConn().Where("mobile = ?", mobile).First(&customer).Error
	if err != nil {
		return customer, err
	}
	return customer, nil
}

func (m *MysqlManager) FindCustomerByMobileAndVerifyCode(mobile, verifyCode string) (Customer, error) {

	encryptPassword := utils.GeneratePasswordHash(verifyCode)
	customer := Customer{}
	err := m.GetConn().Where("mobile = ?", mobile).Where("verifyCode = ?", encryptPassword).First(&customer).Error
	if err != nil {
		return customer, err
	}
	return customer, nil
}

func (m *MysqlManager) UpdateCustomer(dto DTOs.UpdateCustomer) error {

	customer := &Customer{
		FullName:   dto.FullName,
		ProvinceID: dto.ProvinceID,
		CityID:     dto.CityID,
		Address:    dto.Address,
		PostalCode: dto.PostalCode,
		CreatedAt:  utils.NowTime(),
	}
	err := m.GetConn().Create(customer).Error
	if err != nil {
		return err
	}
	return nil
}
