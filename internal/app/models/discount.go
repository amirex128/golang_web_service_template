package models

import (
	"backend/internal/app/DTOs"
	"backend/internal/app/helpers"
	"encoding/gob"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strings"
)

type Discount struct {
	ID         uint64  `json:"id"`
	Code       string  `json:"code"`
	UserID     uint64  `json:"user_id"`
	StartedAt  string  `json:"started_at"`
	Count      uint32  `json:"count"`
	EndedAt    string  `json:"ended_at"`
	Type       string  `json:"type" sql:"type:ENUM('percent','amount')"` // ,
	Amount     float32 `json:"value"`
	Percent    float32 `json:"percent"`
	ProductIDs string  `json:"product_ids"`
	Status     byte    `json:"status"`
}
type DiscountArr []Discount

func (s DiscountArr) Len() int {
	return len(s)
}
func (s DiscountArr) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s DiscountArr) Less(i, j int) bool {
	return s[i].ID < s[j].ID
}

func (c *Discount) Encode(iw io.Writer) error {
	return gob.NewEncoder(iw).Encode(c)
}

func (c *Discount) Decode(ir io.Reader) error {
	return gob.NewDecoder(ir).Decode(c)
}

func initDiscount(manager *MysqlManager) {
	manager.GetConn().AutoMigrate(&Discount{})
}
func (m *MysqlManager) CreateDiscount(c *gin.Context, dto DTOs.CreateDiscount, userID uint64) error {

	for _, pId := range strings.Split(dto.ProductIDs, ",") {
		product, err := m.FindProductById(c, helpers.Uint64Convert(pId))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "محصول یافت نشد"})
			return err
		}
		if product.UserId != userID {
			c.JSON(http.StatusBadRequest, gin.H{"error": "شما اجازه ایجاد کد تخفیف برای این محصول را ندارید"})
			return err
		}
	}

	if m.GetConn().Where("code = ?", dto.Code).First(&Discount{}).RowsAffected > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "کد تخفیف تکراری است"})
		return nil
	}

	discount := Discount{
		Code:       dto.Code,
		UserID:     userID,
		StartedAt:  helpers.DateTimeConvert(dto.StartedAt),
		EndedAt:    helpers.DateTimeConvert(dto.EndedAt),
		Count:      dto.Count,
		Type:       dto.Type,
		Amount:     dto.Amount,
		Percent:    dto.Percent,
		ProductIDs: dto.ProductIDs,
		Status:     dto.Status,
	}
	err := m.GetConn().Create(discount).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "خطا در ایجاد کد تخفیف"})
		return err
	}
	return nil
}
func (m *MysqlManager) UpdateDiscount(c *gin.Context, discountID uint64) (Discount, error) {

	discount := Discount{}
	err := m.GetConn().Where("id = ?", discountID).First(&discount).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "کد تخفیف یافت نشد"})
		return discount, err
	}
	return discount, nil
}
func (m *MysqlManager) DeleteDiscount(c *gin.Context, discountID uint64) (Discount, error) {

	discount := Discount{}
	err := m.GetConn().Where("id = ?", discountID).First(&discount).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "کد تخفیف یافت نشد"})
		return discount, err
	}
	return discount, nil
}

func (m *MysqlManager) FindDiscountById(c *gin.Context, discountID uint64) (Discount, error) {

	discount := Discount{}
	err := m.GetConn().Where("id = ?", discountID).First(&discount).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "کد تخفیف یافت نشد"})
		return discount, err
	}
	return discount, nil
}
