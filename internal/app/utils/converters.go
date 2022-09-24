package utils

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

func StringToInt(value string) int {
	val, _ := strconv.Atoi(value)
	return val
}
func StringToFloat32(value string) float32 {
	val, _ := strconv.ParseFloat(value, 32)
	return float32(val)
}
func StringToUint(value string) uint {
	val, _ := strconv.ParseUint(value, 10, 32)
	return uint(val)
}
func StringToUint64(value string) uint64 {
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
