package utils

import (
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func getUser(c *gin.Context) *uint64 {
	var userID *uint64
	defer func() {
		if r := recover(); r != nil {
			userID = nil
		}
	}()
	userIDString := jwt.ExtractClaims(c)["id"]
	if userIDString != "" && userIDString != nil {
		u := uint64(userIDString.(float64))
		userID = &u
	} else {
		userID = nil
	}

	return userID
}
func IsAdmin(c *gin.Context) bool {
	claims := jwt.ExtractClaims(c)
	if claims["is_admin"] == true {
		return true
	}
	return false
}

func CheckAccess(c *gin.Context, userID *uint64) error {
	id := getUser(c)
	if IsAdmin(c) {
		return nil
	}
	if id == nil {
		return errorx.New("شما سطح دسترسی کافی ندارید", "authorize", nil)
	}
	if userID == nil {
		return errorx.New("شما سطح دسترسی کافی ندارید", "authorize", nil)
	}
	if *id != *userID {
		return errorx.New("شما سطح دسترسی کافی ندارید", "authorize", nil)
	}

	return nil
}
