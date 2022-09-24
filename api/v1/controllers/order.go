package controllers

import (
	"backend/api/v1/validations"
	"backend/internal/app/models"
	"backend/internal/app/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// createOrder ثبت یک سفارش جدید
func createOrder(c *gin.Context) {
	dto, err := validations.CreateOrder(c)
	if err != nil {
		return
	}

	customer, err := models.NewMainManager().FindCustomerById(c, dto.CustomerID)
	if err != nil {
		return
	}

	if customer.VerifyCode != dto.VerifyCode {
		c.JSON(http.StatusBadRequest, gin.H{"error": "کد تایید صحیح نمی باشد"})
	}

	discount, err := models.NewMainManager().FindDiscountById(c, dto.DiscountID)
	if err != nil {
		return
	}

	var productIDs []uint64
	for i := range dto.OrderItems {
		productIDs = append(productIDs, dto.OrderItems[i].ProductID)
	}
	products, err := models.NewMainManager().FindProductByIds(c, productIDs)
	if err != nil {
		return
	}

	productDiscounts := strings.Split(discount.ProductIDs, ",")
	applyDiscount := utils.ApplyDiscount(productDiscounts, discount, productIDs)
	utils.CalculateDiscountProduct(applyDiscount, products, discount)

	err = models.NewMainManager().CreateOrder(c, dto)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "محصول با موفقیت ایجاد شد",
	})
	return
}
