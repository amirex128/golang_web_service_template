package controllers

import (
	"backend/api/v1/controllers/admin"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) {
	//r.LoadHTMLGlob("../../templates/*")
	r.Static("/assets", "../../assets")
	r.NoRoute(func(c *gin.Context) {
		c.Set("template", "404-error.html")
	})

	root := r.Group("/")
	{
		root.GET("/", indexLanding)
		root.GET("blog", blogLanding)
		root.GET("blog/:slug", detailsLanding)
		root.GET("contact", contactLanding)
		root.GET("faq", faqLanding)
		root.GET("pricing", pricingLanding)
		root.GET("services", servicesLanding)
		root.GET("testimonial", testimonialLanding)
		root.GET("learn", learnLanding)
		root.GET("rules", rulesLanding)
		root.GET("return-rules", returnRulesLanding)
	}

	v1 := r.Group("api/v1")
	{
		ad := v1.Group("ad")
		ad.Use(authMiddleware.MiddlewareFunc())
		{
			product := ad.Group("products")
			{
				product.GET("/", indexProduct)
				product.GET("/show/:id", showProduct)
				product.POST("/create", createProduct)
				product.POST("/update/:id", updateProduct)
				product.POST("/delete/:id", deleteProduct)
			}
			discount := ad.Group("discounts")
			{
				discount.GET("/", indexDiscount)
				discount.GET("/show/:id", showDiscount)
				discount.POST("/create", createDiscount)
				discount.POST("/update/:id", updateDiscount)
				discount.POST("/delete/:id", deleteDiscount)
			}
			blog := ad.Group("blog")
			{
				blog.GET("/", admin.IndexPost)
				blog.GET("/show/:id", admin.ShowPost)
				blog.POST("/create", admin.CreatePost)
				blog.POST("/update/:id", admin.UpdatePost)
				blog.POST("/delete/:id", admin.DeletePost)
			}
		}

		user := v1.Group("user")
		{
			r.POST("/login", authMiddleware.LoginHandler)
			r.POST("/register", register)
			r.POST("/forget", forget)
			r.GET("/sadad/pay", sadadPaymentRequest)
			r.POST("/sadad/verify", sadadPaymentVerify)

			user.POST("/order/create", createOrder)
			user.POST("/discount/check", checkDiscount)
			user.POST("/customer/create", createCustomer)
			user.POST("/customer/login", loginCustomer)
			user.POST("/customer/verify", verifyCustomer)
			user.POST("/customer/update", updateCustomer)
		}

	}

}
