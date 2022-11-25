package shop

import (
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
)

func GetInstagramPost(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:getInstagramPost", "request")
	c.Request.WithContext(ctx)
	defer span.End()

}
