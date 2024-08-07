package utils

import (
	"strconv"
	"time"
)

func Atoi(s string) (int, error) {
	return strconv.Atoi(s)
}

func ParseDate(s, layout string) (time.Time, error) {
	return time.Parse(layout, s)
}

func ParseFloat64(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}
