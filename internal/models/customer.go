package models

import (
	"github.com/amirex128/selloora_backend/internal/DTOs"
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/brianvoe/gofakeit/v6"
	"go.elastic.co/apm/v2"
)

type Customer struct {
	ID            uint64 `gorm:"primary_key;auto_increment" json:"id"`
	FullName      string `json:"full_name"`
	Mobile        string `json:"mobile"`
	ProvinceID    uint32 `json:"province_id"`
	CityID        uint32 `json:"city_id"`
	Address       string `json:"address"`
	PostalCode    string `json:"postal_code"`
	VerifyCode    string `json:"verify_code"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
	LastSendSMSAt string `json:"last_send_sms_at"`
}

func initCustomer(manager *MysqlManager) {
	if !manager.GetConn().Migrator().HasTable(&Customer{}) {
		manager.GetConn().Migrator().CreateTable(&Customer{})
		for i := 0; i < 100; i++ {
			model := new(DTOs.CreateUpdateCustomer)
			gofakeit.Struct(model)

			manager.CreateCustomer(*model)
		}

	}

}

func (m *MysqlManager) FindCustomerById(customerID uint64) (*Customer, error) {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:FindCustomerById", "model")
	defer span.End()
	customer := &Customer{}
	err := m.GetConn().Where("id = ?", customerID).First(customer).Error
	if err != nil {
		return nil, errorx.New("مشتری یافت نشد", "model", err)
	}
	return customer, nil
}

func (m *MysqlManager) FindCustomerByMobile(mobile string) (*Customer, error) {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:FindCustomerByMobile", "model")
	defer span.End()
	customer := &Customer{}
	err := m.GetConn().Where("mobile = ?", mobile).Find(customer).Error
	if err != nil {
		return nil, errorx.New("شماره موبایل تکراری یا اشتباه میباشد", "model", err)
	}
	return customer, nil
}

func (m *MysqlManager) FindCustomerByMobileAndVerifyCode(mobile, verifyCode string) (*Customer, error) {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:FindCustomerByMobileAndVerifyCode", "model")
	defer span.End()
	customer := &Customer{}
	err := m.GetConn().Where("mobile = ?", mobile).Where("verify_code = ?", verifyCode).First(customer).Error
	if err != nil {
		return nil, errorx.New("رمز عبور یا کد تایید اشتباه است", "model", err)
	}
	return customer, nil
}

func (m *MysqlManager) CreateCustomer(dto DTOs.CreateUpdateCustomer) error {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:CreateCustomer", "model")
	defer span.End()
	customer := Customer{
		FullName:   dto.FullName,
		Mobile:     dto.Mobile,
		ProvinceID: dto.ProvinceID,
		CityID:     dto.CityID,
		Address:    dto.Address,
		PostalCode: dto.PostalCode,
		VerifyCode: dto.VerifyCode,
	}
	err := m.GetConn().Create(&customer).Error
	if err != nil {
		return errorx.New("مشکلی در ثبت نام شما پیش آمده است", "model", err)
	}
	return nil
}

func (m *MysqlManager) CreateCodeCustomer(dto DTOs.RequestCreateLoginCustomer, encryptPassword string) (*Customer, error) {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:CreateCodeCustomer", "model")
	defer span.End()
	rowsAffected := m.GetConn().Where("mobile = ?", dto.Mobile).First(&Customer{}).RowsAffected
	if rowsAffected > 0 {
		return nil, errorx.New("شماره موبایل قبلا ثبت شده است", "model", nil)
	}

	customer := &Customer{
		Mobile:        dto.Mobile,
		VerifyCode:    encryptPassword,
		LastSendSMSAt: utils.NowTime(),
	}
	err := m.GetConn().Create(customer).Error
	if err != nil {
		return customer, errorx.New("مشکلی در ثبت نام شما پیش آمده است", "model", err)
	}
	return customer, nil
}

func (m *MysqlManager) UpdateCustomer(dto DTOs.CreateUpdateCustomer) (*Customer, error) {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:UpdateCustomer", "model")
	defer span.End()
	customer, err := m.FindCustomerByMobile(dto.Mobile)
	if err != nil {
		return nil, err
	}

	if dto.FullName != "" {
		customer.FullName = dto.FullName
	}
	if dto.ProvinceID != 0 {
		customer.ProvinceID = dto.ProvinceID
	}
	if dto.CityID != 0 {
		customer.CityID = dto.CityID
	}
	if dto.Address != "" {
		customer.Address = dto.Address
	}
	if dto.PostalCode != "" {
		customer.PostalCode = dto.PostalCode
	}
	if dto.VerifyCode != "" {
		customer.VerifyCode = dto.VerifyCode
	}
	err = m.GetConn().Save(&customer).Error
	if err != nil {
		return nil, errorx.New("مشکلی در ویرایش اطلاعات شما پیش آمده است", "model", err)
	}
	return customer, nil
}

func (m *MysqlManager) DeleteCustomer(customerID uint64) error {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:showCustomer", "model")
	defer span.End()
	customer := Customer{}
	err := m.GetConn().Where("id = ?", customerID).First(&customer).Error
	if err != nil {
		return errorx.New("مشتری یافت نشد", "model", err)
	}

	err = m.GetConn().Delete(&customer).Error
	if err != nil {
		return errorx.New("خطا در حذف مشتری", "model", err)
	}
	return nil
}

func (m *MysqlManager) GetAllCustomerWithPagination(dto DTOs.IndexCustomer) (*DTOs.Pagination, error) {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:showCustomer", "model")
	defer span.End()
	conn := m.GetConn()
	var customers []Customer
	pagination := &DTOs.Pagination{PageSize: dto.PageSize, Page: dto.Page}

	conn = conn.Scopes(DTOs.Paginate("customers", pagination, conn))
	if dto.Search != "" {
		conn = conn.Where("title LIKE ?", "%"+dto.Search+"%").Where("shop_id = ? ", dto.ShopID).Order("id DESC")
	}
	err := conn.Find(&customers).Error
	if err != nil {
		return nil, errorx.New("خطا در دریافت مشتری ها", "model", err)
	}
	pagination.Data = customers
	return pagination, nil
}

func (m *MysqlManager) FindCustomerByID(customerID uint64) (*Customer, error) {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:FindCustomerByID", "model")
	defer span.End()
	customer := &Customer{}
	err := m.GetConn().Where("id = ?", customerID).First(customer).Error
	if err != nil {
		return nil, errorx.New("مشکلی در یافتن مشتری پیش آمده است", "model", err)
	}
	return customer, nil
}
