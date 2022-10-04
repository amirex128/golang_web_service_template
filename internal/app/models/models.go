package models

import (
	"backend/internal/pkg/framework/mysql"
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
	ShopTable               = "shops"
	ProvinceTable           = "provinces"
	ProductProvinceTable    = "product_provinces"
	UserTable               = "users"

	CommentTable   = "comments"
	TagTable       = "tags"
	FinancialTable = "financials"
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
	initProvince(manager)
	initCity(manager)
	initUser(manager)
	initShop(manager)
	initCustomer(manager)
	InitProduct(manager)
	initDiscount(manager)
	InitPost(manager)
	initComment(manager)
	initTag(manager)
	initAddress(manager)
	initOrder(manager)
	initOrderItem(manager)

	initEvent(manager)
	initFeatureGroup(manager)
	initFeatureItem(manager)
	initFeatureItemValue(manager)
	initGallery(manager)
	initGuild(manager)
	initManufacturer(manager)
	initOption(manager)

}

func init() {
	mysql.Register("mysql_main", &MysqlManager{})
}
