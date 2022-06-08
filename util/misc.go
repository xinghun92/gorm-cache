package util

import "strings"

func ShouldCacheQuery(sql string) bool {
	if strings.Contains(strings.ToLower(sql), "for update") {
		return false
	}

	return true
}

func ShouldCache(tableName string, tables []string) bool {
	if len(tables) == 0 {
		return true
	}
	return ContainString(tableName, tables)
}

func ContainString(target string, slice []string) bool {
	for _, s := range slice {
		if target == s {
			return true
		}
	}
	return false
}
