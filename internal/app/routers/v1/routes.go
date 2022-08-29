package v1

import (
	"backend/internal/app/helpers"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Routes(r *gin.Engine) {
	r.Static("/assets", "../../../assets")
	authMiddleware := helpers.GetAuthMiddleware()
	r.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})
	r.GET("/", func(c *gin.Context) {
		c.JSONP(http.StatusOK, gin.H{"message": "pong"})
	})
	r.POST("/login", authMiddleware.LoginHandler)

	v1 := r.Group("/v1")
	v1.GET("/refresh_token", authMiddleware.RefreshHandler)
	v1.Use(authMiddleware.MiddlewareFunc())
	{
		v1.GET("/ping", func(c *gin.Context) {
			value, _ := c.Get("JWT_PAYLOAD")
			c.JSON(http.StatusOK, gin.H{
				"message": value,
			})
		})
	}

}
