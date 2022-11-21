package models

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) *uint64 {
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
func GeneratePasswordHash(pass string) string {
	return Sha512EncodeSaltIter(pass, 2, "amirex128-selloora")
}
func Sha512EncodeSaltIter(raw string, iter int, salt string) string {
	salted := fmt.Sprintf("%s{%s}", raw, salt)
	digest := mSha512(salted)
	for i := 1; i < iter; i++ {
		digest = mSha512(string(digest) + salted)
	}
	return hex.EncodeToString(digest)
}

func mSha512(s string) []byte {
	h := sha512.New()
	h.Write([]byte(s))
	return h.Sum(nil)
}

func IsAdmin(c *gin.Context) bool {
	userID := GetUser(c)
	user, err := NewMysqlManager(c.Request.Context()).FindUserByID(c, nil, *userID)
	if err != nil {
		return false
	}

	return user.IsAdmin
}
