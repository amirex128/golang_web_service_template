package discount

import (
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

// ShowDiscount
// @Summary نمایش تخفیف
// @description نمایش تخفیف
// @Tags discount
// @Router       /user/discount/show/{id} [get]
// @Param	Authorization	 header string	true "Authentication"
// @Param	id			 path   string	true "شناسه تخفیف" SchemaExample(1)
func ShowDiscount(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:showDiscount", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	id := utils.StringToUint64(c.Param("id"))
	userID := utils.GetUserID(c)

	discount, err := models.NewMysqlManager(c).FindDiscountById(id)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	if *discount.UserID != *userID {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "شما اجازه مشاهده این تخفیف را ندارید",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"data": discount,
	})
}
