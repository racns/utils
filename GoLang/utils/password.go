package utils

import (
	"github.com/spf13/cast"
	"golang.org/x/crypto/bcrypt"
)

func passwordCreate(password []byte) string {

	result, _ := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)

	return string(result)
}

// passwordVerify 验证密码
func passwordVerify(encode any, password []byte) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(cast.ToString(encode)), password); err != nil {
		return false
	}
	return true
}