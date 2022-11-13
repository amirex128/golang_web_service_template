package models

import (
	"backend/internal/app/DTOs"
	"context"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

type Address struct {
	ID         uint64   `gorm:"primary_key;auto_increment" json:"id"`
	UserID     uint64   `gorm:"not null" json:"user_id"`
	Title      string   `gorm:"not null" json:"title"`
	ProvinceID uint64   `gorm:"not null" json:"province_id"`
	Province   Province `gorm:"foreignKey:province_id" json:"province"`
	CityID     uint64   `gorm:"not null" json:"city_id"`
	City       City     `gorm:"foreignKey:city_id" json:"city"`
	Address    string   `gorm:"not null" json:"address"`
	PostalCode string   `gorm:"not null" json:"postal_code"`
	Mobile     string   `gorm:"not null" json:"mobile"`
	FullName   string   `gorm:"not null" json:"full_name"`
	Lat        string   `json:"lat"`
	Long       string   `json:"long"`
}

func initAddress(manager *MysqlManager) {
	manager.GetConn().AutoMigrate(&Address{})

}

func (m *MysqlManager) CreateAddress(c *gin.Context, ctx context.Context, dto DTOs.CreateAddress) error {
	span, ctx := apm.StartSpan(ctx, "GetTicketWithChildren", "model")
	defer span.End()
	userID := GetUser(c)
	address := Address{
		UserID:     userID,
		Title:      dto.Title,
		ProvinceID: dto.ProvinceID,
		CityID:     dto.CityID,
		Address:    dto.Address,
		PostalCode: dto.PostalCode,
		Mobile:     dto.Mobile,
		FullName:   dto.FullName,
		Lat:        dto.Lat,
		Long:       dto.Long,
	}
	err := m.GetConn().Create(&address).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "خطایی در ایجاد آدرس رخ داده است",
			"error":   err.Error(),
			"type":    "model",
		})
		return err
	}
	return err
}

func (m *MysqlManager) UpdateAddress(c *gin.Context, ctx context.Context, dto DTOs.UpdateAddress) error {
	span, ctx := apm.StartSpan(ctx, "UpdateAddress", "model")
	defer span.End()
	address := Address{}
	err := m.GetConn().First(&address, dto.ID).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "خطایی در بروزرسانی آدرس رخ داده است",
			"error":   err.Error(),
			"type":    "model",
		})
		return err
	}
	userID := GetUser(c)
	if address.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "شما اجازه دسترسی به این آدرس را ندارید",
		})
		return err
	}
	if dto.FullName != "" {
		address.FullName = dto.FullName
	}
	if dto.Title != "" {
		address.Title = dto.Title
	}
	if dto.ProvinceID != 0 {
		address.ProvinceID = dto.ProvinceID
	}
	if dto.CityID != 0 {
		address.CityID = dto.CityID
	}
	if dto.Address != "" {
		address.Address = dto.Address
	}
	if dto.PostalCode != "" {
		address.PostalCode = dto.PostalCode
	}
	if dto.Mobile != "" {
		address.Mobile = dto.Mobile
	}
	if dto.Lat != "" {
		address.Lat = dto.Lat
	}
	if dto.Long != "" {
		address.Long = dto.Long
	}
	err = m.GetConn().Save(&address).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "خطایی در بروزرسانی آدرس رخ داده است",
			"error":   err.Error(),
			"type":    "model",
		})
		return err
	}
	return err
}

func (m *MysqlManager) DeleteAddress(c *gin.Context, ctx context.Context, addressID, userID uint64) error {
	span, ctx := apm.StartSpan(ctx, "DeleteAddress", "model")
	defer span.End()
	address := Address{}
	err := m.GetConn().First(&address, addressID).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "خطایی در حذف آدرس رخ داده است",
			"error":   err.Error(),
			"type":    "model",
		})
		return err
	}
	if address.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "شما اجازه دسترسی به این آدرس را ندارید",
		})
		return err
	}
	err = m.GetConn().Delete(&address).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "خطایی در حذف آدرس رخ داده است",
			"error":   err.Error(),
			"type":    "model",
		})
		return err
	}
	return err
}

func (m *MysqlManager) IndexAddress(c *gin.Context, ctx context.Context, userID uint64) ([]*Address, error) {
	span, ctx := apm.StartSpan(ctx, "IndexAddress", "model")
	defer span.End()
	var addresses []*Address
	err := m.GetConn().Where("user_id = ?", userID).Find(&addresses).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "خطایی در دریافت آدرس ها رخ داده است",
			"error":   err.Error(),
			"type":    "model",
		})
		return addresses, err
	}
	return addresses, err
}
