/*
Project: configen config.go
Created: 2021/11/19 by Landers
*/

package configen

// Default 修改默认值
// field must be a pointer
func Default(filed interface{}, def interface{}) {
	filed = def
}

func DefaultEnv(filed interface{}, key string) {
	
}
