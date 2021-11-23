/*
Project: configen template.go
Created: 2021/11/19 by Landers
*/

package configen

import (
	"io"
	"text/template"
)

// 配置文件模板生成 支持py的gunicorn和nginx
const (
	TpName = "ConfigenTpl"
	TpRaw =
		`# config generate by configen
# copyright github.com/landers1037/configen
{{range $r := . -}}
	{{print $r}}
{{end}}`
)

// MakeTemplateRaw 按照字符串数组生成配置文件
// 需要传入[]string
func MakeTemplateRaw(w io.Writer, data []string) error {
	t, err := template.New(TpName).Parse(TpRaw)
	if err != nil {
		return Err(ErrGenTemp, err.Error())
	}

	err = t.Execute(w, data)
	if err != nil {
		return Err(ErrGenTemp, err.Error())
	}

	return nil
}

// MakeTemplate 生成模板
func MakeTemplate(tpl string, w io.Writer, data interface{}) error {
	t, err := template.ParseFiles(tpl)
	if err != nil {
		return Err(ErrGenTemp, err.Error())
	}

	err = t.Execute(w, data)
	if err != nil {
		return Err(ErrGenTemp, err.Error())
	}

	return nil
}

// MakeTemplateString 使用字符串生成模板
func MakeTemplateString(tpl string, w io.Writer, data interface{}) error {
	t, err := template.New(TpName).Parse(tpl)
	if err != nil {
		return Err(ErrGenTemp, err.Error())
	}

	err = t.Execute(w, data)
	if err != nil {
		return Err(ErrGenTemp, err.Error())
	}

	return nil
}