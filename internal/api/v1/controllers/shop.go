package controllers

import (
	"github.com/amirex128/selloora_backend/internal/DTOs"
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/amirex128/selloora_backend/internal/validations"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
)

func CreateShop(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "createShop", "request")
	defer span.End()
	dto, err := validations.CreateShop(c)
	if err != nil {
		return
	}
	userID := models.GetUser(c)

	err = models.NewMysqlManager(ctx).CreateShop(c, ctx, dto, *userID)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "فروشگاه با موفقیت ایجاد شد",
	})
}

func UpdateShop(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "updateShop", "request")
	defer span.End()
	dto, err := validations.UpdateShop(c)
	if err != nil {
		return
	}

	err = models.NewMysqlManager(ctx).UpdateShop(c, ctx, dto)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "فروشگاه با موفقیت ویرایش شد",
	})
}

func DeleteShop(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "deleteShop", "request")
	defer span.End()
	shopID := utils.StringToUint64(c.Param("id"))
	userID := models.GetUser(c)
	dto, err := validations.DeleteShop(c)
	if err != nil {
		return
	}
	if dto.ProductBehave == "move" {
		err = models.NewMysqlManager(ctx).MoveProducts(c, ctx, shopID, dto.NewShopID, *userID)
		if err != nil {
			return
		}
	} else if dto.ProductBehave == "delete_product" {
		err = models.NewMysqlManager(ctx).DeleteProducts(c, ctx, shopID, *userID)
		if err != nil {
			return
		}
	}

	err = models.NewMysqlManager(ctx).DeleteShop(c, ctx, shopID, *userID)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "فروشگاه با موفقیت حذف شد",
	})
}

func ShowShop(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "showShop", "request")
	defer span.End()
	shopID := utils.StringToUint64(c.Param("id"))
	shop, err := models.NewMysqlManager(ctx).FindShopByID(c, ctx, shopID)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"shop": shop,
	})
}

func IndexShop(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "indexShop", "request")
	defer span.End()
	dto, err := validations.IndexShop(c)
	if err != nil {
		return
	}

	shops, err := models.NewMysqlManager(ctx).GetAllShopWithPagination(c, ctx, dto)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"shops": shops,
	})
}

func CheckSocial(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "checkSocial", "request")
	defer span.End()
	_, err := validations.CheckSocial(c)
	if err != nil {
		return
	}
	// TODO بررسی وضعیت تایید شبکه اجتماعی
	var resultCheck bool
	resultCheck = true
	err = models.NewMysqlManager(ctx).UpdateShop(c, ctx, DTOs.UpdateShop{
		VerifySocial: true,
	})
	if err != nil {
		return
	}
	if resultCheck {
		c.JSON(http.StatusOK, gin.H{
			"message": "تایید شبکه اجتماعی با موفقیت انجام شد",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "تایید شبکه اجتماعی با انجام نشد",
	})
}

func SendPrice(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "sendPrice", "request")
	defer span.End()
	dto, err := validations.SendPrice(c)
	if err != nil {
		return
	}
	err = models.NewMysqlManager(ctx).UpdateShop(c, ctx, DTOs.UpdateShop{
		SendPrice: dto.SendPrice,
	})
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "هزینه ارسال با موفقیت بروزرسانی شد",
	})
}

func GetInstagramPost(c *gin.Context) {
	span, _ := apm.StartSpan(c.Request.Context(), "getInstagramPost", "request")
	defer span.End()

}
