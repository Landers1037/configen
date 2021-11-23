/*
Project: configen parse_json_test.go
Created: 2021/11/22 by Landers
*/

package configen

import (
	"testing"
)

type JsonData struct {
	Name string
	Age  int
}

func TestParseConfig(t *testing.T) {
	v := new(JsonData)
	err := ParseConfig(v, JSON, "test.json")
	t.Log(v)
	t.Log(err)
}

func TestParseJson(t *testing.T) {
	v := new(JsonData)
	err := parseJson(v, "test.json")
	if err != nil {
		t.Error(err.Error())
		return
	}

	t.Logf("%+v", v)
}

func TestParseINI(t *testing.T) {
	v := new(JsonData)
	err := parseINI(v, "test.ini")
	if err != nil {
		t.Error(err.Error())
		return
	}

	t.Logf("%+v", v)
}
