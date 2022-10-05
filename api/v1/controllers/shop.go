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
	image, err := utils.UploadImage(c, dto.Logo, "shop/user_"+utils.Uint64ToString(userID))
	if err != nil {
		return
	}
	dto.LogoPath = image
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

	if dto.Logo != nil {
		image, err := utils.UploadImage(c, dto.Logo, "shop/user_"+utils.Uint64ToString(userID))
		if err != nil {
			return
		}
		dto.LogoPath = image
	}
	if dto.LogoRemove != "" {
		utils.RemoveImages([]string{dto.LogoRemove})
	}

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
	err := models.NewMainManager().DeleteShop(c, shopID, userID)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "فروشگاه با موفقیت حذف شد",
	})
}

func indexShop(c *gin.Context) {
	userID := utils.GetUser(c)
	shops, err := models.NewMainManager().IndexShop(c, userID)
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
	err = models.NewMainManager().UpdateShop(c, DTOs.UpdateShop{
		VerifySocial: true,
	}, dto.ShopID, userID)
	if err != nil {
		return
	}
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
}

func getInstagramPost(c *gin.Context) {
	//TODO

}
