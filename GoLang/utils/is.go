package utils

import (
	"github.com/spf13/cast"
	"regexp"
)

// 是否为邮箱
func isEmail(email any) bool {
	return regexp.MustCompile(`^([a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+(.[a-zA-Z0-9_-])+`).MatchString(cast.ToString(email))
}

// 是否为手机号
func isPhone(phone any) bool {
	return regexp.MustCompile(`^1[3456789]\d{9}$`).MatchString(cast.ToString(phone))
}

// 是否为空
func isEmpty(value any) bool {
	_, empty := typeof(value)
	return empty
}

// 是否为域名
func isDomain(domain any) bool {
	return regexp.MustCompile(`^((https|http|ftp|rtsp|mms)?:\/\/)[^\s]+`).MatchString(cast.ToString(domain))
}

// 是否为真
func isTrue(value any) bool {
	return cast.ToBool(value)
}

// 是否为假
func isFalse(value any) bool {
	return !cast.ToBool(value)
}

// 是否为IP
func isIp(ip any) bool {
	return regexp.MustCompile(`^((2[0-4]\d|25[0-5]|[01]?\d\d?)($|(?!\.$)\.)){4}$`).MatchString(cast.ToString(ip))
}

// 是否为URL
func isUrl(url any) bool {
	return regexp.MustCompile(`^((https|http|ftp|rtsp|mms)?:\/\/)[^\s]+`).MatchString(cast.ToString(url))
}