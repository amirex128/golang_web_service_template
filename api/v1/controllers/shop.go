package controllers

import (
	"backend/api/v1/validations"
	"backend/internal/app/DTOs"
	"backend/internal/app/models"
	"backend/internal/app/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func createShop(c *gin.Context) {
	dto, err := validations.CreateShop(c)
	if err != nil {
		return
	}
	userID := utils.GetUser(c)

	err = models.NewMainManager().CreateShop(c, dto, userID)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "فروشگاه با موفقیت ایجاد شد",
	})
}

func updateShop(c *gin.Context) {
	dto, err := validations.UpdateShop(c)
	if err != nil {
		return
	}
	userID := utils.GetUser(c)
	shopID := utils.StringToUint64(c.Param("id"))

	err = models.NewMainManager().UpdateShop(c, dto, shopID, userID)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "فروشگاه با موفقیت ویرایش شد",
	})
}

func deleteShop(c *gin.Context) {
	shopID := utils.StringToUint64(c.Param("id"))
	userID := utils.GetUser(c)
	dto, err := validations.DeleteShop(c)
	if err != nil {
		return
	}
	if dto.ProductBehave == "move" {
		err = models.NewMainManager().MoveProducts(c, shopID, dto.NewShopID, userID)
		if err != nil {
			return
		}
	} else if dto.ProductBehave == "delete_product" {
		err = models.NewMainManager().DeleteProducts(c, shopID, userID)
		if err != nil {
			return
		}
	}

	err = models.NewMainManager().DeleteShop(c, shopID, userID)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "فروشگاه با موفقیت حذف شد",
	})
}

func indexShop(c *gin.Context) {
	userID := utils.GetUser(c)
	dto, err := validations.IndexShop(c)
	shops, err := models.NewMainManager().GetAllShopWithPagination(c, dto, userID)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"shops": shops,
	})
}

func checkSocial(c *gin.Context) {
	dto, err := validations.CheckSocial(c)
	if err != nil {
		return
	}
	userID := utils.GetUser(c)
	// TODO بررسی وضعیت تایید شبکه اجتماعی
	var resultCheck bool
	resultCheck = true
	err = models.NewMainManager().UpdateShop(c, DTOs.UpdateShop{
		VerifySocial: true,
	}, dto.ShopID, userID)
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

func sendPrice(c *gin.Context) {
	dto, err := validations.SendPrice(c)
	if err != nil {
		return
	}
	userID := utils.GetUser(c)

	err = models.NewMainManager().UpdateShop(c, DTOs.UpdateShop{
		SendPrice: dto.SendPrice,
	}, dto.ShopID, userID)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "هزینه ارسال با موفقیت بروزرسانی شد",
	})
}

func getInstagramPost(c *gin.Context) {
	//TODO

}
