package order

import (
	"github.com/amirex128/selloora_backend/internal/DTOs"
	"github.com/amirex128/selloora_backend/internal/constants"
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/amirex128/selloora_backend/internal/validations"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
	"strings"
)

// CreateOrder
// @Summary ایجاد سفارش جدید
// @description از این سرویس برای ایجاد سفارش در بخش ادمین و کاربر استفاده میشود
// @Tags order
// @Router       /user/order/create [post]
// @Param	Authorization	 header string	true "Authentication"
// @Param	message	 body   DTOs.CreateOrder  	true "ورودی"
func CreateOrder(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:createOrder", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	dto, err := validations.CreateOrder(c)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}

	user, err := models.NewMysqlManager(c).FindUserByID(dto.UserID)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}

	shop, err := models.NewMysqlManager(c).FindShopByID(dto.ShopID)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}

	customer, err := models.NewMysqlManager(c).FindCustomerById(dto.CustomerID)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}

	if customer.VerifyCode != dto.VerifyCode {
		c.JSON(http.StatusBadRequest, gin.H{"error": "کد تایید صحیح نمی باشد"})
		return
	}

	discount, err := models.NewMysqlManager(c).FindDiscountByCodeAndUserID(dto.DiscountCode, dto.UserID)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}

	if utils.DifferentWithNow(discount.StartedAt) < 0 || utils.DifferentWithNow(discount.EndedAt) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "کد تخفیف منقضی شده است"})
		return
	}

	rawProducts, err := models.NewMysqlManager(c).FindProductByIds(extractProductIDs(dto))
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	for i := range rawProducts {
		var count uint32
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
	for i := range rawProducts {
		productCalculate = append(productCalculate, utils.ProductDiscountCalculatorType{
			ProductID: rawProducts[i].ID,
			Price:     rawProducts[i].Price,
			Count: func() *DTOs.OrderItem {
				for j := range dto.OrderItems {
					if dto.OrderItems[j].ProductID == rawProducts[i].ID {
						return &dto.OrderItems[j]
					}
				}
				return nil
			}().Count,
		})
	}
	calculateDiscountProduct := utils.CalculateDiscountProduct(applyDiscount, productCalculate, utils.DiscountPriceType{
		Percent: discount.Percent,
		Amount:  discount.Amount,
		Type:    discount.Type,
	})

	var order *models.Order
	for _, dis := range calculateDiscountProduct {
		order.TotalProductPrice += dis.RawPrice
		order.TotalDiscountPrice += dis.OffPrice
		order.TotalTaxPrice += dis.NewPrice * 0.09
		order.TotalProductDiscountPrice += dis.NewPrice
	}

	order.TotalFinalPrice = order.TotalProductDiscountPrice + order.TotalTaxPrice + order.SendPrice

	order.SendPrice = shop.SendPrice
	order.UserID = user.ID
	order.ShopID = shop.ID
	order.CustomerID = customer.ID
	order.DiscountID = discount.ID
	order.IP = c.ClientIP()
	order.Status = constants.PendingPaymentOrderStatus
	order.Description = dto.Description
	order.LastUpdateStatusAt = utils.NowTime()
	order.CreatedAt = utils.NowTime()

	order, err = models.NewMysqlManager(c).CreateOrder(order)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}

	err = models.NewMysqlManager(c).CreateOrderItem(dto.OrderItems, order.ID)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}

	err = utils.SadadPayRequest(c, 10000000, 10000.0)
	if err != nil {
		errorx.ResponseErrorx(c, errorx.New("خطایی در ارتباط با درگاه پرداخت رخ داده است", "request", err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "سفارش با موفقیت ثبت شد",
		"data":    order,
	})
}
func extractProductIDs(dto DTOs.CreateOrder) []uint64 {
	var productIDs []uint64
	for i := range dto.OrderItems {
		productIDs = append(productIDs, dto.OrderItems[i].ProductID)
	}
	return productIDs
}
