/*
Project: configen parse_tags.go
Created: 2021/11/22 by Landers
*/

package configen

import (
	"reflect"
)

// ValidateTags 自定义标签的解析 用于配置默认值
// eg: `config:"env"` 默认为0值时从环境变量key: env读取
// env 需要读取的环境变量
// f 处理环境变量的函数，返回最终处理后需要更改的值
func ValidateTags(v interface{}, env string, f func() interface{}) error {
	val := reflect.ValueOf(v).Elem()
	typ := reflect.TypeOf(v).Elem()

	for i := 0; i < typ.NumField(); i++ {
		if val.Field(i).IsZero() && typ.Field(i).Tag.Get("config") != "" {
			if typ.Field(i).Tag.Get("config") != env {
				return Err(ErrValidate, "env key invalid")
			}
			envInterface := f()
			if reflect.TypeOf(envInterface).Kind() != val.Field(i).Kind() {
				return Err(ErrValidate)
			}
			val.Field(i).Set(reflect.ValueOf(envInterface))
		}
	}

	return nil
}
