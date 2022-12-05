package models

import (
	"github.com/amirex128/selloora_backend/internal/DTOs"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/brianvoe/gofakeit/v6"
	"go.elastic.co/apm/v2"
)

type Address struct {
	ID         uint64   `gorm:"primary_key;auto_increment" json:"id"`
	UserID     *uint64  `gorm:"default:null" json:"user_id"`
	Title      string   `json:"title"`
	ProvinceID uint64   `json:"province_id"`
	Province   Province `gorm:"foreignKey:province_id" json:"province"`
	CityID     uint64   `json:"city_id"`
	City       City     `gorm:"foreignKey:city_id" json:"city"`
	Address    string   `json:"address"`
	PostalCode string   `json:"postal_code"`
	Mobile     string   `json:"mobile"`
	FullName   string   `json:"full_name"`
	Lat        string   `json:"lat"`
	Long       string   `json:"long"`
}

func initAddress(manager *MysqlManager) {
	if !manager.GetConn().Migrator().HasTable(&Address{}) {
		manager.GetConn().Migrator().CreateTable(&Address{})
		for i := 0; i < 100; i++ {
			model := new(DTOs.CreateAddress)
			gofakeit.Struct(model)

			manager.CreateAddress(*model)
		}
	}

}

func (m *MysqlManager) CreateAddress(dto DTOs.CreateAddress) (*Address, error) {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:CreateAddress", "model")
	defer span.End()
	userID := GetUser(m.Ctx)
	address := &Address{
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
	err := m.GetConn().Create(address).Error
	if err != nil {
		return address, errorx.New("خطایی در ثبت آدرس رخ داده است", "model:panic", err)
	}
	return address, nil
}

func (m *MysqlManager) UpdateAddress(dto DTOs.UpdateAddress) error {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:UpdateAddress", "model")
	defer span.End()
	address := Address{}
	err := m.GetConn().First(&address, dto.ID).Error
	if err != nil {
		return errorx.New("خطایی در بروزرسانی آدرس رخ داده است", "model", err)
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
		return errorx.New("خطایی در بروزرسانی آدرس رخ داده است", "model", err)
	}
	return err
}

func (m *MysqlManager) DeleteAddress(addressID uint64) error {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:DeleteAddress", "model")
	defer span.End()
	address := Address{}
	err := m.GetConn().First(&address, addressID).Error
	if err != nil {
		return errorx.New("خطایی در حذف آدرس رخ داده است", "model", err)
	}
	err = m.GetConn().Delete(&address).Error
	if err != nil {
		return errorx.New("خطایی در حذف آدرس رخ داده است", "model", err)
	}
	return err
}

func (m *MysqlManager) IndexAddress(userID uint64) ([]*Address, error) {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:IndexAddress", "model")
	defer span.End()
	var addresses []*Address
	err := m.GetConn().Where("user_id = ?", userID).Find(&addresses).Error
	if err != nil {
		return nil, errorx.New("خطایی در دریافت آدرس ها رخ داده است", "model", err)
	}
	return addresses, nil
}

func (m *MysqlManager) GetAllAddressWithPagination(dto DTOs.IndexAddress) (*DTOs.Pagination, error) {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:IndexAddress", "model")
	defer span.End()
	conn := m.GetConn()
	var addresses []Address
	pagination := &DTOs.Pagination{PageSize: dto.PageSize, Page: dto.Page}

	userID := GetUser(m.Ctx)
	conn = conn.Scopes(DTOs.Paginate("addresses", pagination, conn))
	if dto.Search != "" {
		conn = conn.Where("title LIKE ?", "%"+dto.Search+"%").Where("user_id = ? ", userID).Order("id DESC")
	}
	err := conn.Find(&addresses).Error
	if err != nil {
		return nil, errorx.New("خطایی در دریافت آدرس ها رخ داده است", "model", err)
	}
	pagination.Data = addresses
	return pagination, nil
}

func (m *MysqlManager) FindAddressByID(addressID uint64) (*Address, error) {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:FindAddressByID", "model")
	defer span.End()
	address := &Address{}
	err := m.GetConn().Where("id = ?", addressID).First(address).Error
	if err != nil {
		return nil, errorx.New("مشکلی در یافتن آدرس پیش آمده است", "model", err)
	}
	return address, nil
}
