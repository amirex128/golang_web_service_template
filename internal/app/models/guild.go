package models

import (
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

func (m *MysqlManager) CreateAllGuilds(files [][]string) {
	guilds := make([]Guild, 0)
	for i := range files {
		value := files[i]
		guilds = append(guilds, Guild{
			ID:         intConvert(value[0]),
			ParentID:   intConvert(value[1]),
			Name:       value[2],
			Icon:       stringConvert(value[3]),
			Equivalent: stringConvert(value[4]),
			Sort:       uintConvert(value[5]),
			Active:     activeConvert(value[6]),
		})
	}
	err := m.GetConn().CreateInBatches(guilds, 100).Error
	if err != nil {
		logrus.Error("seed guilds error: ", err)
	}
}
