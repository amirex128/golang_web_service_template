package timed

import (
	"github.com/amirex128/selloora_backend/internal/pkg/framework/assert"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	epoch time.Time
	once  sync.Once
	loc   *time.Location
)

func TimeStringToTimestamp(timeString string) int64 {
	l, err := time.LoadLocation("Asia/Tehran")
	if err != nil {
		panic(err)
	}
	splitted := strings.Split(timeString, ":")
	if len(splitted) > 1 {
		hour, _ := strconv.Atoi(splitted[0])
		minute, _ := strconv.Atoi(splitted[1])
		second, _ := strconv.Atoi(splitted[2])
		startTime := time.Date(epoch.Year(), epoch.Month(), epoch.Day(), hour, minute, second, 0, l)
		t := startTime.Unix()

		return t
	}

	return 0
}

func GetEpochTimestamps(t time.Time, l *time.Location) int64 {
	if l == nil {
		l = time.UTC
	}
	startTime := time.Date(epoch.Year(), epoch.Month(), epoch.Day(), t.Hour(), t.Minute(), t.Second(), 0, l)
	return startTime.Unix()
}

func GetLocation() *time.Location {
	return loc
}

func init() {
	var err error
	epoch = time.Unix(0, 0)
	loc, err = time.LoadLocation("Asia/Tehran")
	assert.Nil(err)
}
