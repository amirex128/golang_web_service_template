package container

import (
	"backend/internal/pkg/framework/assert"
	"backend/internal/pkg/framework/bredis"
	"backend/internal/pkg/framework/elastic"
	"backend/internal/pkg/framework/mysql"
	"backend/internal/pkg/framework/rabbit"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/spf13/viper"
)

const (
	mysqlCnt   = "mysql"
	redisCnt   = "redis"
	elasticCnt = "elastic"
	rabbitCnt  = "rabbit"
)

type service map[string][]map[string]string

func RegisterServices(appName string) {
	var res service
	configStr, err := ioutil.ReadFile("./configs/services.json")
	assert.Nil(err)
	err = json.Unmarshal(configStr, &res)
	assert.Nil(err)
	mysqlCntData, ok := res[mysqlCnt]
	if ok {
		for i := range mysqlCntData {
			if ok {
				mysql.RegisterMysql(
					mysqlCntData[i]["name"],
					viper.GetString(fmt.Sprintf("%s_%s_%s", appName, mysqlCntData[i]["name"], mysqlCntData[i]["env_host"])),
					viper.GetString(fmt.Sprintf("%s_%s_%s", appName, mysqlCntData[i]["name"], mysqlCntData[i]["env_user"])),
					viper.GetString(fmt.Sprintf("%s_%s_%s", appName, mysqlCntData[i]["name"], mysqlCntData[i]["env_password"])),
					viper.GetString(fmt.Sprintf("%s_%s_%s", appName, mysqlCntData[i]["name"], mysqlCntData[i]["env_database"])),
					viper.GetInt(fmt.Sprintf("%s_%s_%s", appName, mysqlCntData[i]["name"], mysqlCntData[i]["env_port"])),
				)
			}
		}
	}
	redisCntData, ok := res[redisCnt]
	if ok {
		for i := range redisCntData {
			if ok {
				err := bredis.RegisterRedis(
					redisCntData[i]["name"],
					viper.GetString(fmt.Sprintf("%s_%s_%s", appName, redisCntData[i]["name"], redisCntData[i]["env_host"])),
					viper.GetString(fmt.Sprintf("%s_%s_%s", appName, redisCntData[i]["name"], redisCntData[i]["env_password"])),
					redisCntData[i]["kind"],
					viper.GetInt(fmt.Sprintf("%s_%s_%s", appName, redisCntData[i]["name"], redisCntData[i]["env_port"])),
					viper.GetInt(fmt.Sprintf("%s_%s_%s", appName, redisCntData[i]["name"], redisCntData[i]["env_database"])),
				)
				assert.Nil(err)
			}
		}
	}
	elasticCntData, ok := res[elasticCnt]
	if ok {
		for i := range elasticCntData {
			if ok {
				err := elastic.RegisterElastic(
					elasticCntData[i]["name"],
					viper.GetString(fmt.Sprintf("%s_%s", elasticCntData[i]["name"], elasticCntData[i]["env_host"])),
					viper.GetInt(fmt.Sprintf("%s_%s", elasticCntData[i]["name"], elasticCntData[i]["env_port"])),
				)
				assert.Nil(err)
			}
		}
	}
	rabbitCntData, ok := res[rabbitCnt]
	if ok {
		for i := range rabbitCntData {
			if ok {
				rabbit.RegisterRabbit(
					rabbitCntData[i]["name"],
					viper.GetString(fmt.Sprintf("%s_%s", rabbitCntData[i]["name"], rabbitCntData[i]["env_host"])),
					viper.GetString(fmt.Sprintf("%s_%s", rabbitCntData[i]["name"], rabbitCntData[i]["env_user"])),
					viper.GetString(fmt.Sprintf("%s_%s", rabbitCntData[i]["name"], rabbitCntData[i]["env_password"])),
					viper.GetString(fmt.Sprintf("%s_%s", rabbitCntData[i]["name"], rabbitCntData[i]["env_v_host"])),
					viper.GetInt(fmt.Sprintf("%s_%s", rabbitCntData[i]["name"], rabbitCntData[i]["env_port"])),
				)
			}
		}
	}
}
