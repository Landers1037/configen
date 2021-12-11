/*
Project: configen template_test.go
Created: 2021/11/23 by Landers
*/

package configen

import (
	"os"
	"testing"
)

func TestMakeTemplateNginx(t *testing.T) {
	err := MakeTemplateNginxTest("nginx.tpl",
		map[string]interface{}{
			"app":    "Test",
			"listen": "80",
			"domain": "a.b.c",
			"des":    "",
			"port":   "1000",
		})
	if err != nil {
		t.Error(err)
	}
}

func TestMakeTemplateGunicorn(t *testing.T) {
	err := MakeTemplateGunicornTest("gunicorn.tpl", map[string]interface{}{"port": "1000"})
	if err != nil {
		t.Error(err)
	}
}

func TestMakeTemplateRaw(t *testing.T) {
	err := MakeTemplateRaw(os.Stdout, []string{"a", "b", "c"})
	if err != nil {
		t.Error(err)
	}
}
