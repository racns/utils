package utils

import (
	"github.com/spf13/viper"
	"log"
	"strings"
)

func envToml(value ...string) any {

	key, def, path := "app", "", "./config/app.toml"

	if len(value) > 0 {
		for k, v := range value {
			if k == 0 {
				key = v
			} else if k == 1 {
				def = v
			} else {
				path = v
			}
		}
	}

	keys := strings.Split(key, ".")

	// 文件路径
	viper.SetConfigFile(path)

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatal("读取配置文件失败: %v", err)
	}

	if len(keys) == 1 {

		result := viper.GetStringMap(key)

		if empty := Is.Empty(result); empty {
			return def
		}

		return result

	} else {

		result := viper.GetStringMap(keys[0])[keys[1]]

		if empty := Is.Empty(result); empty {
			return def
		}

		return result
	}
}