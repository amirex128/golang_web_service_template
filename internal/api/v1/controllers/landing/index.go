package landing

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
)

const (
	siteName = " | سلورا دستیار فروش اینستاگرام شما"
)

func IndexLanding(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:indexLanding", "request")
	c.Request.WithContext(ctx)
	defer span.End()

	c.Set("template", fmt.Sprintf("themes/1/index.html"))
	c.Set("data", map[string]interface{}{
		"title": "صفحه اصلی" + siteName,
	})
}
