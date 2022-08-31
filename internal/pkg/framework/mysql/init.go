package mysql

import (
	"backend/internal/pkg/framework/assert"
	"backend/internal/pkg/framework/initializer"
	"backend/internal/pkg/framework/safe"
	"backend/internal/pkg/framework/xlog"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"sync"
	"time"

	"gorm.io/gorm"
)

var (
	all     = make(map[string][]initializer.Simple, 0)
	allLock = &sync.RWMutex{}

	once = sync.Once{}

	dbMap  = make(map[string]*SingleManager)
	dbLock = &sync.RWMutex{}

	mysqlConnExpected = make([]mysqlExpected, 0)
)

type Manager struct {
}

type SingleManager struct {
	dbMap *gorm.DB
}

func (m *SingleManager) GetConn() *gorm.DB {
	return m.dbMap
}

func (m *SingleManager) GetSqlDB() *sql.DB {
	db, _ := m.dbMap.DB()
	return db
}

// Begin is for begin transaction
func (m *SingleManager) Begin() *SingleManager {
	return &SingleManager{m.dbMap.Begin()}
}

// Commit is for committing transaction. panic if transaction is not started
func (m *SingleManager) Commit() *gorm.DB {
	return m.dbMap.Commit()
}

// Rollback is for RollBack transaction. panic if transaction is not started
func (m *SingleManager) Rollback() *gorm.DB {
	return m.dbMap.Rollback()
}

type mysqlExpected struct {
	containerName string
	host          string
	port          int
	user          string
	password      string
	database      string
}

func RegisterMysql(cnt, host, user, password, database string, port int) {
	mysqlConnExpected = append(mysqlConnExpected, mysqlExpected{
		containerName: cnt,
		host:          host,
		database:      database,
		user:          user,
		password:      password,
		port:          port,
	})
}

type gorpLogger struct {
	fields logrus.Fields
}

func (g gorpLogger) Print(v ...interface{}) {
	logrus.WithFields(g.fields).Debugf("%v", v...)
}

func Initialize(ctx context.Context) {
	once.Do(func() {
		safe.Try(func() error {
			for i := range mysqlConnExpected {
				dns := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true",
					mysqlConnExpected[i].user,
					mysqlConnExpected[i].password,
					mysqlConnExpected[i].host,
					mysqlConnExpected[i].port,
					mysqlConnExpected[i].database,
				)
				db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
				if err != nil {
					xlog.GetWithError(ctx, errors.New("connect to db failed to ")).Error(dns)
					return err
				}
				dbLock.Lock()
				dbMap[mysqlConnExpected[i].containerName] = &SingleManager{
					dbMap: db,
				}
				dbLock.Unlock()
				if viper.GetBool("develop_mode") {
					//logger := gorpLogger{
					//	fields: logrus.Fields{
					//		"type": "sql",
					//	},
					//}
					dbLock.RLock()
					//dbMap[mysqlConnExpected[i].containerName].dbMap.LogMode(true).SetLogger(logger)
					dbLock.RUnlock()
				}
				dbLock.RLock()
				//err = dbMap[mysqlConnExpected[i].containerName].dbMap.DB().Ping()
				dbLock.RUnlock()
				if err != nil {
					xlog.GetWithError(ctx, errors.New("ping to db failed to ")).Error(dns)
					return err
				}

				allLock.RLock()
				dbPostInitial, ok := all[mysqlConnExpected[i].containerName]
				if ok {
					for j := range dbPostInitial {
						dbPostInitial[j].Initial()
					}
				}
				allLock.RUnlock()
				logrus.Infof("successfully connected to mysql :%s", dns)
			}
			return nil
		}, 30*time.Second)

	})
}

// Register a new initMysql module
func Register(cnt string, m ...initializer.Simple) {
	allLock.Lock()
	all[cnt] = make([]initializer.Simple, 0)
	all[cnt] = append(all[cnt], m...)
	allLock.Unlock()
}

func MustGetMysqlConn(cnt string) *SingleManager {
	dbLock.RLock()
	defer dbLock.RUnlock()
	res, ok := dbMap[cnt]
	assert.True(ok)
	assert.NotNil(res)
	return res
}
