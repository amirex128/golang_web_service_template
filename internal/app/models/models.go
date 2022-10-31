package models

import (
	"backend/internal/pkg/framework/mysql"
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
