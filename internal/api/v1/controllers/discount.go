package controllers

import (
	"github.com/amirex128/selloora_backend/internal/DTOs"
	"github.com/amirex128/selloora_backend/internal/models"
	utils2 "github.com/amirex128/selloora_backend/internal/utils"
	"github.com/amirex128/selloora_backend/internal/validations"
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

	discount, err := models.NewMysqlManager(ctx).FindDiscountByCodeAndUserID(c, ctx, dto.Code)
	if err != nil {
		return
	}

	if utils2.DifferentWithNow(discount.StartedAt) < 0 || utils2.DifferentWithNow(discount.EndedAt) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "کد تخفیف منقضی شده است"})
		return
	}

	var pIDs []uint64
	for i := range dto.ProductIDs {
		pIDs = append(pIDs, dto.ProductIDs[i].ProductID)
	}

	products, err := models.NewMysqlManager(ctx).FindProductByIds(c, ctx, pIDs)
	if err != nil {
		return
	}

	productDiscounts := strings.Split(discount.ProductIDs, ",")

	applyDiscount := utils2.ApplyDiscount(productDiscounts, utils2.DiscountPriceType{
		Percent: discount.Percent,
		Amount:  discount.Amount,
		Type:    discount.Type,
	}, pIDs)
	var productsCalculate []utils2.ProductDiscountCalculatorType
	for i := range products {
		productsCalculate = append(productsCalculate, utils2.ProductDiscountCalculatorType{
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
	calculateDiscountProduct := utils2.CalculateDiscountProduct(applyDiscount, productsCalculate, utils2.DiscountPriceType{
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
	err = models.NewMysqlManager(ctx).CreateDiscount(c, ctx, dto)

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
	err = models.NewMysqlManager(ctx).UpdateDiscount(c, ctx, dto)
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

	discounts, err := models.NewMysqlManager(ctx).GetAllDiscountWithPagination(c, ctx, dto)
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
	id := utils2.StringToUint64(c.Param("id"))

	err := models.NewMysqlManager(ctx).DeleteDiscount(c, ctx, id)
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
	id := utils2.StringToUint64(c.Param("id"))
	userID := models.GetUser(c)

	discount, err := models.NewMysqlManager(ctx).FindDiscountById(c, ctx, id)
	if err != nil {
		return
	}
	if *discount.UserID != *userID {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "شما اجازه مشاهده این تخفیف را ندارید",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"discount": discount,
	})
}
