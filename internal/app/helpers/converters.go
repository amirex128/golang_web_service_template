package helpers

import (
	"database/sql"
	"strconv"
	"time"
)

func ActiveConvert(value interface{}) byte {

	if value == "0" || value == "deactivate" || value == "" || value == "NULL" {
		return 0
	}
	return 1
}

func StringConvert(value string) sql.NullString {
	return sql.NullString{
		Valid:  !(value == "" || value == "NULL"),
		String: value,
	}
}

func Int32Convert(value string) int {
	val, _ := strconv.Atoi(value)
	return val
}
func Float32Convert(value string) float32 {
	val, _ := strconv.ParseFloat(value, 32)
	return float32(val)
}
func UintConvert(value string) uint {
	val, _ := strconv.ParseUint(value, 10, 32)
	return uint(val)
}
func Uint64Convert(value string) uint64 {
	val, _ := strconv.ParseUint(value, 10, 32)
	return val
}
func DateTimeConvert(value string) string {
	if value != "" {
		l, _ := time.LoadLocation("Asia/Tehran")
		res, err := time.ParseInLocation("2006-01-02 15:04:05", value, l)
		if err == nil {
			return res.String()
		}
		return ""
	}
	return ""
}
func NowTime() string {
	l, _ := time.LoadLocation("Asia/Tehran")
	return time.Now().In(l).Format("2006-01-02 15:04:05")
}
