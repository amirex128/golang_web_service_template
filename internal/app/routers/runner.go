package routers

import (
	"backend/internal/app/DTOs"
	"backend/internal/app/models"
	"backend/internal/app/routers/v1/controllers"
	"fmt"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
	"time"
)

func GetAuthMiddleware() *jwt.GinJWTMiddleware {

	// the jwt middleware
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: "mobile",
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*models.User); ok {
				return jwt.MapClaims{
					"id":        v.ID,
					"email":     v.Email,
					"mobile":    v.Mobile,
					"status":    v.Status,
					"firstname": v.Firstname,
					"lastname":  v.Lastname,
					"expire_at": v.ExpireAt,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &models.User{
				ID:        int64(claims["id"].(float64)),
				Email:     claims["email"].(string),
				Mobile:    claims["mobile"].(string),
				Status:    claims["status"].(string),
				Firstname: claims["firstname"].(string),
				Lastname:  claims["lastname"].(string),
				ExpireAt:  claims["expire_at"].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var login DTOs.Login
			if err := c.ShouldBind(&login); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			validate := validator.New()
			err := validate.Struct(login)
			if err != nil {
				if _, ok := err.(*validator.InvalidValidationError); ok {
					c.JSON(http.StatusInternalServerError, gin.H{
						"message": "مقادیر ارسال شده نا درست میباشد",
						"error":   err.Error(),
					})
					return "", jwt.ErrMissingLoginValues
				}
				var errors []gin.H
				for _, err := range err.(validator.ValidationErrors) {
					if err.StructField() == "Mobile" {
						errors = append(errors, gin.H{
							"message": "شماره موبایل نامعتبر میباشد",
						})
					}
					if err.StructField() == "Password" {
						if err.Tag() == "min" {
							errors = append(errors, gin.H{
								"message": "رمز عبور باید حداقل 8 کاراکتر باشد",
							})
						}
						if err.Tag() == "max" {
							errors = append(errors, gin.H{
								"message": "رمز عبور باید حداکثر 20 کاراکتر باشد",
							})
						}
						if err.Tag() == "required" {
							errors = append(errors, gin.H{
								"message": "رمز عبور باید وارد شود",
							})
						}

					}

					c.JSON(http.StatusBadRequest, errors)
					return "", jwt.ErrMissingLoginValues
				}
			}
			user, err := models.NewMainManager().FindUserByMobilePassword(login)
			if err != nil {
				return nil, jwt.ErrFailedAuthentication
			}

			return user, nil
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(*models.User); ok && v.Mobile == "09024809750" {
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	// When you use jwt.New(), the function is already automatically called for checking,
	// which means you don't need to call it again.
	errInit := authMiddleware.MiddlewareInit()

	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}

	return authMiddleware
}

func Runner(host string, port string) {
	r := gin.Default()

	r.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
		}
		c.AbortWithStatus(http.StatusInternalServerError)
	}))

	controllers.Routes(r, GetAuthMiddleware())

	err := r.Run(host + ":" + port)
	if err != nil {
		panic(err)
	}

}
