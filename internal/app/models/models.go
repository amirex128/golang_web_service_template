package models

import (
	"backend/internal/app/helpers"
	"backend/internal/pkg/framework/mysql"
	"database/sql"
	"strconv"
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
	manager.GetConn().AutoMigrate(&Category{})
	manager.GetConn().AutoMigrate(&CategoryOption{})
	manager.GetConn().AutoMigrate(&CategoryProduct{})
	manager.GetConn().AutoMigrate(&CategoryRelated{})
	manager.GetConn().AutoMigrate(&City{})
	manager.GetConn().AutoMigrate(&CityProduct{})
	manager.GetConn().AutoMigrate(&Customer{})
	manager.GetConn().AutoMigrate(&Discount{})
	manager.GetConn().AutoMigrate(&Event{})
	manager.GetConn().AutoMigrate(&FeatureGroup{})
	manager.GetConn().AutoMigrate(&FeatureItem{})
	manager.GetConn().AutoMigrate(&FeatureItemProduct{})
	manager.GetConn().AutoMigrate(&FeatureItemValue{})
	manager.GetConn().AutoMigrate(&Gallery{})
	manager.GetConn().AutoMigrate(&GroupDiscount{})
	manager.GetConn().AutoMigrate(&Guild{})
	manager.GetConn().AutoMigrate(&GuildProduct{})
	manager.GetConn().AutoMigrate(&Manufacturer{})
	manager.GetConn().AutoMigrate(&Option{})
	manager.GetConn().AutoMigrate(&OptionItem{})
	manager.GetConn().AutoMigrate(&Order{})
	manager.GetConn().AutoMigrate(&OrderItem{})
	manager.GetConn().AutoMigrate(&Product{})
	manager.GetConn().AutoMigrate(&Province{})
	manager.GetConn().AutoMigrate(&ProductProvince{})
	manager.GetConn().AutoMigrate(&User{})

	categories := helpers.ReadCsvFile("../../csv/categories.csv")
	manager.CreateAllCategories(categories)

	categoryRelated := helpers.ReadCsvFile("../../csv/category_related.csv")
	manager.CreateAllCategoryRelated(categoryRelated)

	cities := helpers.ReadCsvFile("../../csv/cities.csv")
	manager.CreateAllCities(cities)

	events := helpers.ReadCsvFile("../../csv/events.csv")
	manager.CreateAllEvents(events)

	featureGroups := helpers.ReadCsvFile("../../csv/feature_groups.csv")
	manager.CreateAllFeatureGroups(featureGroups)

	featureItemValues := helpers.ReadCsvFile("../../csv/feature_item_values.csv")
	manager.CreateAllFeatureItemValues(featureItemValues)

	featureItems := helpers.ReadCsvFile("../../csv/feature_items.csv")
	manager.CreateAllFeatureItems(featureItems)

	guilds := helpers.ReadCsvFile("../../csv/guilds.csv")
	manager.CreateAllGuilds(guilds)

	manufacturer := helpers.ReadCsvFile("../../csv/manufacturers.csv")
	manager.CreateAllManufacturer(manufacturer)

	optionItems := helpers.ReadCsvFile("../../csv/option_items.csv")
	manager.CreateAllOptionItems(optionItems)

	options := helpers.ReadCsvFile("../../csv/options.csv")
	manager.CreateAllOptions(options)

	provinces := helpers.ReadCsvFile("../../csv/provinces.csv")
	manager.CreateAllProvinces(provinces)
}

func init() {
	mysql.Register("mysql_main", &MysqlManager{})
}

func activeConvert(value string) byte {

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

func intConvert(value string) int {
	return func() int {
		val, _ := strconv.Atoi(value)
		return val
	}()
}
func uintConvert(value string) uint {
	return func() uint {
		val, _ := strconv.ParseUint(value, 10, 32)
		return uint(val)
	}()
}
