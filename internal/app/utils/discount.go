package utils

import (
	"backend/internal/app/constants"
)

type ApplyDiscountType struct {
	RawPrice float32 `json:"raw_price"`
	OffPrice float32 `json:"off_price"`
	NewPrice float32 `json:"new_price"`
}

type DiscountPriceType struct {
	Percent float32
	Amount  float32
	Type    string
}
type ProductDiscountCalculatorType struct {
	ProductID uint64
	Price     float32
}

func (p ProductDiscountCalculatorType) GetID() uint64 {
	return p.ProductID
}

// CalculateDiscountProduct  محاسبه میزان تخفیف برای محصولات
func CalculateDiscountProduct(applyDiscount map[string][]uint64, products []ProductDiscountCalculatorType, discount DiscountPriceType) map[uint64]ApplyDiscountType {
	var finalProductPrice map[uint64]ApplyDiscountType
	for pType, pIds := range applyDiscount {
		if pType == "percent" {
			for _, pId := range pIds {
				product := FindId(products, pId)
				finalProductPrice[pId] = ApplyDiscountType{
					RawPrice: product.Price,
					OffPrice: product.Price * (discount.Percent / 100),
					NewPrice: product.Price - (product.Price * (discount.Percent / 100)),
				}
			}
		} else {
			for _, pId := range pIds {
				product := FindId(products, pId)
				finalProductPrice[pId] = ApplyDiscountType{
					RawPrice: product.Price,
					OffPrice: discount.Amount,
					NewPrice: product.Price - discount.Amount,
				}
			}
		}
	}
	return finalProductPrice
}

func FindId[T constants.IModel](products []T, pId uint64) T {
	var product T
	for i := range products {
		if products[i].GetID() == pId {
			product = products[i]
			break
		}
	}
	return product
}

// ApplyDiscount اعمال تخفیف روی محصول
func ApplyDiscount(productDiscounts []string, discount DiscountPriceType, productIDs []uint64) map[string][]uint64 {
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
