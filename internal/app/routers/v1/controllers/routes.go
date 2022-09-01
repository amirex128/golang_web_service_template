package controllers

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"log"
)

func Routes(r *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) {
	r.Static("/assets", "../../../assets")
	r.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	r.POST("/login", authMiddleware.LoginHandler)
	r.POST("/register", register)

	v1 := r.Group("v1")
	v1.Use(authMiddleware.MiddlewareFunc())
	{
		v1.GET("/refresh_token", authMiddleware.RefreshHandler)

		product := v1.Group("/products")
		product.GET("/", indexProduct)
		product.GET("/show/:id", showProduct)
		product.POST("/create", createProduct)
		product.POST("/update", updateProduct)
		product.POST("/delete", deleteProduct)

	}

}
