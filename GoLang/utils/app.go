package utils

func init() {
	Is.Ip           = isIp
	Is.Url          = isUrl
	Is.Email        = isEmail
	Is.Phone        = isPhone
	Is.Empty        = isEmpty
	Is.True		= isTrue
	Is.False	= isFalse
	Is.Domain       = isDomain
	Env.Toml	= envToml
	Get.Type        = getType
	In.Array        = inArray
	Array.Filter 	= arrayFilter
	Password.Create = passwordCreate
	Password.Verify = passwordVerify
}

var Is struct{
	Ip     func(ip     any) bool
	Url    func(url    any) bool
	Email  func(email  any) bool
	Phone  func(phone  any) bool
	Empty  func(value  any) bool
	True   func(value  any) bool
	False  func(value  any) bool
	Domain func(domain any) bool
}

var Env struct{
	Toml func(value ...string) any
}

var Get struct{
	Type func(value any) string
}

var In struct{
	Array func(value any, array []any) bool
}

var Array struct{
	Filter func(array []string) []string
}

var Password struct{
	Create func(password []byte) string
	Verify func(encode any, password []byte) bool
}
