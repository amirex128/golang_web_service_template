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
		root.GET("category/:id", categoryLanding)
		root.GET("tag/:slug", tagLanding)
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

	ad := v1.Group("admin")
	ad.Use(authMiddleware.MiddlewareFunc())
	{
		user := ad.Group("user")
		{
			user.POST("/update/:id", updateProfile)
		}
		product := ad.Group("product")
		{
			product.GET("/", indexProduct)
			product.GET("/show/:id", showProduct)
			product.POST("/create", createProduct)
			product.POST("/update/:id", updateProduct)
			product.POST("/delete/:id", deleteProduct)
		}
		discount := ad.Group("discount")
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
		address := ad.Group("address")
		{
			address.GET("/", indexAddress)
			address.POST("/create", createAddress)
			address.POST("/update/:id", updateAddress)
			address.POST("/delete/:id", deleteAddress)
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
		order := ad.Group("order")
		{
			order.POST("/send", sendOrder)
			order.GET("/", indexOrder)
			order.POST("/approve/:id", approveOrder)
			order.POST("/cancel/:id", cancelOrder)
			order.POST("/calculate", calculateSendPrice)
			order.POST("/returned", returnedOrder)
			order.POST("/returned/accept", acceptReturnedOrder)
			order.GET("/show/:id", showOrder)
			order.GET("/tracking/:id", trackingOrder)
		}
	}

	user := v1.Group("user")
	{
		user.POST("/verify", authMiddleware.LoginHandler)
		user.POST("/login/register", register)
		user.POST("/sadad/verify", sadadPaymentVerify)

		user.POST("/order/create", createOrder)
		user.POST("/discount/check", checkDiscount)

		customer := user.Group("customer")
		{
			customer.POST("login/register", requestCreateLoginCustomer)
			customer.POST("verify", verifyCreateLoginCustomer)
			customer.POST("orders", indexCustomerOrders)

		}

	}

}
