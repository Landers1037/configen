/*
Project: configen template_nginx.go
Created: 2021/11/19 by Landers
*/

package configen

import (
	"os"
)

// MakeTemplateNginx 生成nginx配置文件
// 支持自定义域名 服务名 端口 IP 代理转发
func MakeTemplateNginx(tpl, file string, data interface{}) error {
	f, err := os.OpenFile(file, os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return Err(ErrGenTemp, err.Error())
	}
	defer f.Close()
	return MakeTemplate(tpl, f, data)
}

func MakeTemplateNginxTest(tpl string, data interface{}) error {
	return MakeTemplate(tpl, os.Stdout, data)
}

func MakeTemplateNginxString(tpl, file string, data interface{}) error {
	f, err := os.OpenFile(file, os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return Err(ErrGenTemp, err.Error())
	}
	defer f.Close()
	return MakeTemplateString(tpl, f, data)
}
