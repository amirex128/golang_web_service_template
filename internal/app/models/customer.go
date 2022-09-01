package models

import (
	"encoding/gob"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

type Customer struct {
	ID         uint64 `json:"id"`
	FullName   string `json:"full_name"`
	Mobile     string `json:"mobile"`
	ProvinceID int64  `json:"province_id"`
	CityID     int64  `json:"city_id"`
	Address    string `json:"address"`
	PostalCode string `json:"postal_code"`
	VerifyCode uint16 `json:"verify_code"`
	CreatedAt  string `json:"created_at"`
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
