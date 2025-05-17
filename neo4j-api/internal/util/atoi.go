package util

import (
	"strconv"
	"time"
)

func ParseDate(dateStr string) (time.Time, error) {
	return time.Parse("2006-01-02", dateStr)
}

func ParseInt(value string) (int, error) {
	return strconv.Atoi(value)
}
