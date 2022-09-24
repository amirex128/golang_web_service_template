package utils

import (
	"backend/internal/app/models"
)

// CalculateDiscountProduct  محاسبه میزان تخفیف برای محصولات
func CalculateDiscountProduct(applyDiscount map[string][]uint64, products []models.Product, discount models.Discount) map[uint64]float32 {
	var finalProductPrice map[uint64]float32
	for pType, pIds := range applyDiscount {
		if pType == "percent" {
			for _, pId := range pIds {
				var product models.Product
				for i := range products {
					if products[i].ID == pId {
						product = products[i]
						break
					}
				}
				finalProductPrice[pId] = product.Price - (product.Price * (discount.Percent / 100))
			}
		} else {
			for _, pId := range pIds {
				var product models.Product
				for i := range products {
					if products[i].ID == pId {
						product = products[i]
						break
					}
				}
				finalProductPrice[pId] = product.Price - discount.Amount
			}
		}
	}
	return finalProductPrice
}

// ApplyDiscount اعمال تخفیف روی محصول
func ApplyDiscount(productDiscounts []string, discount models.Discount, productIDs []uint64) map[string][]uint64 {
	var applyDiscount map[string][]uint64

	if len(productDiscounts) > 0 {
		// در صورتی که تخفیف درصدی بود
		if discount.Type == "percent" {
			for i := range productDiscounts {
				for i2 := range productIDs {
					if StringToUint64(productDiscounts[i]) == productIDs[i2] {
						applyDiscount["percent"] = append(applyDiscount["percent"], productIDs[i2])
					}
				}
			}
			// در صورتی که تخفیف مقداری بود
		} else {
			for i := range productDiscounts {
				for i2 := range productIDs {
					if StringToUint64(productDiscounts[i]) == productIDs[i2] {
						applyDiscount["amount"] = append(applyDiscount["amount"], productIDs[i2])
					}
				}
			}
		}
	} else {
		// در صورتی که تخفیف درصدی بود
		if discount.Type == "percent" {
			for i := range productIDs {
				applyDiscount["percent"] = append(applyDiscount["percent"], productIDs[i])
			}
			// در صورتی که تخفیف مقداری بود
		} else {
			for i := range productIDs {
				applyDiscount["amount"] = append(applyDiscount["amount"], productIDs[i])
			}
		}

	}
	return applyDiscount

}
