package utils

import (
	"github.com/spf13/cast"
	"strings"
)

// FormatQuery 转 Query 格式
func FormatQuery(data any) (result string) {
	var query string
	for key, val := range cast.ToStringMap(data) {
		query += key + "=" + cast.ToString(val) + "&"
	}
	return strings.TrimRight(query, "&")
}
