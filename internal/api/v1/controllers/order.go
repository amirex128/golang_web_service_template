package controllers

import (
	"github.com/amirex128/selloora_backend/internal/DTOs"
	"github.com/amirex128/selloora_backend/internal/constants"
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils"
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
	span, ctx := apm.StartSpan(c.Request.Context(), "createOrder", "request")
	defer span.End()
	dto, err := validations.CreateOrder(c)
	if err != nil {
		return
	}

	user, err := models.NewMysqlManager(ctx).FindUserByID(c, ctx, dto.UserID)
	if err != nil {
		return
	}

	shop, err := models.NewMysqlManager(ctx).FindShopByID(c, ctx, dto.ShopID)
	if err != nil {
		return
	}

	customer, err := models.NewMysqlManager(ctx).FindCustomerById(c, ctx, dto.CustomerID)
	if err != nil {
		return
	}

	if customer.VerifyCode != dto.VerifyCode {
		c.JSON(http.StatusBadRequest, gin.H{"error": "کد تایید صحیح نمی باشد"})
		return
	}

	discount, err := models.NewMysqlManager(ctx).FindDiscountByCodeAndUserID(c, ctx, dto.DiscountCode)
	if err != nil {
		return
	}

	if utils.DifferentWithNow(discount.StartedAt) < 0 || utils.DifferentWithNow(discount.EndedAt) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "کد تخفیف منقضی شده است"})
		return
	}

	rawProducts, err := models.NewMysqlManager(ctx).FindProductByIds(c, ctx, extractProductIDs(dto))
	if err != nil {
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
	order.ShopID = shop.ID
	order.CustomerID = customer.ID
	order.DiscountID = discount.ID
	order.IP = c.ClientIP()
	order.Status = constants.PendingPaymentOrderStatus
	order.Description = dto.Description
	order.LastUpdateStatusAt = utils.NowTime()
	order.CreatedAt = utils.NowTime()

	orderID, err := models.NewMysqlManager(ctx).CreateOrder(c, ctx, order)
	if err != nil {
		return
	}

	err = models.NewMysqlManager(ctx).CreateOrderItem(c, ctx, dto.OrderItems, orderID)
	if err != nil {
		return
	}

	err = utils.SadadPayRequest(c, 10000000, 10000.0)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "خطایی در ارتباط با درگاه پرداخت رخ داده است",
			"error":   err.Error(),
		})
		return
	}

}

func extractProductIDs(dto DTOs.CreateOrder) []uint64 {
	var productIDs []uint64
	for i := range dto.OrderItems {
		productIDs = append(productIDs, dto.OrderItems[i].ProductID)
	}
	return productIDs
}

func SadadPaymentVerify(c *gin.Context) {
	span, _ := apm.StartSpan(c.Request.Context(), "sadadPaymentVerify", "request")
	defer span.End()
	err := utils.SadadVerify(c, 1, 1000.0, 100000, "")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "خطایی در ارتباط با درگاه پرداخت رخ داده است و مبلغ پرداختی شما تا 72 ساعت آینده به حساب شما برگشت داده میشود لطفا مجدد پرداخت خود را انجام دهید",
			"error":   err.Error(),
		})
		return
	}
	//TODO ارسال پیامک خریدار کسر کردن موجودی محصول و کسر کردن موجودی کد تحفیف و ارسال پیامک ثبت سفارش برای فروشنده
	//text := fmt.Sprintf("سفارش شما با کد %d با موفقیت ثبت شد و در انتظار تایید فروشگاه میباشد", order.ID)
	//err = utils.SendSMS(c, customer.Mobile, text, true)
	//if err != nil {
	//	return
	//}
}

// IndexOrder
// @Summary لیست سفارشات
// @description لیست سفارشات بر اساس حالت های مختلف قابلیت فیلتر شدن داد
// @Tags order
// @Router       /user/order [get]
// @Param	Authorization	 header string	true "Authentication"
// @Param	order_status	 query   string	false "new,processing,returned,completed وضعیت سفارش" example(new)
// @Param	search			 query   string	false "متن جستجو"
// @Param	page			 query   string	false "شماره صفحه"
// @Param	page_size		 query   string	false "تعداد صفحه"
// @Param	sort			 query   string	false "مرتب سازی براساس desc/asc"
func IndexOrder(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "indexOrder", "request")
	defer span.End()
	orderStatus := c.Query("order_status")
	userID := models.GetUser(c)
	var orders []*models.Order
	var err error
	if orderStatus == "new" {
		orders, err = models.NewMysqlManager(ctx).GetOrders(c, ctx, *userID, []string{
			constants.PendingAcceptOrderStatus,
		})
	} else if orderStatus == "processing" {
		orders, err = models.NewMysqlManager(ctx).GetOrders(c, ctx, *userID, []string{
			constants.AcceptedOrderStatus,
			constants.PendingReceivePostOrderStatus,
			constants.ReceivedPostOrderStatus,
			constants.ReceivedCustomerOrderStatus,
			constants.PendingReceivePostReturnOrderStatus,
			constants.ReceivedPostReturnOrderStatus,
			constants.ReceivedOwnerOrderStatus,
		})
	} else if orderStatus == "returned" {
		orders, err = models.NewMysqlManager(ctx).GetOrders(c, ctx, *userID, []string{
			constants.PendingReturnOrderStatus,
			constants.AcceptedReturnOrderStatus,
			constants.RejectedReturnOrderStatus,
		})
	} else if orderStatus == "completed" {
		orders, err = models.NewMysqlManager(ctx).GetOrders(c, nil, *userID, []string{
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

// ApproveOrder
// @Summary تائید سفارش
// @description سفارشات بعد از ثبت شدن باید توسط ادمین تائید شوند و سپس به مرحله انتخاب سرویس ارسال کنند بروند
// @Tags order
// @Router       /user/order/approve/{id} [post]
// @Param	Authorization	 header string	true "Authentication"
// @Param	id			 path   string	true "شناسه سفارش" SchemaExample(1)
func ApproveOrder(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "approveOrder", "request")
	defer span.End()
	orderID := utils.StringToUint64(c.Param("id"))
	userID := models.GetUser(c)
	order, err := models.NewMysqlManager(ctx).FindOrderByID(c, ctx, orderID)
	if err != nil {
		return
	}
	if order.UserID != *userID {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "شما اجازه دسترسی به این سفارش را ندارید",
		})
		return
	}
	order.Status = constants.AcceptedOrderStatus
	err = models.NewMysqlManager(ctx).UpdateOrder(c, ctx, order)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "خطا در تایید سفارش",
		})
		return
	}
}

// CancelOrder
// @Summary کنسل کردن سفارش
// @description سفارشات میتوانند بعد از ثبت شدن یا تائید شوند یا کنسل و به مرحله انتخاب ارسال کنند روند و در انجا هم نیز امکان کنسل شدن داشته باشند
// @Tags order
// @Router       /user/order/cancel/{id} [post]
// @Param	Authorization	 header string	true "Authentication"
// @Param	id			 path   string	true "شناسه سفارش" SchemaExample(1)
func CancelOrder(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "cancelOrder", "request")
	defer span.End()
	orderID := utils.StringToUint64(c.Param("id"))
	userID := models.GetUser(c)
	order, err := models.NewMysqlManager(ctx).FindOrderByID(c, ctx, orderID)
	if err != nil {
		return
	}
	if order.UserID != *userID {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "شما اجازه دسترسی به این سفارش را ندارید",
		})
		return
	}
	order.Status = constants.CanceledOrderStatus
	err = models.NewMysqlManager(ctx).UpdateOrder(c, ctx, order)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "خطا در تایید سفارش",
		})
		return
	}
}

// SendOrder
// @Summary دریافت اطلاعات ارسال و انتخاب ارسال کننده
// @description بعد از تائید سفارش باید اطلاعات سفارش از قبلی وزن وارد شود و هزینه ارسال هر سرویس دهنده محاسبه شود و توسط ادمین انتخاب شود سرویس دهنده جهت ارسال
// @Tags order
// @Router       /user/order/send [post]
// @Param	Authorization	 header string	true "Authentication"
// @Param	message	 body   DTOs.SendOrder  	true "ورودی"
func SendOrder(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "sendOrder", "request")
	defer span.End()
	dto, err := validations.SendOrder(c)
	if err != nil {
		return
	}
	order, err := models.NewMysqlManager(ctx).FindOrderByID(c, ctx, dto.OrderID)
	if err != nil {
		return
	}
	if order.UserID != *models.GetUser(c) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "شما اجازه دسترسی به این سفارش را ندارید",
		})
		return
	}
	order.Courier = dto.Courier
	order.AddressID = dto.AddressID
	order.Status = constants.ChooseCourierOrderStatus
	order.Weight = dto.Weight
	order.PackageSize = dto.PackageSize
	order.LastUpdateStatusAt = utils.NowTime()

	err = models.NewMysqlManager(ctx).UpdateOrder(c, ctx, order)
	if err != nil {
		return
	}
	if dto.Courier == "tipax" {
		err = utils.TipaxSendOrderRequest(c)
		if err != nil {
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "اطلاعات سفارش ثبت شد در اتنظار پیک برای دریافت محصول باشید",
	})
}

// CalculateSendPrice
// @Summary دریافت اطلاعات ارسال و انتخاب ارسال کننده
// @description بعد از تائید سفارش باید اطلاعات سفارش از قبلی وزن وارد شود و هزینه ارسال هر سرویس دهنده محاسبه شود و توسط ادمین انتخاب شود سرویس دهنده جهت ارسال
// @Tags order
// @Router       /user/order/calculate [post]
// @Param	Authorization	 header string	true "Authentication"
// @Param	message	 body   DTOs.CalculateOrder  	true "ورودی"
func CalculateSendPrice(c *gin.Context) {
	span, _ := apm.StartSpan(c.Request.Context(), "calculateSendPrice", "request")
	defer span.End()
	dto, err := validations.CalculateOrder(c)
	if err != nil {
		return
	}

	err = utils.CalculateSendPriceTipax(dto)

	c.JSON(http.StatusOK, gin.H{
		"tipax": "",
	})
}

// ReturnedOrder
// @Summary ثبت درخواست مرجوعی توسط مشتری
// @description مشتری میتواند بعد از دریافت سفارش ان را مرجوع کند
// @Tags order
// @Router       /user/order/returned [post]
func ReturnedOrder(c *gin.Context) {
	span, _ := apm.StartSpan(c.Request.Context(), "returnedOrder", "request")
	defer span.End()
	//TODO

}

// AcceptReturnedOrder
// @Summary تائید درخواست مرجوعی توسط مدیر
// @description بعد از درخواست مرجوعی با این درخواست توسط ادمین بررسی شود و در صورت تائید سفارش مرجوع شود و سرویس دهنده قبلی جهت جمع آوری ارسال شود
// @Tags order
// @Router       /user/order/returned/accept [post]
func AcceptReturnedOrder(c *gin.Context) {
	span, _ := apm.StartSpan(c.Request.Context(), "acceptReturnedOrder", "request")
	defer span.End()
	//TODO

}

// ShowOrder
// @Summary نمایش جزئیات سفارش
// @description مشتری نیاز دارد سفارش خود را از طریق پنل مشتری مشاهده نماید
// @Tags order
// @Router       /user/order/show/{id} [get]
// @Param	Authorization	 header string	true "Authentication"
// @Param	id			 path   string	true "شناسه سفارش" SchemaExample(1)
func ShowOrder(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "showOrder", "request")
	defer span.End()
	orderID := utils.StringToUint64(c.Param("id"))
	order, err := models.NewMysqlManager(ctx).FindOrderWithItemByID(c, ctx, orderID)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"order": order,
	})
}

// TrackingOrder
// @Summary پیگیری وضعیت ارسال سفارش
// @description مشتری میتواند سفارش خود را پیگیری نماید و مشاهده نماید که این سفارش در چه مرحله ای به سر میبرد
// @Tags order
// @Router       /user/order/tracking/{id} [get]
// @Param	Authorization	 header string	true "Authentication"
// @Param	id			 path   string	true "شناسه سفارش" SchemaExample(1)
func TrackingOrder(c *gin.Context) {
	span, _ := apm.StartSpan(c.Request.Context(), "trackingOrder", "request")
	defer span.End()
	trackingCode := c.Param("id")
	utils.TrackingOrder(trackingCode)
}

// IndexCustomerOrders
// @Summary نمایش لیست سفارشات مشتری
// @description مشتری میتواند سفارشات خود را در یک پنل ساده مشاهده نمیاد
// @Tags order
// @Router       /user/customer [get]
// @Param	search			 query   string	false "متن جستجو"
// @Param	page			 query   string	false "شماره صفحه"
// @Param	page_size		 query   string	false "تعداد صفحه"
// @Param	sort			 query   string	false "مرتب سازی براساس desc/asc"
func IndexCustomerOrders(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "indexCustomerOrders", "request")
	defer span.End()
	dto, err := validations.IndexOrderCustomer(c)
	if err != nil {
		return
	}
	customer, err := models.NewMysqlManager(ctx).FindCustomerByMobileAndVerifyCode(c, ctx, dto.Mobile, dto.VerifyCode)
	if err != nil {
		return
	}
	orders, err := models.NewMysqlManager(ctx).FindOrdersByCustomerID(c, ctx, customer.ID)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"orders": orders,
	})
}
