package controllers

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) {
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
	v1.POST("/verify", authMiddleware.LoginHandler)
	v1.POST("/login/register", registerLogin)

	user := v1.Group("user")
	user.Use(authMiddleware.MiddlewareFunc())
	{
		profile := user.Group("profile")
		{
			profile.POST("/update/:id", updateProfile)
			profile.POST("/change-password", changePassword)

		}
		product := user.Group("product")
		{
			product.GET("/list", indexProduct)
			product.GET("/show/:id", showProduct)
			product.POST("/create", createProduct)
			product.POST("/update/:id", updateProduct)
			product.POST("/delete/:id", deleteProduct)
		}
		gallery := user.Group("gallery")
		{
			gallery.POST("/create", createGallery)
		}
		discount := user.Group("discount")
		{
			discount.GET("/list", indexDiscount)
			discount.GET("/show/:id", showDiscount)
			discount.POST("/create", createDiscount)
			discount.POST("/update/:id", updateDiscount)
			discount.POST("/delete/:id", deleteDiscount)
		}
		post := user.Group("post")
		{
			post.GET("/list", indexPost)
			post.GET("/show/:id", showPost)
			post.POST("/create", createPost)
			post.POST("/update/:id", updatePost)
			post.POST("/delete/:id", deletePost)
		}
		address := user.Group("address")
		{
			address.GET("/list", indexAddress)
			address.POST("/create", createAddress)
			address.POST("/update/:id", updateAddress)
			address.POST("/delete/:id", deleteAddress)
		}
		category := user.Group("category")
		{
			category.GET("/list", indexCategory)
		}
		comment := user.Group("comment")
		{
			comment.GET("/list", indexComment)
			comment.POST("/create", createComment)
			comment.POST("/delete/:id", deleteComment)
			comment.POST("/approve/:id", approveComment)
		}
		tag := user.Group("tag")
		{
			tag.GET("/list", indexTag)
			tag.POST("/create", createTag)
			tag.POST("/delete/:id", deleteTag)
			tag.POST("/add", addTag)
		}
		order := user.Group("order")
		{
			order.POST("/send", sendOrder)
			order.GET("/list", indexOrder)
			order.POST("/approve/:id", approveOrder)
			order.POST("/cancel/:id", cancelOrder)
			order.POST("/calculate", calculateSendPrice)
			order.POST("/returned", returnedOrder)
			order.POST("/returned/accept", acceptReturnedOrder)
			order.GET("/show/:id", showOrder)
			order.GET("/tracking/:id", trackingOrder)
		}
		shop := user.Group("shop")
		{
			shop.GET("list", indexShop)
			shop.POST("/create", createShop)
			shop.POST("/update/:id", updateShop)
			shop.POST("/delete/:id", deleteShop)
			shop.POST("/check", checkSocial)
			shop.POST("/send-price", sendPrice)
			shop.GET("/instagram", getInstagramPost)
		}
	}

	admin := v1.Group("admin")
	admin.Use(authMiddleware.MiddlewareFunc())
	{

	}

	customer := v1.Group("customer")
	{
		customer.POST("login/register", requestCreateLoginCustomer)
		customer.POST("verify", verifyCreateLoginCustomer)
		customer.POST("orders", indexCustomerOrders)
		customer.POST("/sadad/verify", sadadPaymentVerify)
		customer.POST("/order/create", createOrder)
		customer.POST("/discount/check", checkDiscount)
	}
}
