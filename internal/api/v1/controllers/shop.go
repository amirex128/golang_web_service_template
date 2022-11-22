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
)

// CreateShop
// @Summary ایجاد فروشگاه
// @description هر کاربر برای این که بتواند محصولی ایجاد کند باید فروشگاه داشته باشد تا محصولات و مقالات خود را بر روی این فروشگاه ذخیره کند این فروشگاه میتواند ربات تلگرام باشد یا سایت باشد یک نمونه مشابه اینستاگرام باشد
// @Tags shop
// @Router       /user/shop/create [post]
// @Param	Authorization	 header string	true "Authentication"
// @Param	message	 body   DTOs.CreateShop  	true "ورودی"
func CreateShop(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:createShop", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	dto, err := validations.CreateShop(c)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	userID := models.GetUser(c)

	err = models.NewMysqlManager(c).CreateShop(dto, *userID)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "فروشگاه با موفقیت ایجاد شد",
	})
}

// UpdateShop
// @Summary ویرایش فروشگاه
// @description هر کاربر برای این که بتواند محصولی ایجاد کند باید فروشگاه داشته باشد تا محصولات و مقالات خود را بر روی این فروشگاه ذخیره کند این فروشگاه میتواند ربات تلگرام باشد یا سایت باشد یک نمونه مشابه اینستاگرام باشد
// @Tags shop
// @Router       /user/shop/update [post]
// @Param	Authorization	 header string	true "Authentication"
// @Param	message	 body   DTOs.UpdateShop  	true "ورودی"
func UpdateShop(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:updateShop", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	dto, err := validations.UpdateShop(c)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}

	err = models.NewMysqlManager(c).UpdateShop(dto)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "فروشگاه با موفقیت ویرایش شد",
	})
}

// DeleteShop
// @Summary حذف فروشگاه
// @description هر کاربر برای این که بتواند محصولی ایجاد کند باید فروشگاه داشته باشد تا محصولات و مقالات خود را بر روی این فروشگاه ذخیره کند این فروشگاه میتواند ربات تلگرام باشد یا سایت باشد یک نمونه مشابه اینستاگرام باشد
// @Tags shop
// @Router       /user/shop/delete/{id} [post]
// @Param	Authorization	 header string	true "Authentication"
// @Param	id			 path   string	true "شناسه فروشگاه" SchemaExample(1)
func DeleteShop(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:deleteShop", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	shopID := utils.StringToUint64(c.Param("id"))
	userID := models.GetUser(c)
	dto, err := validations.DeleteShop(c)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	if dto.ProductBehave == "move" {
		err = models.NewMysqlManager(c).MoveProducts(shopID, dto.NewShopID, *userID)
		if err != nil {
			return
		}
	} else if dto.ProductBehave == "delete_product" {
		err = models.NewMysqlManager(c).DeleteProducts(shopID, *userID)
		if err != nil {
			return
		}
	}

	err = models.NewMysqlManager(c).DeleteShop(shopID, *userID)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "فروشگاه با موفقیت حذف شد",
	})
}

// ShowShop
// @Summary نمایش فروشگاه
// @description هر کاربر برای این که بتواند محصولی ایجاد کند باید فروشگاه داشته باشد تا محصولات و مقالات خود را بر روی این فروشگاه ذخیره کند این فروشگاه میتواند ربات تلگرام باشد یا سایت باشد یک نمونه مشابه اینستاگرام باشد
// @Tags shop
// @Router       /user/shop/show/{id} [get]
// @Param	Authorization	 header string	true "Authentication"
// @Param	id			 path   string	true "شناسه فروشگاه" SchemaExample(1)
func ShowShop(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:showShop", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	shopID := utils.StringToUint64(c.Param("id"))
	shop, err := models.NewMysqlManager(c).FindShopByID(shopID)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"shop": shop,
	})
}

// IndexShop
// @Summary لیست فروشگاه
// @description هر کاربر برای این که بتواند محصولی ایجاد کند باید فروشگاه داشته باشد تا محصولات و مقالات خود را بر روی این فروشگاه ذخیره کند این فروشگاه میتواند ربات تلگرام باشد یا سایت باشد یک نمونه مشابه اینستاگرام باشد
// @Tags shop
// @Router       /user/shop [get]
// @Param	Authorization	 header string	true "Authentication"
// @Param	search			 query   string	false "متن جستجو"
// @Param	page			 query   string	false "شماره صفحه"
// @Param	page_size		 query   string	false "تعداد صفحه"
// @Param	sort			 query   string	false "مرتب سازی براساس desc/asc"
func IndexShop(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:indexShop", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	dto, err := validations.IndexShop(c)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}

	shops, err := models.NewMysqlManager(c).GetAllShopWithPagination(dto)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"shops": shops,
	})
}

// SendPrice
// @Summary ویرایش هزینه ارسال سفارشات به صورت جدا
// @description هر کاربر برای این که بتواند محصولی ایجاد کند باید فروشگاه داشته باشد تا محصولات و مقالات خود را بر روی این فروشگاه ذخیره کند این فروشگاه میتواند ربات تلگرام باشد یا سایت باشد یک نمونه مشابه اینستاگرام باشد
// @Tags shop
// @Router       /user/shop/send-price [post]
// @Param	Authorization	 header string	true "Authentication"
// @Param	message	 body   DTOs.SendPrice  	true "ورودی"
func SendPrice(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:sendPrice", "request")
	c.Request.WithContext(ctx)
	defer span.End()
	dto, err := validations.SendPrice(c)
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	err = models.NewMysqlManager(c).UpdateShop(DTOs.UpdateShop{
		SendPrice: dto.SendPrice,
	})
	if err != nil {
		errorx.ResponseErrorx(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "هزینه ارسال با موفقیت بروزرسانی شد",
	})
}

func GetInstagramPost(c *gin.Context) {
	span, ctx := apm.StartSpan(c.Request.Context(), "controller:getInstagramPost", "request")
	c.Request.WithContext(ctx)
	defer span.End()

}
