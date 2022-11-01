package models

import (
	"backend/internal/app/utils"
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
	manager.GetConn().AutoMigrate(&Guild{})
	manager.GetConn().AutoMigrate(&GuildProduct{})
	guilds := utils.ReadCsvFile("../../csv/guilds.csv")
	manager.CreateAllGuilds(guilds)
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
