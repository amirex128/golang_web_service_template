package api

import (
	"github.com/amirex128/selloora_backend/internal/api/v1/controllers/address"
	"github.com/amirex128/selloora_backend/internal/api/v1/controllers/auth"
	"github.com/amirex128/selloora_backend/internal/api/v1/controllers/category"
	"github.com/amirex128/selloora_backend/internal/api/v1/controllers/comment"
	"github.com/amirex128/selloora_backend/internal/api/v1/controllers/customer"
	"github.com/amirex128/selloora_backend/internal/api/v1/controllers/dev_ops"
	"github.com/amirex128/selloora_backend/internal/api/v1/controllers/discount"
	"github.com/amirex128/selloora_backend/internal/api/v1/controllers/domain"
	"github.com/amirex128/selloora_backend/internal/api/v1/controllers/gallery"
	"github.com/amirex128/selloora_backend/internal/api/v1/controllers/landing"
	"github.com/amirex128/selloora_backend/internal/api/v1/controllers/menu"
	"github.com/amirex128/selloora_backend/internal/api/v1/controllers/order"
	"github.com/amirex128/selloora_backend/internal/api/v1/controllers/page"
	"github.com/amirex128/selloora_backend/internal/api/v1/controllers/post"
	"github.com/amirex128/selloora_backend/internal/api/v1/controllers/product"
	"github.com/amirex128/selloora_backend/internal/api/v1/controllers/shop"
	"github.com/amirex128/selloora_backend/internal/api/v1/controllers/slider"
	"github.com/amirex128/selloora_backend/internal/api/v1/controllers/tag"
	"github.com/amirex128/selloora_backend/internal/api/v1/controllers/theme"
	"github.com/amirex128/selloora_backend/internal/api/v1/controllers/ticket"
	"github.com/amirex128/selloora_backend/internal/api/v1/controllers/user"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) {

	root := r.Group("/")
	{
		root.GET("/", landing.IndexLanding)
		root.GET("blog", landing.BlogLanding)
		root.GET("category/:id", landing.CategoryLanding)
		root.GET("tag/:slug", landing.TagLanding)
		root.GET("blog/:slug", landing.DetailsLanding)
		root.GET("search/:search", landing.SearchLanding)
		root.GET("page/:slug", landing.PageLanding)
	}

	v1 := r.Group("api/v1")
	{
		v1.POST("/verify", authMiddleware.LoginHandler)
		v1.POST("/login/register", auth.RegisterLogin)
		v1.POST("/forget", auth.ForgetPassword)

		v1.POST("/ticket/create", ticket.CreateTicket)
		v1.POST("/comment/create", comment.CreateComment)

		v1.GET("/healthcheck", dev_ops.HealthCheck)
		v1.GET("/metrics", dev_ops.Metrics)
	}

	webShop := v1.Group("web")
	{
		webShop.GET("home")
	}
	_user := v1.Group("user")
	_user.Use(authMiddleware.MiddlewareFunc())
	{
		profile := _user.Group("profile")
		{
			profile.POST("/update", user.UpdateUser)
			profile.POST("/change-password", auth.ChangePassword)
		}
		_product := _user.Group("product")
		{
			_product.GET("/list", product.IndexProduct)
			_product.GET("/show/:id", product.ShowProduct)
			_product.POST("/create", product.CreateProduct)
			_product.POST("/update/*id", product.UpdateProduct)
			_product.POST("/delete/:id", product.DeleteProduct)
		}
		_ticket := _user.Group("ticket")
		{
			_ticket.POST("/create", ticket.CreateTicket)
			_ticket.GET("/list", ticket.IndexTicket)
			_ticket.POST("/delete/:id", ticket.DeleteTicket)
			_ticket.GET("/show/:id", ticket.ShowTicket)
		}
		_gallery := _user.Group("gallery")
		{
			_gallery.POST("/create", gallery.CreateGallery)
			_gallery.GET("/show/:id", gallery.ShowGallery)
			_gallery.GET("/list", gallery.IndexGallery)
			_gallery.POST("/delete/:id", gallery.DeleteGallery)
		}
		_discount := _user.Group("discount")
		{
			_discount.GET("/list", discount.IndexDiscount)
			_discount.GET("/show/:id", discount.ShowDiscount)
			_discount.POST("/create", discount.CreateDiscount)
			_discount.POST("/update/*id", discount.UpdateDiscount)
			_discount.POST("/delete/:id", discount.DeleteDiscount)
		}
		_address := _user.Group("address")
		{
			_address.GET("/list", address.IndexAddress)
			_address.POST("/create", address.CreateAddress)
			_address.POST("/update/*id", address.UpdateAddress)
			_address.GET("/show/:id", address.ShowAddress)
			_address.POST("/delete/:id", address.DeleteAddress)
		}
		_domain := _user.Group("domain")
		{
			_domain.GET("/list", domain.IndexDomain)
			_domain.POST("/create", domain.CreateDomain)
			_domain.GET("/show/:id", domain.ShowDomain)
			_domain.POST("/delete/:id", domain.DeleteDomain)
		}
		_order := _user.Group("order")
		{
			_order.POST("/send/:id", order.SendOrder)
			_order.GET("/list", order.IndexOrder)
			_order.POST("/approve/:id", order.ApproveOrder)
			_order.POST("/cancel/:id", order.CancelOrder)
			_order.POST("/calculate", order.CalculateSendPrice)
			_order.POST("/returned/:id", order.ReturnedOrder)
			_order.POST("/returned/accept/:id", order.AcceptReturnedOrder)
			_order.GET("/show/:id", order.ShowOrder)
			_order.GET("/tracking/:id", order.TrackingOrder)
			_order.POST("/delete/:id", order.DeleteOrder)

		}
		_shop := _user.Group("shop")
		{
			_shop.GET("list", shop.IndexShop)
			_shop.POST("/create", shop.CreateShop)
			_shop.POST("/update/*id", shop.UpdateShop)
			_shop.GET("/show/:id", shop.ShowShop)
			_shop.POST("/delete/:id", shop.DeleteShop)
			_shop.POST("/send-price/:id", shop.SendPriceShop)
			_shop.POST("/select/theme", shop.SelectThemeShop)
			_shop.GET("/instagram", shop.GetInstagramPost)
		}

		_slider := _user.Group("slider")
		{
			_slider.GET("/list", slider.IndexSlider)
			_slider.POST("/create", slider.CreateSlider)
			_slider.GET("/show/:id", slider.ShowSlider)
			_slider.POST("/update/*id", slider.UpdateSlider)
			_slider.POST("/delete/:id", slider.DeleteSlider)
		}

		_post := _user.Group("post")
		{
			_post.GET("/list", post.IndexPost)
			_post.GET("/show/:id", post.ShowPost)
			_post.POST("/create", post.CreatePost)
			_post.POST("/update/*id", post.UpdatePost)
			_post.POST("/delete/:id", post.DeletePost)
		}
		_page := _user.Group("page")
		{
			_page.GET("/list", page.IndexPage)
			_page.GET("/show/:id", page.ShowPage)
			_page.POST("/create", page.CreatePage)
			_page.POST("/update/*id", page.UpdatePage)
			_page.POST("/delete/:id", page.DeletePage)
		}
		_customer := _user.Group("customer")
		{
			_customer.GET("/list", customer.IndexCustomer)
			_customer.GET("/show/:id", customer.ShowCustomer)
			_customer.POST("/delete/:id", customer.DeleteCustomer)
		}
		_menu := _user.Group("menu")
		{
			_menu.GET("/list", menu.IndexMenu)
			_menu.GET("/show/:id", menu.ShowMenu)
			_menu.POST("/create", menu.CreateMenu)
			_menu.POST("/update/*id", menu.UpdateMenu)
			_menu.POST("/delete/:id", menu.DeleteMenu)
		}
		_theme := _user.Group("theme")
		{
			_theme.GET("/list", theme.IndexTheme)
		}
		_category := _user.Group("category")
		{
			_category.GET("/list", category.IndexCategory)
			_category.GET("/show/:id", category.ShowCategory)
			_category.POST("/create", category.CreateCategory)
			_category.POST("/update/*id", category.UpdateCategory)
			_category.POST("/delete/:id", category.DeleteCategory)

		}
		_comment := _user.Group("comment")
		{
			_comment.GET("/list", comment.IndexComment)
			_comment.GET("/show/:id", comment.ShowComment)
			_comment.POST("/delete/:id", comment.DeleteComment)
			_comment.POST("/approve/:id", comment.ApproveComment)
		}
		_tag := _user.Group("tag")
		{
			_tag.GET("/list", tag.IndexTag)
			_tag.POST("/create", tag.CreateTag)
			_tag.POST("/delete/:id", tag.DeleteTag)
			_tag.POST("/add", tag.AddTag)
		}
	}

	admin := v1.Group("admin")
	admin.Use(authMiddleware.MiddlewareFunc())
	{

	}

	_customer := v1.Group("customer")
	{
		_customer.POST("login/register", customer.RequestCreateLoginCustomer)
		_customer.POST("verify", customer.VerifyCreateLoginUpdateCustomer)
		_customer.POST("orders", order.IndexCustomerOrders)
		_customer.POST("/sadad/verify", order.SadadPaymentVerify)
		_customer.POST("/order/create", order.CreateOrder)
		_customer.POST("/discount/check", discount.CheckDiscount)
	}
}
