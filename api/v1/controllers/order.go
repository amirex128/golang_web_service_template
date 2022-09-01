package controllers

import (
	"backend/api/v1/validations"
	"backend/internal/app/DTOs"
	"backend/internal/app/helpers"
	"backend/internal/app/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

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

	applyDiscount, err := applyDiscount(dto, discount)
	if err != nil {
		return
	}

	// محاسبه میزان تخفیف برای محصولات
	var finalProductPrice map[uint64]float32
	for pType, pIds := range applyDiscount {
		if pType == "percent" {
			for _, pId := range pIds {
				product, err := models.NewMainManager().FindProductById(c, pId)
				if err != nil {
					return
				}
				finalProductPrice[pId] = product.Price - (product.Price * (discount.Percent / 100))
			}
		} else {
			for _, pId := range pIds {
				product, err := models.NewMainManager().FindProductById(c, pId)
				if err != nil {
					return
				}
				finalProductPrice[pId] = product.Price - discount.Amount
			}
		}
	}

	err = models.NewMainManager().CreateOrder(c, dto)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "محصول با موفقیت ایجاد شد",
	})
	return
}

func applyDiscount(dto DTOs.CreateOrder, discount models.Discount) (map[string][]uint64, error) {

	var applyProductDiscount map[string][]uint64
	productDiscounts := strings.Split(discount.ProductIDs, ",")
	// درصورت وجود تخفیف برای محصولات خاص
	if len(productDiscounts) > 0 {
		// در صورتی که تخفیف درصدی بود
		if discount.Type == "percent" {
			for i := range productDiscounts {
				for i2 := range dto.OrderItems {
					if helpers.Uint64Convert(productDiscounts[i]) == dto.OrderItems[i2].ProductID {
						applyProductDiscount["percent"] = append(applyProductDiscount["percent"], dto.OrderItems[i2].ProductID)
					}
				}
			}
			// در صورتی که تخفیف مقداری بود
		} else {
			for i := range productDiscounts {
				for i2 := range dto.OrderItems {
					if helpers.Uint64Convert(productDiscounts[i]) == dto.OrderItems[i2].ProductID {
						applyProductDiscount["amount"] = append(applyProductDiscount["amount"], dto.OrderItems[i2].ProductID)
					}
				}
			}
		}
	} else {
		// در صورتی که تخفیف درصدی بود
		if discount.Type == "percent" {
			for i := range dto.OrderItems {
				applyProductDiscount["percent"] = append(applyProductDiscount["percent"], dto.OrderItems[i].ProductID)
			}
			// در صورتی که تخفیف مقداری بود
		} else {
			for i := range dto.OrderItems {
				applyProductDiscount["amount"] = append(applyProductDiscount["amount"], dto.OrderItems[i].ProductID)
			}
		}
	}
	return applyProductDiscount, nil
}
