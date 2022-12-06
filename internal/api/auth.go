package api

import (
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/validations"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"log"
	"time"
)

func authMiddleware() *jwt.GinJWTMiddleware {

	// the jwt middleware
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     999999 * time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: "mobile",
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*models.User); ok {
				return jwt.MapClaims{
					"id":        v.ID,
					"is_admin":  v.IsAdmin,
					"expire_at": v.ExpireAt,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &models.User{
				ID:       uint64(claims["id"].(float64)),
				ExpireAt: claims["expire_at"].(string),
				IsAdmin:  claims["is_admin"].(bool),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			span, ctx := apm.StartSpan(c.Request.Context(), "controller:createTicket", "request")
			c.Request.WithContext(ctx)
			defer span.End()
			dto, err := validations.Verify(c)
			if err != nil {
				return nil, jwt.ErrFailedAuthentication
			}
			if dto.Password == "" {
				user, err := models.NewMysqlManager(c).FindUserByMobileAndCodeVerify(dto)
				if err != nil {
					return nil, jwt.ErrFailedAuthentication
				}
				return user, nil

			} else {
				user, err := models.NewMysqlManager(c).FindUserByMobileAndPassword(dto)
				if err != nil {
					return nil, jwt.ErrFailedAuthentication
				}
				return user, nil
			}
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			v, ok := data.(*models.User)
			if !ok {
				return false
			}
			if v.IsAdmin {
				return true
			}
			if v.ExpireAt == "" {
				return true
			}
			expireAt, err := time.Parse("2006-01-02 15:04:05", v.ExpireAt)
			if err != nil {
				return false
			}
			if time.Now().After(expireAt) {
				return false
			}

			return true
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
