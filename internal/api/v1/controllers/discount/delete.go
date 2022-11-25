package discount

import (
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

// DeleteDiscount
// @Summary حذف تخفیف
// @description حذف تخفیف
// @Tags discount
// @Router       /user/discount/delete/{id} [post]
// @Param	Authorization	 header string	true "Authentication"
// @Param	id			 path   string	true "شناسه تخفیف" SchemaExample(1)
func DeleteDiscount(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:deleteDiscount", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	id := utils.StringToUint64(c.Param("id"))

	err := models.NewMysqlManager(c).DeleteDiscount(id)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "تخفیف با موفقیت حذف شد",
	})
}
