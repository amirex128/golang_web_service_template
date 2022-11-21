package models

import (
	utils2 "github.com/amirex128/selloora_backend/internal/utils"
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
	guilds := utils2.ReadCsvFile("./csv/guilds.csv")
	manager.CreateAllGuilds(guilds)
}

func (m *MysqlManager) CreateAllGuilds(files [][]string) {
	guilds := make([]Guild, 0)
	for i := range files {
		value := files[i]
		guilds = append(guilds, Guild{
			ID:         utils2.StringToInt(value[0]),
			ParentID:   utils2.StringToInt(value[1]),
			Name:       value[2],
			Equivalent: value[4],
			Sort:       utils2.StringToUint(value[5]),
			Active:     utils2.ActiveConvert(value[6]),
		})
	}
	err := m.GetConn().CreateInBatches(guilds, 100).Error
	if err != nil {
		logrus.Error("seed guilds error: ", err)
	}
}
