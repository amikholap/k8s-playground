package util

import (
	"log"
	"os"
	"strconv"
)

func GetEnvInt(name string, fallback int) int {
	var value int
	var err error

	valueStr := os.Getenv(name)
	if valueStr == "" {
		value = fallback
	} else {
		value, err = strconv.Atoi(valueStr)
		if err != nil {
			log.Panic(err)
		}
	}

	return value
}
