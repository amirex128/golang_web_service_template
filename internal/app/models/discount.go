package models

import (
	"backend/internal/app/DTOs"
	"backend/internal/app/utils"
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
	"strings"
)

type Discount struct {
	ID         uint64  `gorm:"primary_key;auto_increment" json:"id"`
	Code       string  `json:"code"`
	UserID     uint64  `json:"user_id"`
	User       User    `gorm:"foreignKey:user_id" json:"user"`
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

func initDiscount(manager *MysqlManager) {
	manager.GetConn().AutoMigrate(&Discount{})
	manager.CreateDiscount(&gin.Context{}, context.Background(), DTOs.CreateDiscount{
		Code:       "test",
		StartedAt:  "2021-01-01 00:00:00",
		EndedAt:    "2024-01-01 00:00:00",
		Count:      10,
		Type:       "percent",
		Amount:     0,
		Percent:    20,
		ProductIDs: []uint64{1, 2, 3},
		Status:     1,
	})
}
func (m *MysqlManager) CreateDiscount(c *gin.Context, ctx context.Context, dto DTOs.CreateDiscount) error {
	span, ctx := apm.StartSpan(ctx, "showDiscount", "model")
	defer span.End()
	userID := GetUser(c)
	for _, pId := range dto.ProductIDs {
		product, err := m.FindProductById(c, ctx, pId)
		if err != nil {
			return err
		}
		if product.UserID != userID {
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
	err := m.GetConn().Create(&discount).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "خطا در ایجاد کد تخفیف",
			"error":   err.Error(),
			"type":    "model",
		})
		return err
	}
	return nil
}
func (m *MysqlManager) UpdateDiscount(c *gin.Context, ctx context.Context, dto DTOs.UpdateDiscount) error {
	span, ctx := apm.StartSpan(ctx, "showDiscount", "model")
	defer span.End()
	userID := GetUser(c)

	for _, pId := range dto.ProductIDs {
		product, err := m.FindProductById(c, ctx, pId)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "محصول یافت نشد",
				"error":   err.Error(),
				"type":    "model",
			})
			return err
		}
		if product.UserID != userID {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "شما اجازه ایجاد کد تخفیف برای این محصول را ندارید",
				"error":   err.Error(),
				"type":    "model",
			})
			return err
		}
	}

	discount := &Discount{}
	err := m.GetConn().Where("id = ?", dto.ID).First(discount).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "تخفیف یافت نشد",
			"error":   err.Error(),
			"type":    "model",
		})
		return err
	}

	if discount.UserID != userID {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "شما اجازه ویرایش این تخفیف را ندارید",
			"error":   err.Error(),
			"type":    "model",
		})
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
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "خطا در ویرایش تخفیف"})
		return err
	}
	return nil
}
func (m *MysqlManager) DeleteDiscount(c *gin.Context, ctx context.Context, discountID uint64) error {
	span, ctx := apm.StartSpan(ctx, "showDiscount", "model")
	defer span.End()
	discount := Discount{}
	err := m.GetConn().Where("id = ?", discountID).First(&discount).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "تخفیف یافت نشد",
			"error":   err.Error(),
			"type":    "model",
		})
		return err
	}
	userID := GetUser(c)
	if discount.UserID != userID {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "شما اجازه حذف این تخفیف را ندارید",
			"error":   err.Error(),
			"type":    "model",
		})
		return errors.New("")
	}

	err = m.GetConn().Delete(&discount).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "خطا در حذف تخفیف",
			"error":   err.Error(),
			"type":    "model",
		})
		return err
	}
	return nil
}
func (m *MysqlManager) GetAllDiscountWithPagination(c *gin.Context, ctx context.Context, dto DTOs.IndexDiscount) (*DTOs.Pagination, error) {
	span, ctx := apm.StartSpan(ctx, "showDiscount", "model")
	defer span.End()
	conn := m.GetConn()
	var discounts []Discount
	pagination := &DTOs.Pagination{PageSize: dto.PageSize, Page: dto.Page}
	userID := GetUser(c)
	conn = conn.Scopes(DTOs.Paginate("discounts", pagination, conn))
	if dto.Search != "" {
		conn = conn.Where("name LIKE ?", "%"+dto.Search+"%").Where("user_id = ? ", userID).Order("id DESC")
	}
	err := conn.Find(&discounts).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "خطا در دریافت تخفیف ها",
			"error":   err.Error(),
			"type":    "model",
		})
		return pagination, err
	}
	pagination.Data = discounts
	return pagination, nil
}

func (m *MysqlManager) FindDiscountById(c *gin.Context, ctx context.Context, discountID uint64) (Discount, error) {
	span, ctx := apm.StartSpan(ctx, "showDiscount", "model")
	defer span.End()

	discount := Discount{}
	err := m.GetConn().Where("id = ?", discountID).First(&discount).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "کد تخفیف یافت نشد",
			"error":   err.Error(),
			"type":    "model",
		})
		return discount, err
	}
	return discount, nil
}

func (m *MysqlManager) FindDiscountByCodeAndUserID(c *gin.Context, ctx context.Context, code string) (Discount, error) {
	span, ctx := apm.StartSpan(ctx, "showDiscount", "model")
	defer span.End()
	userID := GetUser(c)
	discount := Discount{}
	err := m.GetConn().Where("code = ?", code).Where("user_id = ?", userID).First(&discount).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "کد تخفیف یافت نشد",
			"error":   err.Error(),
			"type":    "model",
		})
		return discount, err
	}
	return discount, nil
}
