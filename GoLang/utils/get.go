package utils

import (
	"io"
	"net"
	"net/http"
	"strings"
	"syscall"
)

// GetType - 获取数据类型
func GetType(value any) (result string) {
	result, _ = typeof(value)
	return result
}

// GetIp - 获取客户端IP
func GetIp(key ...string) (result any) {

	item := map[string]any{
		"intranet": "",
		"extranet": "",
	}

	// 替代品：https://api.ipify.org https://ipinfo.io/ip https://api.ip.sb/ip
	extranet, _ := http.Get("https://myexternalip.com/raw")
	defer extranet.Body.Close()
	content, _ := io.ReadAll(extranet.Body)
	item["extranet"] = string(content)

	conn, _ := net.Dial("udp", "8.8.8.8:80")
	defer conn.Close()
	localAddr := conn.LocalAddr().String()
	idx := strings.LastIndex(localAddr, ":")
	item["intranet"] = localAddr[0:idx]

	if len(key) > 0 {
		return item[key[0]]
	}

	return item
}

// GetResolution - 获取屏幕分辨率
func GetResolution(index int) (size int) {
	item, _, _ := syscall.NewLazyDLL(`User32.dll`).NewProc(`GetSystemMetrics`).Call(uintptr(index))
	return int(item)
}

//func GetMapStringAnyVlaue(data any, key ...string) any {
//	params := data.(map[string]any)
//	if len(key) > 0 {
//		for _, v := range key {
//			params = params[v].(map[string]any)
//		}
//	}
//}
