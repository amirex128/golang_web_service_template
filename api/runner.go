package api

import (
	"backend/api/v1/controllers"
	"backend/api/v1/validations"
	"backend/internal/app/models"
	"fmt"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/flosch/pongo2"
	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func GetAuthMiddleware() *jwt.GinJWTMiddleware {

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
				ID:        uint64(claims["id"].(float64)),
				Email:     claims["email"].(string),
				Mobile:    claims["mobile"].(string),
				Status:    claims["status"].(string),
				Firstname: claims["firstname"].(string),
				Lastname:  claims["lastname"].(string),
				ExpireAt:  claims["expire_at"].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			login, err := validations.Login(c)
			if err != nil {
				return nil, jwt.ErrFailedAuthentication
			}
			user, err := models.NewMainManager().FindUserByMobilePassword(login)
			if err != nil {
				return nil, jwt.ErrFailedAuthentication
			}

			return user, nil
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			//if v, ok := data.(*models.User); ok && v.Mobile == "09024809750" {
			//	return true
			//}

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

func Runner(host string, port string) {
	r := gin.Default()
	r.Use(Pongo2())
	r.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
		}
		c.AbortWithStatus(http.StatusInternalServerError)
	}))
	r.Use(sentrygin.New(sentrygin.Options{
		Repanic:         true,
		WaitForDelivery: false,
		Timeout:         5 * time.Second,
	}))
	r.Use(func(ctx *gin.Context) {
		if hub := sentrygin.GetHubFromContext(ctx); hub != nil {
			hub.Scope().SetTag("someRandomTag", "maybeYouNeedIt")
		}
		ctx.Next()
	})

	controllers.Routes(r, GetAuthMiddleware())

	err := r.Run(host + ":" + port)
	if err != nil {
		sentry.CaptureException(err)
		panic(err)
	}

}
func Pongo2() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		name := stringFromContext(c, "template")
		data, _ := c.Get("data")

		if name == "" {
			return
		}

		template := pongo2.Must(pongo2.FromFile(fmt.Sprintf("%s/%s", "../../templates", name)))
		err := template.ExecuteWriter(convertContext(data), c.Writer)
		if err != nil {
			sentry.CaptureException(err)
			http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		}
	}
}

func stringFromContext(c *gin.Context, input string) string {
	raw, ok := c.Get(input)
	if ok {
		strVal, ok := raw.(string)
		if ok {
			return strVal
		}
	}
	return ""
}

func convertContext(thing interface{}) pongo2.Context {
	if thing != nil {
		context, isMap := thing.(map[string]interface{})
		if isMap {
			return context
		}
	}
	return nil
}
