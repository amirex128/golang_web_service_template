package models

import (
	"backend/internal/app/helpers"
	"database/sql"
	"encoding/gob"
	"github.com/sirupsen/logrus"
	"io"
)

type Guild struct {
	ID         int            `json:"id"`
	ParentID   int            `json:"parent_id"`
	Name       string         `json:"name"`
	Icon       sql.NullString `json:"icon"`
	Equivalent sql.NullString `json:"equivalent"`
	Sort       uint           `json:"sort"`
	Active     byte           `json:"active"`
}

type GuildProduct struct {
	GuildID   int   `json:"guild_id"`
	ProductID int64 `json:"product_id"`
}
type GuildArr []Guild

func (s GuildArr) Len() int {
	return len(s)
}
func (s GuildArr) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s GuildArr) Less(i, j int) bool {
	return s[i].ID < s[j].ID
}

func (c *Guild) Encode(iw io.Writer) error {
	return gob.NewEncoder(iw).Encode(c)
}

func (c *Guild) Decode(ir io.Reader) error {
	return gob.NewDecoder(ir).Decode(c)
}

func initGuild(manager *MysqlManager) {
	manager.GetConn().AutoMigrate(&Guild{})
	manager.GetConn().AutoMigrate(&GuildProduct{})
	guilds := helpers.ReadCsvFile("../../csv/guilds.csv")
	manager.CreateAllGuilds(guilds)
}

func (m *MysqlManager) CreateAllGuilds(files [][]string) {
	guilds := make([]Guild, 0)
	for i := range files {
		value := files[i]
		guilds = append(guilds, Guild{
			ID:         helpers.Int32Convert(value[0]),
			ParentID:   helpers.Int32Convert(value[1]),
			Name:       value[2],
			Icon:       helpers.StringConvert(value[3]),
			Equivalent: helpers.StringConvert(value[4]),
			Sort:       helpers.UintConvert(value[5]),
			Active:     helpers.ActiveConvert(value[6]),
		})
	}
	err := m.GetConn().CreateInBatches(guilds, 100).Error
	if err != nil {
		logrus.Error("seed guilds error: ", err)
	}
}
