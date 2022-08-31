package controllers

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
)

var (
	validate *validator.Validate
)

func Routes(r *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) {
	validate = validator.New()

	r.Static("/assets", "../../../assets")
	r.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})
	r.GET("/", func(c *gin.Context) {
		c.JSONP(http.StatusOK, gin.H{"message": "pong"})
	})
	r.POST("/login", authMiddleware.LoginHandler)
	r.POST("/register", register)

	v1 := r.Group("/v1")
	v1.Use(authMiddleware.MiddlewareFunc())
	{
		v1.GET("/refresh_token", authMiddleware.RefreshHandler)

	}

}
