package models

import (
	"context"
	"github.com/amirex128/selloora_backend/internal/providers"
	"github.com/samber/do"
	"github.com/sirupsen/logrus"
)

type MysqlManager struct {
	*providers.MysqlProvider
}
type RedisManager struct {
	*providers.RedisProvider
}

var (
	redisManager     *RedisManager
	mysqlManager     *MysqlManager
	mysqlMockManager *MysqlManager
)

func Initialize(ctx context.Context) {
	if mysqlManager == nil {
		providers.Initialize(ctx)
		mysqlProvider, err := do.InvokeNamed[*providers.MysqlProvider](providers.Injector, "main_mysql")
		if err != nil {
			panic(err)
		}

		mysqlManager = &MysqlManager{mysqlProvider}
		providers.Once.Do(func() {
			mysqlManager.initializeTables()
		})
	}

}
func NewRedisManager(ctx context.Context) *RedisManager {
	redisManager.Conn = redisManager.Conn.WithContext(ctx)
	redisManager.Ctx = ctx
	return redisManager
}

func NewMysqlManager(ctx context.Context) *MysqlManager {
	mysqlManager.Conn = mysqlManager.Conn.WithContext(ctx)
	mysqlManager.Ctx = ctx
	return mysqlManager
}

func NewMysqlMockManager() *MysqlManager {
	if mysqlMockManager == nil {
		providers.Initialize(context.Background())
		mysqlMockProvider, err := do.InvokeNamed[*providers.MysqlProvider](providers.Injector, "main_mysql_mock")
		if err != nil {
			panic(err)
		}

		mysqlMockManager = &MysqlManager{mysqlMockProvider}
	}

	return mysqlMockManager
}

func (m *MysqlManager) initializeTables() {
	logrus.Info("mysql initialized started")
	defer logrus.Info("mysql initialized finished")
	manager := NewMysqlManager(context.Background())
	if !initCategory(manager) {
		return
	}
	initGallery(manager)
	initProvince(manager)
	initCity(manager)
	initUser(manager)
	initTheme(manager)
	initShop(manager)
	initPage(manager)
	initMenu(manager)
	initSlider(manager)
	initPage(manager)
	initDomain(manager)
	initCustomer(manager)
	InitProduct(manager)
	initDiscount(manager)
	InitPost(manager)
	InitComment(manager)
	InitTicket(manager)
	initTag(manager)
	initAddress(manager)
	initOrder(manager)
	initOrderItem(manager)

	initEvent(manager)
	initFeatureGroup(manager)
	initFeatureItem(manager)
	initFeatureItemValue(manager)
	initGuild(manager)
	initManufacturer(manager)
	initOption(manager)
}
