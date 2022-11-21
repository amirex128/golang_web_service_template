package api

import (
	"github.com/amirex128/selloora_backend/internal/api/v1/controllers"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) {

	root := r.Group("/")
	{
		root.GET("/", controllers.IndexLanding)
		root.GET("blog", controllers.BlogLanding)
		root.GET("category/:id", controllers.CategoryLanding)
		root.GET("tag/:slug", controllers.TagLanding)
		root.GET("blog/:slug", controllers.DetailsLanding)
		root.GET("search/:search", controllers.SearchLanding)
		root.GET("page/:slug", controllers.PageLanding)

		//root.GET("contact", controllers.ContactLanding)
		//root.GET("faq", controllers.FaqLanding)
		//root.GET("pricing", controllers.PricingLanding)
		//root.GET("services", controllers.ServicesLanding)
		//root.GET("testimonial", controllers.TestimonialLanding)
		//root.GET("learn", controllers.LearnLanding)
		//root.GET("rules", controllers.RulesLanding)
		//root.GET("return-rules", controllers.ReturnRulesLanding)
	}
	v1 := r.Group("api/v1")
	v1.POST("/verify", authMiddleware.LoginHandler)
	v1.POST("/login/register", controllers.RegisterLogin)

	v1.POST("/user/ticket/create", controllers.CreateTicket)
	v1.POST("user/comment/create", controllers.CreateComment)

	user := v1.Group("user")
	user.Use(authMiddleware.MiddlewareFunc())
	{
		profile := user.Group("profile")
		{
			profile.POST("/update", controllers.UpdateProfile)
			profile.POST("/change-password", controllers.ChangePassword)
		}
		product := user.Group("product")
		{
			product.GET("/list", controllers.IndexProduct)
			product.GET("/show/:id", controllers.ShowProduct)
			product.POST("/create", controllers.CreateProduct)
			product.POST("/update", controllers.UpdateProduct)
			product.POST("/delete/:id", controllers.DeleteProduct)
		}
		ticket := user.Group("ticket")
		{
			ticket.GET("/list", controllers.IndexTicket)
			ticket.GET("/show/:id", controllers.ShowTicket)
		}
		gallery := user.Group("gallery")
		{
			gallery.POST("/create", controllers.CreateGallery)
			gallery.POST("/delete/:id", controllers.DeleteGallery)
		}
		discount := user.Group("discount")
		{
			discount.GET("/list", controllers.IndexDiscount)
			discount.GET("/show/:id", controllers.ShowDiscount)
			discount.POST("/create", controllers.CreateDiscount)
			discount.POST("/update", controllers.UpdateDiscount)
			discount.POST("/delete/:id", controllers.DeleteDiscount)
		}
		address := user.Group("address")
		{
			address.GET("/list", controllers.IndexAddress)
			address.POST("/create", controllers.CreateAddress)
			address.POST("/update", controllers.UpdateAddress)
			address.POST("/delete/:id", controllers.DeleteAddress)
		}
		order := user.Group("order")
		{
			order.POST("/send", controllers.SendOrder)
			order.GET("/list", controllers.IndexOrder)
			order.POST("/approve/:id", controllers.ApproveOrder)
			order.POST("/cancel/:id", controllers.CancelOrder)
			order.POST("/calculate", controllers.CalculateSendPrice)
			order.POST("/returned", controllers.ReturnedOrder)
			order.POST("/returned/accept", controllers.AcceptReturnedOrder)
			order.GET("/show/:id", controllers.ShowOrder)
			order.GET("/tracking/:id", controllers.TrackingOrder)
		}
		shop := user.Group("shop")
		{
			shop.GET("list", controllers.IndexShop)
			shop.POST("/create", controllers.CreateShop)
			shop.POST("/update", controllers.UpdateShop)
			shop.GET("/show/:id", controllers.ShowShop)
			shop.POST("/delete/:id", controllers.DeleteShop)
			shop.POST("/check", controllers.CheckSocial)
			shop.POST("/send-price", controllers.SendPrice)
			shop.GET("/instagram", controllers.GetInstagramPost)
		}
		post := user.Group("post")
		{
			post.GET("/list", controllers.IndexPost)
			post.GET("/show/:id", controllers.ShowPost)
			post.POST("/create", controllers.CreatePost)
			post.POST("/update", controllers.UpdatePost)
			post.POST("/delete/:id", controllers.DeletePost)
		}
		category := user.Group("category")
		{
			category.GET("/list", controllers.IndexCategory)
			category.GET("/create", controllers.CreateCategory)
			category.POST("/update", controllers.UpdateCategory)
			category.POST("/delete/:id", controllers.DeleteCategory)

		}
		comment := user.Group("comment")
		{
			comment.GET("/list", controllers.IndexCommentAdmin)
			comment.POST("/delete/:id", controllers.DeleteCommentAdmin)
			comment.POST("/approve/:id", controllers.ApproveCommentAdmin)
		}
		tag := user.Group("tag")
		{
			tag.GET("/list", controllers.IndexTag)
			tag.POST("/create", controllers.CreateTag)
			tag.POST("/delete/:id", controllers.DeleteTag)
			tag.POST("/add", controllers.AddTag)
		}
	}

	admin := v1.Group("admin")
	admin.Use(authMiddleware.MiddlewareFunc())
	{

	}

	customer := v1.Group("customer")
	{
		customer.POST("login/register", controllers.RequestCreateLoginCustomer)
		customer.POST("verify", controllers.VerifyCreateLoginCustomer)
		customer.POST("orders", controllers.IndexCustomerOrders)
		customer.POST("/sadad/verify", controllers.SadadPaymentVerify)
		customer.POST("/order/create", controllers.CreateOrder)
		customer.POST("/discount/check", controllers.CheckDiscount)
	}
}
