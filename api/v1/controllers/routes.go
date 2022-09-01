package controllers

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) {
	r.Static("/assets", "../../../assets")
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	r.POST("/login", authMiddleware.LoginHandler)
	r.POST("/register", register)
	r.POST("/forget", forget)

	v1 := r.Group("v1")
	{
		admin := v1.Group("admin")
		admin.Use(authMiddleware.MiddlewareFunc())
		{
			product := admin.Group("products")
			{
				product.GET("/", indexProduct)
				product.GET("/show/:id", showProduct)
				product.POST("/create", createProduct)
				product.POST("/update/:id", updateProduct)
				product.POST("/delete/:id", deleteProduct)
			}

		}

		user := v1.Group("user")
		{
			user.POST("/order/create", createOrder)
			user.POST("/discount/check", checkDiscount)
			user.POST("/customer/request", requestCustomer)
			user.POST("/customer/verify", verifyCustomer)
			user.POST("/customer/update", updateCustomer)
		}

	}

}
