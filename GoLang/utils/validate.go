package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spf13/cast"
	"reflect"
	"strings"
)

type validator struct {
	model   any
	message map[string]string
	rule    func(name string, value any, rule string, message map[string]string) error
}

// Validate - 验证器入口
var Validate = func(model any) *validator {
	return &validator{
		model: model,
		rule:  ValidateRules,
	}
}

// Validate - 验证模型
func (this *validator) Validate(model any) *validator {
	this.model = model
	return this
}

// Message - 自定义错误信息
func (this *validator) Message(message map[string]string) *validator {
	this.message = message
	return this
}

// Rule - 自定义校验规则
func (this *validator) Rule(callback func(name string, value any, rule string, message map[string]string) error) *validator {
	this.rule = callback
	return this
}

// Check - 验证
func (this *validator) Check(params any) (err error) {

	if this.model == nil {
		return errors.New("model is nil")
	}

	// if this.rule == nil {
	// 	this.rule = ValidateRules
	// }

	// params 转 string
	data, _ := json.Marshal(params)

	// string 转 map
	var tempMap map[string]any
	err = json.Unmarshal(data, &tempMap)

	if err != nil {
		return err
	}

	// 获取结构体的字段和标签
	for i := 0; i < reflect.TypeOf(this.model).NumField(); i++ {

		field := reflect.TypeOf(this.model).Field(i)
		tag := field.Tag
		rule := tag.Get("rule")

		// 字段名
		var name string
		// 字段值
		var value any
		if !Is.Empty(cast.ToString(tag.Get("json"))) {
			name = tag.Get("json")
		} else {
			name = field.Name
		}
		value = tempMap[name]

		if Is.Empty(rule) {
			// 结束本次循环，进入下一次循环
			continue
		}

		err := this.rule(name, value, rule, this.message)
		if err != nil {
			return err
		}
	}

	return nil
}

// ValidateRules - 验证规则
/**
 * @rule - 内置规则 - 如下：
 * required：必填
 * min：最小值
 * max：最大值
 * email：是否为邮箱
 * number：是否为数字
 * float：是否为浮点数
 * bool：是否为布尔值
 * accepted：验证某个字段是否为为 yes, on, 或是 1
 * date：是否为日期
 * alpha：只能包含字母
 * alphaNum：只能包含字母和数字
 * alphaDash：只能包含字母、数字和下划线_及破折号-
 * chs：只能包含汉字
 * chsAlpha：只能包含汉字、字母
 * chsAlphaNum：只能包含汉字、字母和数字
 * chsDash：只能是汉字、字母、数字和下划线_及破折号-
 * cntrl：是否为控制字符 - （换行、缩进、空格）
 * graph：是否为可见字符 - （除空格外的所有可打印字符）
 * lower：是否为小写字母
 * upper：是否为大写字母
 * space：是否为空白字符 - （空格、制表符、换页符等）
 * xdigit：是否为十六进制字符 - （0-9、a-f、A-F）
 * activeUrl：是否为有效的域名或者IP
 * url：是否为有效的URL地址
 * ip：是否为IP地址
 * mobile：是否为手机号
 * idCard：是否为身份证号
 * MacAddr：是否为MAC地址
 * zip：是否为邮政编码
 **/
func ValidateRules(name string, value any, rule string, message map[string]string) (err error) {

	// 获取 rule 中的规则 - 字符串切片，逗号分隔
	rules := strings.Split(rule, ",")

	// 判断 rules 是否包含 required
	if !InArray("required", cast.ToSlice(rules)) {
		// 判断 value 是否为空
		if IsEmpty(value) {
			// 结束本次循环，进入下一次循环
			return nil
		}
	}

	for _, val := range rules {
		if strings.Contains(val, "=") {
			// 获取规则和参数
			ruleAndParam := strings.Split(val, "=")
			first := ruleAndParam[0]
			second := ruleAndParam[1]

			switch first {
			case "min":
				// 判断长度是否小于最小值
				if second != "" && len(cast.ToString(value)) < cast.ToInt(second) {
					msg := message[name+"."+first]
					if msg == "" {
						msg = name + " length cannot be less than " + second + "！"
					}
					return fmt.Errorf(msg)
				}
			case "max":
				// 判断长度是否大于最大值
				if second != "" && len(cast.ToString(value)) > cast.ToInt(second) {
					msg := message[name+"."+first]
					if msg == "" {
						msg = name + " length cannot be greater than " + second + "！"
					}
					return fmt.Errorf(msg)
				}
			}

		} else {

			switch val {
			case "required":
				if IsEmpty(value) {
					msg := message[name+"."+val]
					if msg == "" {
						msg = name + " is not empty!"
					}
					return fmt.Errorf(msg)
				}
			case "email":
				if !IsEmail(value) {
					msg := message[name+"."+val]
					if msg == "" {
						msg = name + " format is incorrect!"
					}
					return fmt.Errorf(msg)
				}
			case "number":
				if !IsNumber(value) {
					msg := message[name+"."+val]
					if msg == "" {
						msg = name + " format is incorrect!"
					}
					return fmt.Errorf(msg)
				}
			case "float":
				if !IsFloat(value) {
					msg := message[name+"."+val]
					if msg == "" {
						msg = name + " format is incorrect!"
					}
					return fmt.Errorf(msg)
				}
			case "bool":
				if !IsBool(value) {
					msg := message[name+"."+val]
					if msg == "" {
						msg = name + " format is incorrect!"
					}
					return fmt.Errorf(msg)
				}
			case "accepted":
				if !IsAccepted(value) {
					msg := message[name+"."+val]
					if msg == "" {
						msg = name + " format is incorrect!"
					}
					return fmt.Errorf(msg)
				}
			case "date":
				if !IsDate(value) {
					msg := message[name+"."+val]
					if msg == "" {
						msg = name + " format is incorrect!"
					}
					return fmt.Errorf(msg)
				}
			case "alpha":
				if !IsAlpha(value) {
					msg := message[name+"."+val]
					if msg == "" {
						msg = name + " format is incorrect!"
					}
					return fmt.Errorf(msg)
				}
			case "alphaNum":
				if !IsAlphaNum(value) {
					msg := message[name+"."+val]
					if msg == "" {
						msg = name + " format is incorrect!"
					}
					return fmt.Errorf(msg)
				}
			case "alphaDash":
				if !IsAlphaDash(value) {
					msg := message[name+"."+val]
					if msg == "" {
						msg = name + " format is incorrect!"
					}
					return fmt.Errorf(msg)
				}
			case "chs":
				if !IsChs(value) {
					msg := message[name+"."+val]
					if msg == "" {
						msg = name + " format is incorrect!"
					}
					return fmt.Errorf(msg)
				}
			case "chsAlpha":
				if !IsChsAlpha(value) {
					msg := message[name+"."+val]
					if msg == "" {
						msg = name + " format is incorrect!"
					}
					return fmt.Errorf(msg)
				}
			case "chsAlphaNum":
				if !IsChsAlphaNum(value) {
					msg := message[name+"."+val]
					if msg == "" {
						msg = name + " format is incorrect!"
					}
					return fmt.Errorf(msg)
				}
			case "chsDash":
				if !IsChsDash(value) {
					msg := message[name+"."+val]
					if msg == "" {
						msg = name + " format is incorrect!"
					}
					return fmt.Errorf(msg)
				}
			case "cntrl":
				if !IsCntrl(value) {
					msg := message[name+"."+val]
					if msg == "" {
						msg = name + " format is incorrect!"
					}
					return fmt.Errorf(msg)
				}
			case "graph":
				if !IsGraph(value) {
					msg := message[name+"."+val]
					if msg == "" {
						msg = name + " format is incorrect!"
					}
					return fmt.Errorf(msg)
				}
			case "lower":
				if !IsLower(value) {
					msg := message[name+"."+val]
					if msg == "" {
						msg = name + " format is incorrect!"
					}
					return fmt.Errorf(msg)
				}
			case "upper":
				if !IsUpper(value) {
					msg := message[name+"."+val]
					if msg == "" {
						msg = name + " format is incorrect!"
					}
					return fmt.Errorf(msg)
				}
			case "space":
				if !IsSpace(value) {
					msg := message[name+"."+val]
					if msg == "" {
						msg = name + " format is incorrect!"
					}
					return fmt.Errorf(msg)
				}
			case "xdigit":
				if !IsXdigit(value) {
					msg := message[name+"."+val]
					if msg == "" {
						msg = name + " format is incorrect!"
					}
					return fmt.Errorf(msg)
				}
			case "activeUrl":
				if !IsActiveUrl(value) {
					msg := message[name+"."+val]
					if msg == "" {
						msg = name + " format is incorrect!"
					}
					return fmt.Errorf(msg)
				}
			case "url":
				if !IsUrl(value) {
					msg := message[name+"."+val]
					if msg == "" {
						msg = name + " format is incorrect!"
					}
					return fmt.Errorf(msg)
				}
			case "ip":
				if !IsIp(value) {
					msg := message[name+"."+val]
					if msg == "" {
						msg = name + " format is incorrect!"
					}
					return fmt.Errorf(msg)
				}
			case "mobile":
				if !IsMobile(value) {
					msg := message[name+"."+val]
					if msg == "" {
						msg = name + " format is incorrect!"
					}
					return fmt.Errorf(msg)
				}
			case "idCard":
				if !IsIdCard(value) {
					msg := message[name+"."+val]
					if msg == "" {
						msg = name + " format is incorrect!"
					}
					return fmt.Errorf(msg)
				}
			case "MacAddr":
				if !IsMacAddr(value) {
					msg := message[name+"."+val]
					if msg == "" {
						msg = name + " format is incorrect!"
					}
					return fmt.Errorf(msg)
				}
			case "zip":
				if !IsZip(value) {
					msg := message[name+"."+val]
					if msg == "" {
						msg = name + " format is incorrect!"
					}
					return fmt.Errorf(msg)
				}
			}
		}
	}

	return nil
}
