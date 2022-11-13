package models

import (
	"backend/internal/app/DTOs"
	"context"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

type Domain struct {
	ID        uint64  `gorm:"primary_key;auto_increment" json:"id"`
	ShopID    *uint64 `gorm:"default:null" json:"shop_id"`
	Shop      *Shop   `json:"shop"`
	Name      string  `json:"name"`
	Type      string  `json:"type" sql:"type:ENUM('subdomain','domain')"`
	DnsStatus string  `gorm:"default:pending" json:"dns_status" sql:"type:ENUM('pending','verified','failed')"`
}

func initDomain(manager *MysqlManager) {
	manager.GetConn().AutoMigrate(&Domain{})
	manager.CreateDomain(&gin.Context{}, context.Background(), DTOs.CreateDomain{
		Name:   "localhost:8585",
		ShopID: 1,
		Type:   "domain",
	})
	manager.CreateDomain(&gin.Context{}, context.Background(), DTOs.CreateDomain{
		Name:   "selloora.test",
		ShopID: 2,
		Type:   "domain",
	})
	manager.CreateDomain(&gin.Context{}, context.Background(), DTOs.CreateDomain{
		Name:   "amir.test",
		ShopID: 3,
		Type:   "domain",
	})
	manager.CreateDomain(&gin.Context{}, context.Background(), DTOs.CreateDomain{
		Name:   "sell.selloora.test",
		ShopID: 4,
		Type:   "subdomain",
	})
	manager.CreateDomain(&gin.Context{}, context.Background(), DTOs.CreateDomain{
		Name:   "amir.selloora.test",
		ShopID: 5,
		Type:   "subdomain",
	})
}

func (m *MysqlManager) CreateDomain(c *gin.Context, ctx context.Context, dto DTOs.CreateDomain) error {
	span, ctx := apm.StartSpan(ctx, "CreateDomain", "model")
	defer span.End()
	err := m.GetConn().Create(&Domain{
		ShopID:    &dto.ShopID,
		Name:      dto.Name,
		Type:      dto.Type,
		DnsStatus: "pending",
	}).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "خطایی در سرور رخ داده است",
			"error":   err.Error(),
			"type":    "model",
		})
		return err
	}
	return nil
}

func (m *MysqlManager) FindDomainByName(c *gin.Context, ctx context.Context, name string) (*Domain, error) {
	span, ctx := apm.StartSpan(ctx, "FindDomainByName", "model")
	defer span.End()
	domain := &Domain{}
	err := m.GetConn().Where("name = ?", name).First(domain).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "دامنه یافت نشد",
			"error":   err.Error(),
			"type":    "model",
		})
		return domain, err
	}
	return domain, nil
}
func (m *MysqlManager) DeleteDomain(c *gin.Context, ctx context.Context, domainID uint64) error {
	span, ctx := apm.StartSpan(ctx, "showDomain", "model")
	defer span.End()
	domain := Domain{}
	err := m.GetConn().Where("id = ?", domainID).First(&domain).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "دامنه یافت نشد",
			"error":   err.Error(),
			"type":    "model",
		})
		return err
	}

	err = m.GetConn().Delete(&domain).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "خطا در حذف دامنه",
			"error":   err.Error(),
			"type":    "model",
		})
		return err
	}
	return nil
}
func (m *MysqlManager) GetAllDomainWithPagination(c *gin.Context, ctx context.Context, dto DTOs.IndexDomain) (*DTOs.Pagination, error) {
	span, ctx := apm.StartSpan(ctx, "showDomain", "model")
	defer span.End()
	conn := m.GetConn()
	var domains []Domain
	pagination := &DTOs.Pagination{PageSize: dto.PageSize, Page: dto.Page}

	conn = conn.Scopes(DTOs.Paginate("domains", pagination, conn))
	if dto.Search != "" {
		conn = conn.Where("name LIKE ?", "%"+dto.Search+"%").Where("shop_id = ? ", dto.ShopID).Order("id DESC")
	}
	err := conn.Find(&domains).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "خطا در دریافت دامنه ها",
			"error":   err.Error(),
			"type":    "model",
		})
		return pagination, err
	}
	pagination.Data = domains
	return pagination, nil
}
