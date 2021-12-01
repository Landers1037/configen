/*
Project: configen config.go
Created: 2021/11/19 by Landers
*/

package configen

import (
	"reflect"
)

// Default 修改默认值
// field must be a pointer
func Default(filed interface{}, def interface{}) {
	if reflect.TypeOf(filed).Kind() != reflect.Ptr {
		panic("field must be a pointer")
	}
	if reflect.ValueOf(filed).Elem().CanSet() {
		reflect.ValueOf(filed).Elem().Set(reflect.ValueOf(def))
	}
}


