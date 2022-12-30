package api

import (
	"github.com/amirex128/selloora_backend/docs"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/k0kubun/pp/v3"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.elastic.co/apm/module/apmgin/v2"
	"runtime/debug"
	"time"
)

func Runner() *gin.Engine {
	r := gin.Default()
	r.Use(apmgin.Middleware(r))
	//r.Use(pongo())
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "OPTION"},
		AllowHeaders:     []string{"Authorization", "type_auth", "content-type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	r.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(error); ok {
			pp.Fatalln("------------------Stack-Trace------------------")
			debug.PrintStack()
			pp.Fatalln("------------------Error-Message------------------")
			pp.Println(err.Error())
			return
		}
	}))
	r.Static("/public", "./public")
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "not found"})
	})
	docs.SwaggerInfo.BasePath = "/api/v1"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	Routes(r, authMiddleware())

	return r
}
