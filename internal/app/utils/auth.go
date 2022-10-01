package utils

import (
	"backend/internal/pkg/framework/hash"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) uint64 {
	return uint64(jwt.ExtractClaims(c)["id"].(float64))
}
func GeneratePasswordHash(pass string) string {
	return hash.Sha512EncodeSaltIter(pass, 2, "amirex128-selloora")
}
