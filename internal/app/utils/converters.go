package utils

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"strconv"
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
func StringToUint32(value string) uint32 {
	val, _ := strconv.ParseUint(value, 10, 32)
	return uint32(val)
}

func Uint64ToString(value uint64) string {
	return fmt.Sprintf("%d", value)
}

func IntToString(value int) string {
	return fmt.Sprintf("%d", value)
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
