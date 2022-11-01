package models

import (
	"backend/internal/app/utils"
	"github.com/sirupsen/logrus"
	"strconv"
)

type Event struct {
	ID        int    `gorm:"primary_key;auto_increment" json:"id"`
	Name      string `json:"name"`
	Active    byte   `json:"active"`
	StartedAt string `json:"started_at"`
	EndedAt   string `json:"ended_at"`
}

func initEvent(manager *MysqlManager) {
	manager.GetConn().AutoMigrate(&Event{})
	events := utils.ReadCsvFile("../../csv/events.csv")
	manager.CreateAllEvents(events)
}
func (m *MysqlManager) CreateAllEvents(files [][]string) {
	events := make([]Event, 0)
	for i := range files {
		value := files[i]
		events = append(events, Event{
			ID: func() int {
				val, _ := strconv.Atoi(value[0])
				return val
			}(),
			Name:      value[1],
			Active:    utils.ActiveConvert(value[2]),
			StartedAt: value[3],
			EndedAt:   value[4],
		})
	}
	err := m.GetConn().CreateInBatches(events, 100).Error
	if err != nil {
		logrus.Error("seed events error: ", err)
	}
}
