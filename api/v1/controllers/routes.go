package controllers

import (
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
		ad := v1.Group("admin")
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
			post := ad.Group("post")
			{
				post.GET("/", indexPost)
				post.GET("/show/:id", showPost)
				post.POST("/create", createPost)
				post.POST("/update/:id", updatePost)
				post.POST("/delete/:id", deletePost)
			}
			category := ad.Group("category")
			{
				category.GET("/", indexCategory)
			}
			comment := ad.Group("comment")
			{
				comment.GET("/", indexComment)
				comment.POST("/create", createComment)
				comment.POST("/delete/:id", deleteComment)
				comment.POST("/approve/:id", approveComment)
			}
			tag := ad.Group("tag")
			{
				tag.GET("/", indexTag)
				tag.POST("/create", createTag)
				tag.POST("/delete/:id", deleteTag)
				tag.POST("/add", addTag)
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

			customer := user.Group("customer")
			customer.POST("/customer/create", createCustomer)
			customer.POST("/customer/login", loginCustomer)
			customer.POST("/customer/verify", verifyCustomer)
			customer.POST("/customer/update", updateCustomer)
		}

	}

}
