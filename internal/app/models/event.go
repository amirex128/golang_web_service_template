package models

import (
	"backend/internal/app/helpers"
	"encoding/gob"
	"github.com/sirupsen/logrus"
	"io"
	"strconv"
)

type Event struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Active    byte   `json:"active"`
	StartedAt string `json:"started_at"`
	EndedAt   string `json:"ended_at"`
}

type EventArr []Event

func (s EventArr) Len() int {
	return len(s)
}
func (s EventArr) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s EventArr) Less(i, j int) bool {
	return s[i].ID < s[j].ID
}

func (c *Event) Encode(iw io.Writer) error {
	return gob.NewEncoder(iw).Encode(c)
}

func (c *Event) Decode(ir io.Reader) error {
	return gob.NewDecoder(ir).Decode(c)
}
func initEvent(manager *MysqlManager) {
	manager.GetConn().AutoMigrate(&Event{})
	events := helpers.ReadCsvFile("../../csv/events.csv")
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
			Active:    helpers.ActiveConvert(value[2]),
			StartedAt: value[3],
			EndedAt:   value[4],
		})
	}
	err := m.GetConn().CreateInBatches(events, 100).Error
	if err != nil {
		logrus.Error("seed events error: ", err)
	}
}
