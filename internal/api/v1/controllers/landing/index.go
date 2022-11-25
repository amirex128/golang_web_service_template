package landing

import (
	"fmt"
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
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

	shop, domain, theme, err := models.NewMysqlManager(c).FindShopByDomain(c.Request.Host)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}

	c.Set("template", fmt.Sprintf("themes/%d/%s", theme.ID, "index.html"))
	c.Set("data", map[string]interface{}{
		"theme":  theme,
		"domain": domain,
		"shop":   shop,
		"title":  "صفحه اصلی" + siteName,
	})
}
