package controllers

import (
	"backend/internal/app/DTOs"
	"backend/internal/app/models"
	"backend/internal/app/utils"
	"backend/internal/app/validations"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
	"strings"
)

func CheckDiscount(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "checkDiscount", "request")
	defer span.End()
	dto, err := validations.CheckDiscount(c)
	if err != nil {
		return
	}

	discount, err := models.NewMainManager().FindDiscountByCodeAndUserID(c, ctx, dto.Code)
	if err != nil {
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

	products, err := models.NewMainManager().FindProductByIds(c, ctx, pIDs)
	if err != nil {
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

func CreateDiscount(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "createDiscount", "request")
	defer span.End()
	dto, err := validations.CreateDiscount(c)
	if err != nil {
		return
	}
	err = models.NewMainManager().CreateDiscount(c, ctx, dto)

	c.JSON(http.StatusOK, gin.H{
		"message": "تخفیف با موفقیت ایجاد شد",
	})

}

func UpdateDiscount(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "updateDiscount", "request")
	defer span.End()
	dto, err := validations.UpdateDiscount(c)
	if err != nil {
		return
	}
	err = models.NewMainManager().UpdateDiscount(c, ctx, dto)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "تخفیف با موفقیت بروزرسانی شد",
	})
}

func IndexDiscount(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "indexDiscount", "request")
	defer span.End()
	dto, err := validations.IndexDiscount(c)
	if err != nil {
		return
	}

	discounts, err := models.NewMainManager().GetAllDiscountWithPagination(c, ctx, dto)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"discounts": discounts,
	})
}

func DeleteDiscount(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "deleteDiscount", "request")
	defer span.End()
	id := utils.StringToUint64(c.Param("id"))

	err := models.NewMainManager().DeleteDiscount(c, ctx, id)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "تخفیف با موفقیت حذف شد",
	})
}

func ShowDiscount(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "showDiscount", "request")
	defer span.End()
	id := utils.StringToUint64(c.Param("id"))
	userID := models.GetUser(c)

	discount, err := models.NewMainManager().FindDiscountById(c, ctx, id)
	if err != nil {
		return
	}
	if discount.UserID != userID {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "شما اجازه مشاهده این تخفیف را ندارید",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"discount": discount,
	})
}
