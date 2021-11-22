/*
Project: configen save_test.go
Created: 2021/11/23 by Landers
*/

package configen

import (
	"testing"
)

func TestSaveJson(t *testing.T) {
	v := JsonData{
		Name: "test",
		Age:  2,
	}

	err := saveJson(v, "test.json")
	if err != nil {
		t.Error(err.Error())
	}
}

func TestSaveINI(t *testing.T) {
	v := JsonData{
		Name: "test",
		Age:  2,
	}

	err := saveINI(&v, "test.ini")
	if err != nil {
		t.Error(err.Error())
	}
}

func TestSaveGob(t *testing.T) {
	v := JsonData{
		Name: "test",
		Age:  2,
	}

	err := saveGdx(v, "test.gdx")
	if err != nil {
		t.Error(err.Error())
	}
}
