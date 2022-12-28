package models

import (
	"github.com/amirex128/selloora_backend/internal/DTOs"
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"go.elastic.co/apm/v2"
)

type Domain struct {
	ID        uint64  `gorm:"primary_key;auto_increment" json:"id"`
	UserID    *uint64 `json:"user_id"`
	ShopID    *uint64 `gorm:"default:null" json:"shop_id"`
	Shop      *Shop   `json:"shop"`
	Name      string  `json:"name"`
	Type      string  `json:"type" sql:"type:ENUM('subdomain','domain')"`
	DnsStatus string  `gorm:"default:pending" json:"dns_status" sql:"type:ENUM('pending','verified','failed')"`
}

func initDomain(manager *MysqlManager) {
	if !manager.GetConn().Migrator().HasTable(&Domain{}) {
		manager.GetConn().Migrator().CreateTable(&Domain{})
		manager.CreateDomain(DTOs.CreateDomain{
			Name:   "subdomain.selloora.conf",
			ShopID: 1,
			Type:   "domain",
		})
		manager.CreateDomain(DTOs.CreateDomain{
			Name:   "selloora.test",
			ShopID: 2,
			Type:   "domain",
		})
		manager.CreateDomain(DTOs.CreateDomain{
			Name:   "amir.test",
			ShopID: 3,
			Type:   "domain",
		})
		manager.CreateDomain(DTOs.CreateDomain{
			Name:   "sell.selloora.test",
			ShopID: 4,
			Type:   "subdomain",
		})
		manager.CreateDomain(DTOs.CreateDomain{
			Name:   "amir.selloora.test",
			ShopID: 5,
			Type:   "subdomain",
		})
		manager.CreateDomain(DTOs.CreateDomain{
			Name:   "localhost:8585",
			ShopID: 6,
			Type:   "domain",
		})
	}

}

func (m *MysqlManager) CreateDomain(dto DTOs.CreateDomain) (*Domain, error) {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:CreateDomain", "model")
	defer span.End()
	domain := &Domain{
		ShopID:    &dto.ShopID,
		Name:      dto.Name,
		UserID:    utils.GetUserID(m.Ctx),
		Type:      dto.Type,
		DnsStatus: "pending",
	}
	err := m.GetConn().Create(domain).Error
	if err != nil {
		return domain, errorx.New("خطایی در سرور رخ داده است", "model", err)
	}
	return domain, nil
}

func (m *MysqlManager) FindDomainByName(name string) (*Domain, error) {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:FindDomainByName", "model")
	defer span.End()
	domain := &Domain{}
	err := m.GetConn().Where("name = ?", name).First(domain).Error
	if err != nil {
		return nil, errorx.New("دامنه یافت نشد", "model", err)
	}
	return domain, nil
}
func (m *MysqlManager) DeleteDomain(domainID uint64) error {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:showDomain", "model")
	defer span.End()
	domain := Domain{}
	err := m.GetConn().Where("id = ?", domainID).First(&domain).Error
	if err != nil {
		return errorx.New("دامنه یافت نشد", "model", err)
	}
	if err := utils.CheckAccess(m.Ctx, domain.UserID); err != nil {
		return err
	}
	err = m.GetConn().Delete(&domain).Error
	if err != nil {
		return errorx.New("خطا در حذف دامنه", "model", err)
	}
	return nil
}
func (m *MysqlManager) GetAllDomainWithPagination(dto DTOs.IndexDomain) (*DTOs.Pagination, error) {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:showDomain", "model")
	defer span.End()
	conn := m.GetConn()
	var domains []Domain
	pagination := &DTOs.Pagination{PageSize: dto.PageSize, Page: dto.Page}

	conn = conn.Scopes(DTOs.Paginate("domains", pagination, conn))
	if dto.Search != "" {
		conn = conn.Where("name LIKE ?", "%"+dto.Search+"%")
	}
	err := conn.Where("user_id = ?", utils.GetUserID(m.Ctx)).Where("shop_id = ? ", dto.ShopID).Order("id DESC").Find(&domains).Error
	if err != nil {
		return pagination, errorx.New("خطا در دریافت دامنه ها", "model", err)
	}
	pagination.Data = domains
	return pagination, nil
}

func (m *MysqlManager) FindDomainByID(domainID uint64) (*Domain, error) {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:FindDomainByID", "model")
	defer span.End()
	domain := &Domain{}
	err := m.GetConn().Where("id = ?", domainID).First(domain).Error
	if err != nil {
		return nil, errorx.New("مشکلی در یافتن دامنه پیش آمده است", "model", err)
	}
	if err := utils.CheckAccess(m.Ctx, domain.UserID); err != nil {
		return nil, err
	}
	return domain, nil
}
