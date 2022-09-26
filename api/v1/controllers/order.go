package controllers

import (
	"backend/api/v1/validations"
	"backend/internal/app/DTOs"
	"backend/internal/app/constants"
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

	user, err := models.NewMainManager().FindUserByID(c, dto.UserID)
	if err != nil {
		return
	}

	customer, err := models.NewMainManager().FindCustomerById(c, dto.CustomerID)
	if err != nil {
		return
	}

	if customer.VerifyCode != utils.GeneratePasswordHash(dto.VerifyCode) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "کد تایید صحیح نمی باشد"})
		return
	}

	discount, err := models.NewMainManager().FindDiscountById(c, dto.DiscountID)
	if err != nil {
		return
	}

	rawProducts, err := models.NewMainManager().FindProductByIds(c, extractProductIDs(dto))
	if err != nil {
		return
	}
	var products []models.Product
	for i := range rawProducts {
		var count uint64
		for j := range dto.OrderItems {
			if rawProducts[i].ID == dto.OrderItems[j].ProductID {
				count = dto.OrderItems[j].Count
			}
		}
		if rawProducts[i].Active == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"message":    "محصول " + rawProducts[i].Name + " غیر فعال است",
				"product_id": rawProducts[i].ID,
			})
			return
		}
		if rawProducts[i].Quantity < count {
			c.JSON(http.StatusBadRequest, gin.H{
				"message":    "موجودی محصول " + rawProducts[i].Name + " کافی نمی باشد",
				"product_id": rawProducts[i].ID,
			})
			return
		}
		if rawProducts[i].Status == constants.BlockProductStatus {
			c.JSON(http.StatusBadRequest, gin.H{
				"message":    "محصول " + rawProducts[i].Name + " مسدود شده است",
				"product_id": rawProducts[i].ID,
			})
			return
		}
	}

	productDiscounts := strings.Split(discount.ProductIDs, ",")
	applyDiscount := utils.ApplyDiscount(productDiscounts, utils.DiscountPriceType{
		Percent: discount.Percent,
		Amount:  discount.Amount,
		Type:    discount.Type,
	}, extractProductIDs(dto))
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

	var order models.Order
	for _, dis := range calculateDiscountProduct {
		order.TotalProductPrice += dis.RawPrice
		order.TotalDiscountPrice += dis.OffPrice
		order.TotalTaxPrice += dis.NewPrice * 0.09
		order.TotalProductDiscountPrice += dis.NewPrice
	}

	freeSend := true
	for i := range products {
		if products[i].FreeSend == 0 {
			freeSend = false
		}
	}
	if freeSend {
		order.SendPrice = 0
	} else {
		order.SendPrice = utils.CalculateSendPrice(user.ProvinceID, user.CityID, customer.ProvinceID, customer.CityID)
	}

	order.TotalFinalPrice = order.TotalProductDiscountPrice + order.TotalTaxPrice + order.SendPrice

	order.UserID = user.ID
	order.CustomerID = customer.ID
	order.DiscountID = discount.ID
	order.IP = c.ClientIP()
	order.Status = constants.PendingPaymentOrderStatus
	order.PaymentStatus = constants.PendingPaymentOrderPaymentStatus
	order.SendType = dto.SendType
	order.LastUpdateStatusAt = utils.NowTime()
	order.CreatedAt = utils.NowTime()

	err = models.NewMainManager().CreateOrder(c, order)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "محصول با موفقیت ایجاد شد",
	})
	return
}

func extractProductIDs(dto DTOs.CreateOrder) []uint64 {
	var productIDs []uint64
	for i := range dto.OrderItems {
		productIDs = append(productIDs, dto.OrderItems[i].ProductID)
	}
	return productIDs
}

func updateOrder() {

}
