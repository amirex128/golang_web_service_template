package models

import (
	"backend/internal/app/DTOs"
	"backend/internal/app/utils"
	"encoding/gob"
	"io"
)

type User struct {
	ID          int64  `json:"id"`
	Gender      string `json:"gender"`
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Email       string `json:"email"`
	ShopName    string `json:"shop_name"`
	GuildID     string `json:"guild_id"`
	ProvinceID  string `json:"province_id"`
	CityID      string `json:"city_id"`
	Lat         string `json:"lat"`
	Long        string `json:"long"`
	Logo        string `json:"logo"`
	Address     string `json:"address"`
	Mobile      string `json:"mobile"`
	Password    string `json:"password"`
	Phone       string `json:"phone"`
	Description string `json:"desc"`
	ExpireAt    string `json:"expire_at"`
	Status      string `json:"status"`
	PostalCode  string `json:"postal_code"`
	VerifyCode  uint16 `json:"verify_code"`
	UpdatedAt   string `json:"updated_at"`
	CreatedAt   string `json:"created_at"`
}
type UserArr []User

func (s UserArr) Len() int {
	return len(s)
}
func (s UserArr) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s UserArr) Less(i, j int) bool {
	return s[i].ID < s[j].ID
}

func (c *User) Encode(iw io.Writer) error {
	return gob.NewEncoder(iw).Encode(c)
}

func (c *User) Decode(ir io.Reader) error {
	return gob.NewDecoder(ir).Decode(c)
}
func initUser(manager *MysqlManager) {
	manager.GetConn().AutoMigrate(&User{})
}
func (m *MysqlManager) CreateUser(user *User) string {
	find := m.GetConn().Where("mobile = ? and password = ?", user.Mobile, utils.GeneratePasswordHash(user.Password)).Find(&User{}).RowsAffected
	if find > 0 {
		return "کاربری با این مشخصات قبلا ثبت شده است"
	}
	user.Password = utils.GeneratePasswordHash(user.Password)
	err := m.GetConn().Create(user).Error
	if err != nil {
		return "خطایی در فرایند ثبت نام شما رخ داده است لطفا مجدد تلاش نمایید"
	}
	return ""
}
func (m *MysqlManager) FindUserByMobilePassword(user DTOs.Login) (*User, error) {
	res := &User{}
	err := m.GetConn().Where("mobile = ? and password = ?", user.Mobile, utils.GeneratePasswordHash(user.Password)).First(res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}
