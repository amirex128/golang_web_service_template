package models

import (
	"context"
	"github.com/amirex128/selloora_backend/internal/providers"
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
	"github.com/sirupsen/logrus"
	"net/http/httptest"
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
	rManager := redisManager
	rManager.Conn = rManager.Conn.WithContext(ctx)
	rManager.Ctx = ctx
	return rManager
}

func NewMysqlManager(ctx *gin.Context) *MysqlManager {
	mManager := mysqlManager
	mManager.Conn = mManager.Conn.WithContext(ctx)
	mManager.Ctx = ctx
	return mManager
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
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Request = httptest.NewRequest("GET", "/", nil)
	manager := NewMysqlManager(c)
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
