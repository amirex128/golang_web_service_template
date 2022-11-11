package controllers

import (
	"backend/api/v1/validations"
	"backend/internal/app/DTOs"
	"backend/internal/app/models"
	"backend/internal/app/utils"
	"github.com/gin-gonic/gin"
	"go.elastic.co/apm/v2"
	"net/http"
	"strings"
)

func checkDiscount(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "checkDiscount", "request")
	defer span.End()
	dto, err := validations.CheckDiscount(c)
	if err != nil {
		return
	}

	discount, err := models.NewMainManager().FindDiscountByCodeAndUserID(c, ctx, dto.Code, dto.UserID)
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

func createDiscount(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "createDiscount", "request")
	defer span.End()
	dto, err := validations.CreateDiscount(c)
	if err != nil {
		return
	}

	userID := models.GetUser(c)

	err = models.NewMainManager().CreateDiscount(c, ctx, dto, userID)

	c.JSON(http.StatusOK, gin.H{
		"message": "تخفیف با موفقیت ایجاد شد",
	})

}

func updateDiscount(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "updateDiscount", "request")
	defer span.End()
	dto, err := validations.UpdateDiscount(c)
	if err != nil {
		return
	}
	userID := models.GetUser(c)
	discountID := c.Param("id")

	err = models.NewMainManager().UpdateDiscount(c, ctx, dto, userID, discountID)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "تخفیف با موفقیت بروزرسانی شد",
	})
}

func indexDiscount(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "indexDiscount", "request")
	defer span.End()
	dto, err := validations.IndexDiscount(c)
	if err != nil {
		return
	}
	userID := models.GetUser(c)

	discounts, err := models.NewMainManager().GetAllDiscountWithPagination(c, ctx, dto, userID)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"discounts": discounts,
	})
}

func deleteDiscount(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "deleteDiscount", "request")
	defer span.End()
	id := utils.StringToUint64(c.Param("id"))
	userID := models.GetUser(c)

	err := models.NewMainManager().DeleteDiscount(c, ctx, id, userID)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "تخفیف با موفقیت حذف شد",
	})
}

func showDiscount(c *gin.Context) {
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
