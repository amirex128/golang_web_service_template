package landing

import (
	"fmt"
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
)

func PageLanding(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:PageLanding", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	shop, domain, theme, err := models.NewMysqlManager(c).FindShopByDomain(c.Request.Host)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}

	page, err := models.NewMysqlManager(c).FindPageBySlug(c.Param("slug"))
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	c.Set("template", fmt.Sprintf("themes/%d/%s", theme.ID, "page.html"))

	c.Set("data", map[string]interface{}{
		"theme":  theme,
		"domain": domain,
		"shop":   shop,
		"page":   page,
	})
}
