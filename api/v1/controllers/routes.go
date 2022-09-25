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
			discount := admin.Group("discounts")
			{
				discount.GET("/", indexDiscount)
				discount.GET("/show/:id", showDiscount)
				discount.POST("/create", createDiscount)
				discount.POST("/update/:id", updateDiscount)
				discount.POST("/delete/:id", deleteDiscount)
			}
		}

		user := v1.Group("user")
		{
			user.POST("/order/create", createOrder)
			user.POST("/discount/check", checkDiscount)
			user.POST("/customer/create", createCustomer)
			user.POST("/customer/login", loginCustomer)
			user.POST("/customer/verify", verifyCustomer)
			user.POST("/customer/update", updateCustomer)
		}

	}

}
