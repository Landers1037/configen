/*
Project: configen parse_tags_test.go
Created: 2021/11/23 by Landers
*/

package configen

import (
	"os"
	"testing"
)

type TagTest struct {
	Name string `config:"env"`
	Age  int
}

func TestParseTags(t *testing.T) {
	_ = os.Setenv("env", "Conf")
	v := new(TagTest)
	err := ValidateTags(v, "env", func() interface{} {
		return os.Getenv("env")
	})

	t.Error(err)
	t.Logf("%+v", v)
}
