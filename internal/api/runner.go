package api

import (
	"github.com/amirex128/selloora_backend/docs"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/k0kubun/pp/v3"
	"github.com/mandrigin/gin-spa/spa"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.elastic.co/apm/module/apmgin/v2"
	"time"
)

func Runner() *gin.Engine {
	r := gin.Default()
	r.Use(apmgin.Middleware(r))
	r.Use(pongo())
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "OPTION"},
		AllowHeaders:     []string{"Authorization", "type_auth", "content-type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	r.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		pp.Println("-------------------------error-----------------------")
		if err, ok := recovered.(error); ok {
			e2 := errorx.New("خطای پانیک رخ داده است", "panic", err)
			errorx.ResponseErrorx(c, e2)
			return
		}
	}))
	r.Static("/public", "./public")
	r.NoRoute(func(c *gin.Context) {
		c.Set("template", "404-error.html")
	})
	docs.SwaggerInfo.BasePath = "/api/v1"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	Routes(r, authMiddleware())

	r.Use(spa.Middleware("/", "./frontend"))
	return r
}
