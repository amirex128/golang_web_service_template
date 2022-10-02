package utils

import (
	"fmt"
	ptime "github.com/yaa110/go-persian-calendar"
	"time"
)

func NowTime() string {
	l, _ := time.LoadLocation("Asia/Tehran")
	return time.Now().In(l).Format("2006-01-02 15:04:05")
}

func DateToJalaali(date string) string {
	l, _ := time.LoadLocation("Asia/Tehran")
	parse, err := time.ParseInLocation("2006-01-02 15:04:05", date, l)
	if err != nil {
		return ""
	}
	pt := ptime.New(parse)
	return fmt.Sprintf("%d %s %d   %d:%d", pt.Day(), pt.Month(), pt.Year(), pt.Hour(), pt.Minute())
}

func DifferentWithNow(date string) int64 {
	l, _ := time.LoadLocation("Asia/Tehran")
	parse, err := time.ParseInLocation("2006-01-02 15:04:05", date, l)
	if err != nil {
		return 0
	}
	return time.Now().In(l).Unix() - parse.Unix()
}
func DateTimeConvert(value string) string {
	if value != "" {
		l, _ := time.LoadLocation("Asia/Tehran")
		res, err := time.ParseInLocation("2006-01-02 15:04:05", value, l)
		if err == nil {
			return res.Format("2006-01-02 15:04:05")
		}
		return ""
	}
	return ""
}
