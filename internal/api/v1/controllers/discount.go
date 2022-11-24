package controllers

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
// @Router       /user/customer/discount/check [post]
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

	discount, err := models.NewMysqlManager(c).FindDiscountByCodeAndUserID(dto.Code)
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

// CreateDiscount
// @Summary ایجاد تخفیف
// @description ایجاد یک تخفیف بر روی یک محصصول یا چند محصول
// @Tags discount
// @Router       /user/discount/create [post]
// @Param	Authorization	 header string	true "Authentication"
// @Param	message	 body   DTOs.CreateDiscount  	true "ورودی"
func CreateDiscount(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:createDiscount", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	dto, err := validations.CreateDiscount(c)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	discount, err := models.NewMysqlManager(c).CreateDiscount(dto)

	c.JSON(http.StatusOK, gin.H{
		"message": "تخفیف با موفقیت ایجاد شد",
		"data":    discount,
	})

}

// UpdateDiscount
// @Summary ویرایش تخفیف
// @description ویرایش تخفیف
// @Tags discount
// @Router       /user/discount/update/{id} [post]
// @Param	Authorization	 header string	true "Authentication"
// @Param	id			 path   string	false "شناسه تخفیف" SchemaExample(1)
// @Param	message	 body   DTOs.UpdateDiscount  	true "ورودی"
func UpdateDiscount(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:updateDiscount", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	dto, err := validations.UpdateDiscount(c)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	err = models.NewMysqlManager(c).UpdateDiscount(dto)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "تخفیف با موفقیت بروزرسانی شد",
	})
}

// IndexDiscount
// @Summary لیست تخفیف
// @description لیست تخفیفات
// @Tags discount
// @Router       /user/discount [get]
// @Param	Authorization	 header string	true "Authentication"
// @Param	search			 query   string	false "متن جستجو"
// @Param	page			 query   string	false "شماره صفحه"
// @Param	page_size		 query   string	false "تعداد صفحه"
// @Param	sort			 query   string	false "مرتب سازی براساس desc/asc"
func IndexDiscount(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:indexDiscount", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	dto, err := validations.IndexDiscount(c)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}

	discounts, err := models.NewMysqlManager(c).GetAllDiscountWithPagination(dto)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"discounts": discounts,
	})
}

// DeleteDiscount
// @Summary حذف تخفیف
// @description حذف تخفیف
// @Tags discount
// @Router       /user/discount/delete/{id} [post]
// @Param	Authorization	 header string	true "Authentication"
// @Param	id			 path   string	true "شناسه تخفیف" SchemaExample(1)
func DeleteDiscount(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:deleteDiscount", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	id := utils.StringToUint64(c.Param("id"))

	err := models.NewMysqlManager(c).DeleteDiscount(id)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "تخفیف با موفقیت حذف شد",
	})
}

// ShowDiscount
// @Summary نمایش تخفیف
// @description نمایش تخفیف
// @Tags discount
// @Router       /user/discount/show/{id} [get]
// @Param	Authorization	 header string	true "Authentication"
// @Param	id			 path   string	true "شناسه تخفیف" SchemaExample(1)
func ShowDiscount(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:showDiscount", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	id := utils.StringToUint64(c.Param("id"))
	userID := models.GetUser(c)

	discount, err := models.NewMysqlManager(c).FindDiscountById(id)
	if err != nil {
		errorx.ResponseErrorx(c, err)
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
