package discount

import (
	"github.com/amirex128/selloora_backend/internal/DTOs"
	"github.com/amirex128/selloora_backend/internal/models"
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/amirex128/selloora_backend/internal/utils/errorx"
	"github.com/amirex128/selloora_backend/internal/validations"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
	"strings"
)

// CheckDiscount
// @Summary بررسی تخفیف
// @description کاربر بعد از وارد کردن محصولات به سبد خرید خود باید کد تخفیف خود را وارد نمایید تا بر روی محصولات اش اعمال شوند
// @Tags discount
// @Router       /customer/discount/check [post]
// @Param	Authorization	 header string	true "Authentication"
// @Param	message	 body   DTOs.CheckDiscount  	true "ورودی"
func CheckDiscount(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:checkDiscount", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	dto, err := validations.CheckDiscount(c)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}

	discount, err := models.NewMysqlManager(c).FindDiscountByCodeAndUserID(dto.Code, dto.UserID)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}

	if utils.DifferentWithNow(discount.StartedAt) < 0 || utils.DifferentWithNow(discount.EndedAt) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "کد تخفیف منقضی شده است"})
		return
	}

	var pIDs []uint64
	for i := range dto.ProductIDs {
		pIDs = append(pIDs, dto.ProductIDs[i].ProductID)
	}

	products, err := models.NewMysqlManager(c).FindProductByIds(pIDs)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}

	productDiscounts := strings.Split(discount.ProductIDs, ",")

	applyDiscount := utils.ApplyDiscount(productDiscounts, utils.DiscountPriceType{
		Percent: discount.Percent,
		Amount:  discount.Amount,
		Type:    discount.Type,
	}, pIDs)
	var productsCalculate []utils.ProductDiscountCalculatorType
	for i := range products {
		productsCalculate = append(productsCalculate, utils.ProductDiscountCalculatorType{
			ProductID: products[i].ID,
			Price:     products[i].Price,
			Count: func() *DTOs.ProductListDiscount {
				for j := range dto.ProductIDs {
					if dto.ProductIDs[j].ProductID == products[i].ID {
						return &dto.ProductIDs[j]
					}
				}
				return nil
			}().Count,
		})
	}
	calculateDiscountProduct := utils.CalculateDiscountProduct(applyDiscount, productsCalculate, utils.DiscountPriceType{
		Percent: discount.Percent,
		Amount:  discount.Amount,
		Type:    discount.Type,
	})

	c.JSON(http.StatusOK, gin.H{
		"result": calculateDiscountProduct,
	})
}
