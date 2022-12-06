package models

import (
	"github.com/amirex128/selloora_backend/internal/utils"
	"github.com/sirupsen/logrus"
)

type Guild struct {
	ID         int    `gorm:"primary_key;auto_increment" json:"id"`
	ParentID   int    `json:"parent_id"`
	Name       string `json:"name"`
	Equivalent string `json:"equivalent"`
	Sort       uint   `json:"sort"`
	Active     byte   `json:"active"`
}

type GuildProduct struct {
	GuildID   int   `json:"guild_id"`
	ProductID int64 `json:"product_id"`
}

func initGuild(manager *MysqlManager) {
	if !manager.GetConn().Migrator().HasTable(&Guild{}) {
		manager.GetConn().Migrator().CreateTable(&Guild{})
		manager.GetConn().Migrator().CreateTable(&GuildProduct{})
		var guilds [][]string
		if utils.IsTest() {
			guilds = utils.ReadCsvFile("../../../csv/guilds.csv")
		} else {
			guilds = utils.ReadCsvFile("./csv/guilds.csv")
		}
		manager.CreateAllGuilds(guilds)

	}
}

func (m *MysqlManager) CreateAllGuilds(files [][]string) {
	guilds := make([]Guild, 0)
	for i := range files {
		value := files[i]
		guilds = append(guilds, Guild{
			ID:         utils.StringToInt(value[0]),
			ParentID:   utils.StringToInt(value[1]),
			Name:       value[2],
			Equivalent: value[4],
			Sort:       utils.StringToUint(value[5]),
			Active:     utils.ActiveConvert(value[6]),
		})
	}
	err := m.GetConn().CreateInBatches(guilds, 100).Error
	if err != nil {
		logrus.Error("seed guilds error: ", err)
	}
}
