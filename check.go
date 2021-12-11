/*
Project: configen check.go
Created: 2021/11/23 by Landers
*/

package configen

import (
	"reflect"
)

func checkStructPtr(s interface{}) bool {
	if reflect.ValueOf(s).Type().Kind() == reflect.Ptr {
		return true
	}

	return false
}

func checkStruct(s interface{}) bool {
	if reflect.ValueOf(s).Type().Kind() == reflect.Struct {
		return true
	}

	return false
}
