package models

import (
	"github.com/amirex128/selloora_backend/internal/DTOs"
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/brianvoe/gofakeit/v6"
	"go.elastic.co/apm/v2"
	"strings"
)

type Discount struct {
	ID         uint64  `gorm:"primary_key;auto_increment" json:"id"`
	Code       string  `json:"code"`
	UserID     *uint64 `gorm:"default:null" json:"user_id"`
	User       *User   `gorm:"foreignKey:user_id" json:"user"`
	StartedAt  string  `json:"started_at"`
	Count      uint32  `json:"count"`
	EndedAt    string  `json:"ended_at"`
	Type       string  `json:"type" sql:"type:ENUM('percent','amount')"` // ,
	Amount     float32 `json:"value"`
	Percent    float32 `json:"percent"`
	ProductIDs string  `json:"product_ids"`
	Status     bool    `json:"status"`
	CreatedAt  string  `json:"created_at"`
	UpdatedAt  string  `json:"updated_at"`
}

func initDiscount(manager *MysqlManager) {
	if !manager.GetConn().Migrator().HasTable(&Discount{}) {
		manager.GetConn().Migrator().CreateTable(&Discount{})

		for i := 0; i < 100; i++ {
			model := new(DTOs.CreateDiscount)
			gofakeit.Struct(model)

			manager.CreateDiscount(*model)
		}

	}

}
func (m *MysqlManager) CreateDiscount(dto DTOs.CreateDiscount) (*Discount, error) {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:showDiscount", "model")
	defer span.End()
	userID := utils.GetUserID(m.Ctx)
	for _, pId := range dto.ProductIDs {
		product, err := m.FindProductById(pId)
		if err != nil {
			return nil, err
		}
		if err := utils.CheckAccess(m.Ctx, product.UserID); err != nil {
			return nil, err
		}
	}

	if m.GetConn().Where("code = ?", dto.Code).First(&Discount{}).RowsAffected > 0 {
		return nil, errorx.New("کد تخفیف تکراری است", "model", nil)
	}

	discount := &Discount{
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
		return nil, errorx.New("خطا در ایجاد کد تخفیف", "model", err)
	}
	return discount, nil
}
func (m *MysqlManager) UpdateDiscount(dto DTOs.UpdateDiscount) error {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:showDiscount", "model")
	defer span.End()
	userID := utils.GetUserID(m.Ctx)

	for _, pId := range dto.ProductIDs {
		product, err := m.FindProductById(pId)
		if err != nil {
			return errorx.New("محصول یافت نشد", "model", err)
		}
		if err := utils.CheckAccess(m.Ctx, product.UserID); err != nil {
			return err
		}
	}

	discount := &Discount{}
	err := m.GetConn().Where("id = ?", dto.ID).First(discount).Error
	if err != nil {
		return errorx.New("تخفیف یافت نشد", "model", err)
	}

	if *discount.UserID != *userID {
		return errorx.New("شما اجازه ویرایش این تخفیف را ندارید", "model", err)
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
	discount.Status = dto.Status
	discount.UpdatedAt = utils.NowTime()
	err = m.GetConn().Save(discount).Error
	if err != nil {
		return errorx.New("خطا در ویرایش تخفیف", "model", err)
	}
	return nil
}
func (m *MysqlManager) DeleteDiscount(discountID uint64) error {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:showDiscount", "model")
	defer span.End()
	discount := Discount{}
	err := m.GetConn().Where("id = ?", discountID).First(&discount).Error
	if err != nil {
		return errorx.New("تخفیف یافت نشد", "model", err)
	}
	if err := utils.CheckAccess(m.Ctx, discount.UserID); err != nil {
		return err
	}

	err = m.GetConn().Delete(&discount).Error
	if err != nil {
		return errorx.New("خطا در حذف تخفیف", "model", err)
	}
	return nil
}
func (m *MysqlManager) GetAllDiscountWithPagination(dto DTOs.IndexDiscount) (*DTOs.Pagination, error) {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:showDiscount", "model")
	defer span.End()
	conn := m.GetConn()
	var discounts []Discount
	pagination := &DTOs.Pagination{PageSize: dto.PageSize, Page: dto.Page}
	conn = conn.Scopes(DTOs.Paginate("discounts", pagination, conn))
	if dto.Search != "" {
		conn = conn.Where("name LIKE ?", "%"+dto.Search+"%")
	}
	err := conn.Where("user_id = ? ", utils.GetUserID(m.Ctx)).Order("id DESC").Find(&discounts).Error
	if err != nil {
		return nil, errorx.New("خطا در دریافت تخفیف ها", "model", err)
	}
	pagination.Data = discounts
	return pagination, nil
}

func (m *MysqlManager) FindDiscountById(discountID uint64) (*Discount, error) {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:showDiscount", "model")
	defer span.End()

	discount := &Discount{}
	err := m.GetConn().Where("id = ?", discountID).First(discount).Error
	if err != nil {
		return discount, errorx.New("کد تخفیف یافت نشد", "model", err)
	}
	if err := utils.CheckAccess(m.Ctx, discount.UserID); err != nil {
		return nil, err
	}
	return discount, nil
}

func (m *MysqlManager) FindDiscountByCodeAndUserID(code string, userOwnerID uint64) (*Discount, error) {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:showDiscount", "model")
	defer span.End()
	discount := &Discount{}
	err := m.GetConn().Where("code = ?", code).Where("user_id = ?", userOwnerID).Find(discount).Error
	if err != nil {
		return nil, errorx.New("کد تخفیف یافت نشد", "model", err)
	}
	return discount, nil
}
