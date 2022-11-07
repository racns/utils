package utils

func arrayFilter(array []string) (result []string) {
	for key, val := range array {
		if (key > 0 && array[key-1] == val) || len(val) == 0 {
			continue
		}
		result = append(result, val)
	}
	return
}