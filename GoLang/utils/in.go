package utils

func inArray(value any, array []any) bool {
	for _, v := range array {
		if v == value {
			return true
		}
	}
	return false
}