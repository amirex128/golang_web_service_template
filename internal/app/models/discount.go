package models

import (
	"backend/internal/app/DTOs"
	"backend/internal/app/utils"
	"encoding/gob"
	"errors"
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
	CreatedAt  string  `json:"created_at"`
	UpdatedAt  string  `json:"updated_at"`
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

	for _, pId := range dto.ProductIDs {
		product, err := m.FindProductById(c, pId)
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
		StartedAt:  utils.DateTimeConvert(dto.StartedAt),
		EndedAt:    utils.DateTimeConvert(dto.EndedAt),
		Count:      dto.Count,
		Type:       dto.Type,
		Amount:     dto.Amount,
		Percent:    dto.Percent,
		ProductIDs: strings.Join(utils.Uint64ToStringArray(dto.ProductIDs), ","),
		Status:     dto.Status,
		CreatedAt:  utils.NowTime(),
		UpdatedAt:  utils.NowTime(),
	}
	err := m.GetConn().Create(discount).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "خطا در ایجاد کد تخفیف"})
		return err
	}
	return nil
}
func (m *MysqlManager) UpdateDiscount(c *gin.Context, dto DTOs.UpdateDiscount, userID uint64) error {

	for _, pId := range dto.ProductIDs {
		product, err := m.FindProductById(c, pId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "محصول یافت نشد"})
			return err
		}
		if product.UserId != userID {
			c.JSON(http.StatusBadRequest, gin.H{"error": "شما اجازه ایجاد کد تخفیف برای این محصول را ندارید"})
			return err
		}
	}

	discount := &Discount{}
	err := m.GetConn().Where("id = ?", dto.DiscountID).First(discount).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "تخفیف یافت نشد"})
		return nil
	}

	if discount.UserID != userID {
		c.JSON(http.StatusBadRequest, gin.H{"error": "شما اجازه ویرایش این تخفیف را ندارید"})
		return errors.New("")
	}

	if dto.Code != "" {
		discount.Code = dto.Code
	}
	if dto.StartedAt != "" {
		discount.StartedAt = utils.DateTimeConvert(dto.StartedAt)
	}
	if dto.EndedAt != "" {
		discount.EndedAt = utils.DateTimeConvert(dto.EndedAt)
	}
	if dto.Count != 0 {
		discount.Count = dto.Count
	}
	if dto.Type != "" {
		discount.Type = dto.Type
	}
	if dto.Amount != 0 {
		discount.Amount = dto.Amount
	}
	if dto.Percent != 0 {
		discount.Percent = dto.Percent
	}
	if dto.ProductIDs != nil {
		discount.ProductIDs = strings.Join(utils.Uint64ToStringArray(dto.ProductIDs), ",")
	}
	if dto.Status != 0 {
		discount.Status = dto.Status
	}
	discount.UpdatedAt = utils.NowTime()
	err = m.GetConn().Save(discount).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "خطا در ویرایش تخفیف"})
		return err
	}
	return nil
}
func (m *MysqlManager) DeleteDiscount(c *gin.Context, discountID uint64, userID uint64) error {
	discount := Discount{}
	err := m.GetConn().Where("id = ?", discountID).First(&discount).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "تخفیف یافت نشد"})
		return err
	}
	if discount.UserID != userID {
		c.JSON(http.StatusBadRequest, gin.H{"error": "شما اجازه حذف این تخفیف را ندارید"})
		return errors.New("")
	}

	err = m.GetConn().Delete(&discount).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "خطا در حذف تخفیف"})
		return err
	}
	return nil
}
func (m *MysqlManager) IndexDiscount(c *gin.Context, search string, userID uint64) ([]Discount, error) {
	var discounts []Discount
	err := m.GetConn().Where("user_id = ?", userID).Where("code LIKE ?", "%"+search+"%").Find(&discounts).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "خطا در دریافت تخفیف ها"})
		return nil, err
	}
	return discounts, nil
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

func (m *MysqlManager) FindDiscountByCodeAndUserID(c *gin.Context, code string, userID uint64) (Discount, error) {

	discount := Discount{}
	err := m.GetConn().Where("code = ?", code).Where("user_id = ?", userID).First(&discount).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "کد تخفیف یافت نشد"})
		return discount, err
	}
	return discount, nil
}
