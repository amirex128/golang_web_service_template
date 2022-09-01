package models

import (
	"backend/internal/pkg/framework/mysql"
	"database/sql"
	"strconv"
	"time"
)

const (
	CategoryTable           = "categories"
	CategoryOptionTable     = "category_option"
	CategoryProductTable    = "category_product"
	CategoryRelatedTable    = "category_related"
	CityTable               = "cities"
	CityProductTable        = "city_product"
	CustomerTable           = "customers"
	DiscountTable           = "discounts"
	FeatureGroupTable       = "feature_groups"
	FeatureItemTable        = "feature_items"
	FeatureItemProductTable = "feature_item_product"
	FeatureItemValueTable   = "feature_item_values"
	GalleryTable            = "galleries"
	GroupDiscountTable      = "group_discounts"
	GuildTable              = "guilds"
	GuildProductTable       = "guild_product"
	ManufacturerTable       = "manufacturers"
	OptionTable             = "options"
	OptionValueTable        = "option_values"
	OrderTable              = "orders"
	OrderItemTable          = "order_items"
	ProductTable            = "products"
	ProvinceTable           = "provinces"
	ProductProvinceTable    = "product_provinces"
	UserTable               = "users"
)

type MysqlManager struct {
	*mysql.SingleManager
}

func NewMainManager() *MysqlManager {
	return &MysqlManager{mysql.MustGetMysqlConn("mysql_main")}
}
func (r *MysqlManager) Initial() {
	manager := NewMainManager()

	if !initCategory(manager) {
		return
	}
	initCity(manager)
	initCustomer(manager)
	initDiscount(manager)
	initEvent(manager)
	initFeatureGroup(manager)
	initFeatureItem(manager)
	initFeatureItemValue(manager)
	initGallery(manager)
	initGroupDiscount(manager)
	initGuild(manager)
	initManufacturer(manager)
	initOption(manager)
	initOrder(manager)
	initOrderItem(manager)
	InitProduct(manager)
	initProvince(manager)
	initUser(manager)

}

func init() {
	mysql.Register("mysql_main", &MysqlManager{})
}

func activeConvert(value interface{}) byte {

	if value == "0" || value == "deactivate" || value == "" || value == "NULL" {
		return 0
	}
	return 1
}

func stringConvert(value string) sql.NullString {
	return sql.NullString{
		Valid:  !(value == "" || value == "NULL"),
		String: value,
	}
}

func int32Convert(value string) int {
	val, _ := strconv.Atoi(value)
	return val
}
func float32Convert(value string) float32 {
	val, _ := strconv.ParseFloat(value, 32)
	return float32(val)
}
func uintConvert(value string) uint {
	val, _ := strconv.ParseUint(value, 10, 32)
	return uint(val)
}
func dateTimeConvert(value string) string {
	if value != "" {
		l, _ := time.LoadLocation("Asia/Tehran")
		res, err := time.ParseInLocation("2006-01-02 15:04:05", value, l)
		if err == nil {
			return res.String()
		}
		return ""
	}
	return ""
}
func nowTime() string {
	l, _ := time.LoadLocation("Asia/Tehran")
	return time.Now().In(l).Format("2006-01-02 15:04:05")
}
