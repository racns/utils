package utils

import (
	"os"
	"path/filepath"
)

// FileList 获取指定目录下的所有文件
func FileList(path string, opt ...map[string]any) (slice []string) {

	// 默认参数
	defOpt := map[string]any{
		// 获取指定后缀的文件
		"ext": []string{"*"},
		// 包含子目录
		"sub": true,
		// 返回路径格式
		"format": "network",
		// 域名
		"domain": "",
		// 过滤前缀
		"prefix": "",
	}

	if len(opt) != 0 {
		// 合并参数
		for key, val := range defOpt {
			if opt[0][key] == nil {
				opt[0][key] = val
			}
		}
	} else {
		// 默认参数
		opt = append(opt, defOpt)
	}

	conf := opt[0]
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		// 忽略当前目录
		if info.IsDir() {
			return nil
		}
		// 忽略子目录
		if !conf["sub"].(bool) && filepath.Dir(path) != path {
			return nil
		}
		// []string 转 []any
		var exts []any
		for _, v := range conf["ext"].([]string) {
			exts = append(exts, v)
		}
		// 忽略指定后缀
		if !InArray("*", exts) && !InArray(filepath.Ext(path), exts) {
			return nil
		}
		slice = append(slice, path)
		return nil
	})

	if err != nil {
		return []string{}
	}

	// 转码为网络路径
	if conf["format"] == "network" {
		for key, val := range slice {
			slice[key] = filepath.ToSlash(val)
			if !IsEmpty(conf["domain"]) {
				root, _ := os.Getwd()
				slice[key] = conf["domain"].(string) + slice[key][len(root)+len(conf["prefix"].(string)):]
			}
		}
	}

	return
}

// FileBytes 获取文件字节
func FileBytes(path string) []byte {
	// 读取文件
	file, err := os.Open(path)
	if err != nil {
		return []byte{}
	}
	defer file.Close()
	// 获取文件大小
	info, _ := file.Stat()
	size := info.Size()
	// 读取文件
	data := make([]byte, size)
	file.Read(data)
	return data
}
