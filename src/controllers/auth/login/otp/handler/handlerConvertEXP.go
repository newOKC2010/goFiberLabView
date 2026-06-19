package handlerLogin

import (
	"strconv"
	"strings"
)

func GetExpiresInSeconds(expiresIn string, defaultSeconds int) int {
	if expiresIn == "" {
		return defaultSeconds
	}

	if value, err := strconv.Atoi(expiresIn); err == nil {
		return value
	}

	unit := expiresIn[len(expiresIn)-1:]
	valueStr := expiresIn[:len(expiresIn)-1]
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return defaultSeconds
	}

	switch strings.ToLower(unit) {
	case "s":
		return value
	case "m":
		return value * 60
	case "h":
		return value * 3600
	case "d":
		return value * 86400
	case "w":
		return value * 604800
	default:
		return defaultSeconds
	}
}
