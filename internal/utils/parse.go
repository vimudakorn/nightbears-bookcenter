package utils

import (
	"strconv"
	"strings"
)

func ParsePrice(s string) float64 {
	s = strings.ReplaceAll(s, "ราคา", "")
	s = strings.ReplaceAll(s, "บาท", "")
	s = strings.TrimSpace(s)
	f, _ := strconv.ParseFloat(s, 64)
	return f
}

func GetString(ptr *string) string {
	if ptr == nil {
		return ""
	}
	return *ptr
}
