package utils

// 获取数据类型
func getType(value any) string {
	attr, _ := typeof(value)
	return attr
}