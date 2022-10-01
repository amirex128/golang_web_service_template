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

	shop, err := models.NewMainManager().FindShopByID(c, dto.ShopID)
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

	order.TotalFinalPrice = order.TotalProductDiscountPrice + order.TotalTaxPrice + order.SendPrice

	order.SendPrice = shop.SendPrice
	order.UserID = user.ID
	order.CustomerID = customer.ID
	order.DiscountID = discount.ID
	order.IP = c.ClientIP()
	order.Status = constants.PendingPaymentOrderStatus
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

// indexOrder لیست کردن تمامی سفارشات پنل ادمین و از طریق فیلتر قابلیت تقسیم به سه بخش سفارش جدید در حتال انجام و تمام شده
func indexOrder(c *gin.Context) {
	orderStatus := c.Param("order_status")
	userID := utils.GetUser(c)
	var orders []models.Order
	var err error
	if orderStatus == "new" {
		orders, err = models.NewMainManager().GetOrders(c, userID, []string{
			constants.PendingAcceptOrderStatus,
		})
	} else if orderStatus == "processing" {
		orders, err = models.NewMainManager().GetOrders(c, userID, []string{
			constants.AcceptedOrderStatus,
			constants.PendingReceivePostOrderStatus,
			constants.ReceivedPostOrderStatus,
			constants.ReceivedCustomerOrderStatus,
			constants.PendingReturnOrderStatus,
			constants.AcceptedReturnOrderStatus,
			constants.RejectedReturnOrderStatus,
			constants.PendingReceivePostReturnOrderStatus,
			constants.ReceivedPostReturnOrderStatus,
			constants.ReceivedOwnerOrderStatus,
		})
	} else if orderStatus == "completed" {
		orders, err = models.NewMainManager().GetOrders(c, userID, []string{
			constants.FinishedOrderStatus,
		})
	}
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"orders": orders,
	})

}
func acceptOrder(c *gin.Context) {
	orderID := utils.StringToUint64(c.Param("orderID"))
	userID := utils.GetUser(c)
	order, err := models.NewMainManager().FindOrderByID(c, orderID, userID)
	err = models.NewMainManager().UpdateOrder(c, order)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "خطا در تایید سفارش",
		})
		return
	}
}
func sendOrder(c *gin.Context) {
	// TODO دریافت اطلاعات وزن بسته و ارسال به این وبسرویس برای محاسبه هزینه ارسال سیستم های پستی مختلف
}

func chooseSenderOrder(c *gin.Context) {
	// TODOki انتخاب سیستم پستی مورد نظر برای ارسال سفارش
}

func updateOrder(c *gin.Context) {
	//TODO

}

func cancelOrder(c *gin.Context) {
	//TODO

}

func returnedOrder(c *gin.Context) {
	//TODO

}

func acceptReturnedOrder(c *gin.Context) {
	//TODO

}

func showOrder(c *gin.Context) {
	//TODO

}

func trackingOrder(c *gin.Context) {
	// TODO
}

func indexCustomerOrders(c *gin.Context) {
	// TODO لیست کردن سفارشات قبلی مشتری بر اساس اعتبارسنجی موبایلی پیامک
}