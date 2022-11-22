package providers

import (
	"context"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/amirex128/selloora_backend/internal/providers/safe"
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	mysqlApm "go.elastic.co/apm/module/apmgormv2/v2/driver/mysql"
	mysqlGorm "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type MysqlProvider struct {
	Conn *gorm.DB
	Ctx  *gin.Context
	Mock sqlmock.Sqlmock
}

func (m *MysqlProvider) GetConn() *gorm.DB {
	return m.Conn
}

func NewMysql(i *do.Injector, ctx context.Context) (*MysqlProvider, error) {

	mysqlIns := new(MysqlProvider)

	safe.Try(func() error {
		dns := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true",
			viper.GetString("SELLOORA_MYSQL_MAIN_USER"),
			viper.GetString("SELLOORA_MYSQL_MAIN_PASSWORD"),
			viper.GetString("SELLOORA_MYSQL_MAIN_HOST"),
			viper.GetInt("SELLOORA_MYSQL_MAIN_PORT"),
			viper.GetString("SELLOORA_MYSQL_MAIN_DB"),
		)
		db, err := gorm.Open(mysqlApm.Open(dns), &gorm.Config{})
		if err != nil {
			logrus.Errorf("failed to register mysql service")
			return err
		}
		mysqlIns.Conn = db

		return nil
	}, 30*time.Second)
	logrus.Infof("mysql service is registered")
	return mysqlIns, nil
}

func NewMysqlMock(i *do.Injector, ctx context.Context) (*MysqlProvider, error) {

	mysqlIns := new(MysqlProvider)

	safe.Try(func() error {

		db, mock, err := sqlmock.New()
		if err != nil {
			return err
		}

		if db == nil {
			return err
		}

		if mock == nil {
			return err
		}

		dialector := mysqlGorm.New(mysqlGorm.Config{
			DSN:        "sqlmock_db_0",
			DriverName: "mysql",
			Conn:       db,
		})
		gormDB, err := gorm.Open(dialector, &gorm.Config{
			SkipDefaultTransaction: true,
		})
		if err != nil {
			logrus.Errorf("failed to register mysql service")
			return err
		}

		gormDB.WithContext(ctx)
		mysqlIns.Conn = gormDB
		mysqlIns.Mock = mock

		return nil
	}, 30*time.Second)
	logrus.Infof("mysql mock service is registered")
	return mysqlIns, nil
}
