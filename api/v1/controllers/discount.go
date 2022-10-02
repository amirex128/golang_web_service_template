package controllers

import (
	"backend/api/v1/validations"
	"backend/internal/app/models"
	"backend/internal/app/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// checkDiscount بررسی درستی تخفیف
func checkDiscount(c *gin.Context) {
	dto, err := validations.CheckDiscount(c)
	if err != nil {
		return
	}
	productIDs := dto.ProductIDs

	discount, err := models.NewMainManager().FindDiscountByCodeAndUserID(c, dto.Code, dto.UserID)
	if err != nil {
		return
	}
	products, err := models.NewMainManager().FindProductByIds(c, productIDs)
	if err != nil {
		return
	}

	productDiscounts := strings.Split(discount.ProductIDs, ",")

	applyDiscount := utils.ApplyDiscount(productDiscounts, utils.DiscountPriceType{
		Percent: discount.Percent,
		Amount:  discount.Amount,
		Type:    discount.Type,
	}, productIDs)
	var productCalculate []utils.ProductDiscountCalculatorType
	for i := range products {
		productCalculate = append(productCalculate, utils.ProductDiscountCalculatorType{
			ProductID: products[i].ID,
			Price:     products[i].Price,
		})
	}
	calculateDiscountProduct := utils.CalculateDiscountProduct(applyDiscount, productCalculate, utils.DiscountPriceType{
		Percent: discount.Percent,
		Amount:  discount.Amount,
		Type:    discount.Type,
	})

	c.JSON(http.StatusOK, gin.H{
		"result": calculateDiscountProduct,
	})
}

// createDiscount ایجاد تخفیف
func createDiscount(c *gin.Context) {
	dto, err := validations.CreateDiscount(c)
	if err != nil {
		return
	}

	userID := utils.GetUser(c)

	err = models.NewMainManager().CreateDiscount(c, dto, userID)

	c.JSON(http.StatusOK, gin.H{
		"message": "تخفیف با موفقیت ایجاد شد",
	})

}

// updateDiscount بروزرسانی تخفیف
func updateDiscount(c *gin.Context) {
	dto, err := validations.UpdateDiscount(c)
	if err != nil {
		return
	}
	userID := utils.GetUser(c)

	err = models.NewMainManager().UpdateDiscount(c, dto, userID)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "تخفیف با موفقیت بروزرسانی شد",
	})
}

// indexDiscount لیست تخفیفات
func indexDiscount(c *gin.Context) {
	dto, err := validations.IndexDiscount(c)
	if err != nil {
		return
	}
	userID := utils.GetUser(c)

	discounts, err := models.NewMainManager().GetAllDiscountWithPagination(c, dto, userID)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"result": discounts,
	})
}

// deleteDiscount حذف تخفیف
func deleteDiscount(c *gin.Context) {
	id := utils.StringToUint64(c.Param("id"))
	userID := utils.GetUser(c)

	err := models.NewMainManager().DeleteDiscount(c, id, userID)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "تخفیف با موفقیت حذف شد",
	})
}

// showDiscount نمایش تخفیف
func showDiscount(c *gin.Context) {
	id := utils.StringToUint64(c.Param("id"))
	userID := utils.GetUser(c)

	discount, err := models.NewMainManager().FindDiscountById(c, id)
	if err != nil {
		return
	}
	if discount.UserID != userID {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "شما اجازه مشاهده این تخفیف را ندارید",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "تخفیف با موفقیت حذف شد",
	})
}
