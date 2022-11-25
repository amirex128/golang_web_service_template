package landing

import (
	"fmt"
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/amirex128/selloora_backend/internal/validations"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
)

func BlogLanding(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:blogLanding", "request")
	c.Request.WithContext(ctx)
	defer span.End()

	dto, err := validations.IndexPost(c)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	posts, err := models.NewMysqlManager(c).GetAllPostWithPagination(dto)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	postArray := posts.Data.([]models.Post)
	for i := range postArray {
		postArray[i].CreatedAt = utils.DateToJalaali(postArray[i].CreatedAt)
		postArray[i].UpdatedAt = utils.DateToJalaali(postArray[i].UpdatedAt)
	}
	posts.Data = postArray
	tags, err := models.NewMysqlManager(c).RandomTags(20)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}

	categories, err := models.NewMysqlManager(c).GetLevel1Categories()
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}

	randomPosts, err := models.NewMysqlManager(c).RandomPost(6)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	for i := range randomPosts {
		randomPosts[i].CreatedAt = func() string {
			ago := utils.DateAgo(randomPosts[i].CreatedAt)
			if ago == 0 {
				return "امروز"
			}
			return fmt.Sprintf("%d روز قبل", ago)
		}()
		randomPosts[i].UpdatedAt = func() string {
			ago := utils.DateAgo(randomPosts[i].UpdatedAt)
			if ago == 0 {
				return "امروز"
			}
			return fmt.Sprintf("%d روز قبل", ago)
		}()
	}

	shop, domain, theme, err := models.NewMysqlManager(c).FindShopByDomain(c.Request.Host)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}

	c.Set("template", fmt.Sprintf("themes/%d/%s", theme.ID, "blog.html"))

	c.Set("data", map[string]interface{}{
		"theme":        theme,
		"domain":       domain,
		"shop":         shop,
		"title":        "بلاگ" + siteName,
		"posts":        posts,
		"random_posts": randomPosts,
		"categories":   categories,
		"tags":         tags,
	})
}
