package models

import (
	"backend/internal/app/DTOs"
	"backend/internal/app/utils"
	"encoding/gob"
	"errors"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

type User struct {
	ID         uint64      `gorm:"primary_key;auto_increment" json:"id"`
	Gender     string      `json:"gender" sql:"type:ENUM('man','woman')"`
	Firstname  string      `json:"firstname"`
	Lastname   string      `json:"lastname"`
	Email      string      `json:"email"`
	ProvinceID uint32      `json:"province_id"`
	CityID     uint32      `json:"city_id"`
	Lat        string      `json:"lat"`
	Long       string      `json:"long"`
	Address    string      `json:"address"`
	Mobile     string      `json:"mobile"`
	Password   string      `json:"password"`
	ExpireAt   string      `json:"expire_at"`
	Status     string      `json:"status"`
	PostalCode string      `json:"postal_code"`
	VerifyCode uint16      `json:"verify_code"`
	CartNumber string      `json:"cart_number"`
	Shaba      string      `json:"shaba"`
	IsAdmin    byte        `json:"is_admin"`
	Financial  []Financial `gorm:"foreignKey:user_id" json:"financial"`
	UpdatedAt  string      `json:"updated_at"`
	CreatedAt  string      `json:"created_at"`
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
	manager.CreateUser(&gin.Context{}, &User{
		ID:         1,
		Gender:     "man",
		Firstname:  "امیر",
		Lastname:   "شیردلی",
		Email:      "amirex128@gmail.com",
		ProvinceID: 1,
		CityID:     2,
		Lat:        "35",
		Long:       "35",
		Address:    "",
		Mobile:     "",
		Password:   "",
		ExpireAt:   "",
		Status:     "",
		PostalCode: "",
		VerifyCode: 0,
		CartNumber: "",
		Shaba:      "",
		IsAdmin:    1,
		Financial:  nil,
		UpdatedAt:  utils.NowTime(),
		CreatedAt:  utils.NowTime(),
	})
}
func (m *MysqlManager) CreateUser(c *gin.Context, user *User) error {
	find := m.GetConn().Where("mobile = ? and password = ?", user.Mobile, utils.GeneratePasswordHash(user.Password)).Find(&User{}).RowsAffected
	if find > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "کاربری با این مشخصات قبلا ثبت شده است",
		})
		return errors.New("")
	}
	user.Password = utils.GeneratePasswordHash(user.Password)
	err := m.GetConn().Create(user).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "خطایی در فرایند ثبت نام شما رخ داده است لطفا مجدد تلاش نمایید",
			"error":   err.Error(),
		})
		return err
	}
	return nil
}
func (m *MysqlManager) FindUserByMobilePassword(user DTOs.Login) (*User, error) {
	res := &User{}
	err := m.GetConn().Where("mobile = ? and password = ?", user.Mobile, utils.GeneratePasswordHash(user.Password)).First(res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (m *MysqlManager) FindUserByID(c *gin.Context, userID uint64) (*User, error) {
	res := &User{}
	err := m.GetConn().Where("id = ?", userID).First(res).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "فروشگاه یافت نشد",
			"error":   err.Error(),
		})
		return nil, err
	}
	return res, nil
}
