package utils

import "os"

type Mode = string

const Production Mode = "production"

func GetMode() string {
	return os.Getenv("APP_MODE")
}
