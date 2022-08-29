package routers

import (
	"backend/internal/app/routers/v1"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Runner(host string, port string) {
	r := gin.Default()

	r.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
		}
		c.AbortWithStatus(http.StatusInternalServerError)
	}))

	v1.Routes(r)

	err := r.Run(host + ":" + port)
	if err != nil {
		panic(err)
	}

}
