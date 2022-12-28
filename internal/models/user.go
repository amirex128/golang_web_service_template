package models

import (
	"github.com/amirex128/selloora_backend/internal/DTOs"
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/brianvoe/gofakeit/v6"
	"go.elastic.co/apm/v2"
)

type User struct {
	ID            uint64      `gorm:"primary_key;auto_increment" json:"id"  fake:"{custom_uint64:0}"`
	Gender        string      `json:"gender" sql:"type:ENUM('man','woman')" fake:"{custom_enum:man,woman}"`
	Firstname     string      `json:"firstname" fake:"{firstname}"`
	Lastname      string      `json:"lastname" fake:"{lastname}"`
	Email         string      `json:"email" fake:"{email}"`
	Mobile        string      `json:"mobile" fake:"{phone}"`
	ExpireAt      string      `json:"expire_at" fake:"{date}"`
	Status        string      `json:"status" sql:"type:ENUM('ok','block')" fake:"{custom_enum:ok,block}"`
	VerifyCode    string      `json:"verify_code" fake:"{number:1000,9999}"`
	CartNumber    string      `json:"cart_number"`
	Shaba         string      `json:"shaba"`
	IsAdmin       bool        `json:"is_admin" fake:"{bool}"`
	Financial     []Financial `gorm:"foreignKey:user_id" json:"financial" fake:"skip"`
	LastSendSMSAt string      `json:"last_send_sms_at" fake:"{date}"`
	Password      string      `json:"password" fake:"{password:true,false,false,true,false,10}"`
	UpdatedAt     string      `json:"updated_at" fake:"{date}"`
	CreatedAt     string      `json:"created_at" fake:"{date}"`
}

func initUser(manager *MysqlManager) {

	if !manager.GetConn().Migrator().HasTable(&User{}) {
		manager.GetConn().Migrator().CreateTable(&User{})

		for i := 0; i < 100; i++ {
			model := new(User)
			gofakeit.Struct(model)

			manager.CreateUser(model)
		}
		model := new(User)
		gofakeit.Struct(model)
		model.ID = 101
		manager.CreateUser(model)

	}
}

func (m *MysqlManager) CreateUser(user *User) error {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:CreateUser", "model")
	defer span.End()

	find := m.GetConn().Where("mobile = ?", user.Mobile).Find(&User{}).RowsAffected
	if find > 0 {
		return errorx.New("کاربری با این مشخصات قبلا ثبت شده است", "model", nil)
	}
	err := m.GetConn().Create(user).Error
	if err != nil {
		return errorx.New("خطایی در فرایند ثبت نام شما رخ داده است لطفا مجدد تلاش نمایید", "model", err)
	}
	return nil
}

func (m *MysqlManager) FindUserByMobileAndCodeVerify(user DTOs.Verify) (*User, error) {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:FindUserByMobileAndCodeVerify", "model")
	defer span.End()
	res := &User{}
	err := m.GetConn().Where("mobile = ? and verify_code = ?", user.Mobile, user.VerifyCode).First(res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (m *MysqlManager) FindUserByMobileAndPassword(user DTOs.Verify) (*User, error) {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:FindUserByMobileAndPassword", "model")
	defer span.End()
	res := &User{}
	err := m.GetConn().Where("mobile = ? and password = ?", user.Mobile, GeneratePasswordHash(user.Password)).First(res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (m *MysqlManager) FindUserByMobile(mobile string) (*User, error) {

	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:FindUserByMobile", "model")
	defer span.End()

	res := &User{}
	err := m.GetConn().Where("mobile = ?", mobile).First(res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (m *MysqlManager) FindUserByID(userID uint64) (*User, error) {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:FindUserByID", "model")
	defer span.End()
	res := &User{}
	err := m.GetConn().Where("id = ?", userID).First(res).Error
	if err != nil {
		return nil, errorx.New("کاربر یافت نشد", "model", err)
	}
	return res, nil
}

func (m *MysqlManager) UpdateUser(user *User) error {
	span, _ := apm.StartSpan(m.Ctx.Request.Context(), "model:UpdateUser", "model")
	defer span.End()

	var newUser User
	err := m.GetConn().Where("id = ?", user.ID).First(&newUser).Error
	if err != nil {
		return errorx.New("کاربر یافت نشد", "model", err)
	}
	if user.Gender != "" {
		newUser.Gender = user.Gender
	}
	if user.Firstname != "" {
		newUser.Firstname = user.Firstname
	}
	if user.Lastname != "" {
		newUser.Lastname = user.Lastname
	}
	if user.Email != "" {
		newUser.Email = user.Email
	}
	if user.Mobile != "" {
		newUser.Mobile = user.Mobile
	}
	if user.ExpireAt != "" {
		newUser.ExpireAt = user.ExpireAt
	}
	if user.Status != "" {
		newUser.Status = user.Status
	}
	if user.VerifyCode != "" {
		newUser.VerifyCode = user.VerifyCode
	}
	if user.CartNumber != "" {
		newUser.CartNumber = user.CartNumber
	}
	if user.Shaba != "" {
		newUser.Shaba = user.Shaba
	}
	if user.LastSendSMSAt != "" {
		newUser.LastSendSMSAt = user.LastSendSMSAt
	}
	if user.Password != "" {
		newUser.Password = user.Password
	}
	newUser.UpdatedAt = utils.NowTime()

	err = m.GetConn().Save(&newUser).Error
	if err != nil {
		return errorx.New("خطایی در فرایند ویرایش شما رخ داده است لطفا مجدد تلاش نمایید", "model", err)
	}
	return nil
}
