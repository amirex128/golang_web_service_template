package models

import (
	"backend/internal/pkg/framework/hash"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) uint64 {
	var userID uint64
	defer func() {
		if r := recover(); r != nil {
			userID = 0
		}
	}()
	userIDString := jwt.ExtractClaims(c)["id"]
	if userIDString != "" && userIDString != nil {
		userID = uint64(userIDString.(float64))
	} else {
		userID = 0
	}

	return userID
}
func GeneratePasswordHash(pass string) string {
	return hash.Sha512EncodeSaltIter(pass, 2, "amirex128-selloora")
}

func IsAdmin(c *gin.Context) bool {
	userID := GetUser(c)
	user, err := NewMainManager().FindUserByID(c, nil, userID)
	if err != nil {
		return false
	}

	return user.IsAdmin
}
