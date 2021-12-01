/*
Project: configen config_default_test.go
Created: 2021/12/1 by Landers
*/

package configen

import (
	"testing"
)

// 默认参数配置
func TestDefault(t *testing.T) {
	type T struct {
		Name string
	}

	data := T{Name: "a"}
	Default(&data.Name, "b")
	t.Log(data)
}
