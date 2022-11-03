package models

import (
	"backend/internal/app/DTOs"
	"backend/internal/app/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	ID            uint64      `gorm:"primary_key;auto_increment" json:"id"`
	Gender        string      `json:"gender" sql:"type:ENUM('man','woman')"`
	Firstname     string      `json:"firstname"`
	Lastname      string      `json:"lastname"`
	Email         string      `json:"email"`
	Mobile        string      `json:"mobile"`
	ExpireAt      string      `json:"expire_at"`
	Status        string      `json:"status" sql:"type:ENUM('ok','block')"`
	VerifyCode    string      `json:"verify_code"`
	CartNumber    string      `json:"cart_number"`
	Shaba         string      `json:"shaba"`
	IsAdmin       bool        `json:"is_admin"`
	Financial     []Financial `gorm:"foreignKey:user_id" json:"financial"`
	Address       Address     `gorm:"foreignKey:user_id" json:"address"`
	LastSendSMSAt string      `json:"last_send_sms_at"`
	Password      string      `json:"password"`
	GalleryID     uint64      `json:"gallery_id"`
	Gallery       Gallery     `gorm:"foreignKey:gallery_id" json:"gallery"`
	UpdatedAt     string      `json:"updated_at"`
	CreatedAt     string      `json:"created_at"`
}

func initUser(manager *MysqlManager) {
	manager.GetConn().AutoMigrate(&User{})
	manager.CreateUser(&gin.Context{}, &User{
		ID:         1,
		Gender:     "man",
		Firstname:  "امیر",
		Lastname:   "شیردلی",
		Email:      "amirex128@gmail.com",
		Mobile:     "09024809750",
		ExpireAt:   "",
		Status:     "",
		VerifyCode: "",
		CartNumber: "",
		Shaba:      "",
		IsAdmin:    true,
		Financial:  nil,
		GalleryID:  1,
		UpdatedAt:  utils.NowTime(),
		CreatedAt:  utils.NowTime(),
	})
}

func (m *MysqlManager) CreateUser(c *gin.Context, user *User) error {
	find := m.GetConn().Where("mobile = ?", user.Mobile).Find(&User{}).RowsAffected
	if find > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "کاربری با این مشخصات قبلا ثبت شده است",
			"error":   errors.New("کاربری با این مشخصات قبلا ثبت شده است"),
		})
		return errors.New("")
	}
	err := m.GetConn().Create(user).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "خطایی در فرایند ثبت نام شما رخ داده است لطفا مجدد تلاش نمایید",
			"error":   err.Error(),
			"type":    "model",
		})
		return err
	}
	return nil
}

func (m *MysqlManager) FindUserByMobileAndCodeVerify(user DTOs.Verify) (*User, error) {
	res := &User{}
	err := m.GetConn().Where("mobile = ? and verify_code = ?", user.Mobile, user.VerifyCode).First(res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (m *MysqlManager) FindUserByMobileAndPassword(user DTOs.Verify) (*User, error) {
	res := &User{}
	err := m.GetConn().Where("mobile = ? and password = ?", user.Mobile, GeneratePasswordHash(user.Password)).First(res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (m *MysqlManager) FindUserByMobile(mobile string) (*User, error) {
	res := &User{}
	err := m.GetConn().Where("mobile = ?", mobile).First(res).Error
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
			"type":    "model",
		})
		return nil, err
	}
	return res, nil
}

func (m *MysqlManager) UpdateUser(c *gin.Context, user *User) error {
	var newUser User
	err := m.GetConn().Where("id = ?", user.ID).First(&newUser).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "کاربر یافت نشد",
			"error":   err.Error(),
			"type":    "model",
		})
		return err
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
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "خطایی در فرایند ویرایش شما رخ داده است لطفا مجدد تلاش نمایید",
			"error":   err.Error(),
			"type":    "model",
		})
		return err
	}
	return nil
}
