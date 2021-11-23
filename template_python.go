/*
Project: configen template_python.go
Created: 2021/11/19 by Landers
*/

package configen

import (
	"os"
)

// MakeTemplateGunicorn 生成nginx配置文件
// 支持端口
func MakeTemplateGunicorn(tpl, file string, data interface{}) error {
	f, err := os.OpenFile(file, os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return Err(ErrGenTemp, err.Error())
	}
	defer f.Close()
	return MakeTemplate(tpl, f, data)
}

func MakeTemplateGunicornString(tpl, file string, data interface{}) error {
	f, err := os.OpenFile(file, os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return Err(ErrGenTemp, err.Error())
	}
	defer f.Close()
	return MakeTemplateString(tpl, f, data)
}

func MakeTemplateGunicornTest(tpl string, data interface{}) error {
	return MakeTemplate(tpl, os.Stdout, data)
}
