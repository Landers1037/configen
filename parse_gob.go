/*
Project: configen parse_gob.go
Created: 2021/11/19 by Landers
*/

package configen

import (
	"bytes"
	"encoding/gob"
	"os"
)

// RegToConfig 注册gob
func RegToConfig(v interface{}) {
	gob.Register(v)
}

func parseGdx(c interface{}, file string) error {
	if !checkStructPtr(c) {
		return Err(ErrNotPointer)
	}
	data, err := loadFile(file)
	if err != nil {
		return Err(ErrGdx, err.Error())
	}

	de := gob.NewDecoder(bytes.NewReader(data))
	err = de.Decode(c)
	if err != nil {
		return Err(ErrGdx, err.Error())
	}

	return nil
}

func saveGdx(c interface{}, file string) error {
	if !checkStructPtr(c) {
		return Err(ErrNotPointer)
	}
	f, err := os.OpenFile(file, os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return Err(ErrSaveGdx, err.Error())
	}
	en := gob.NewEncoder(f)
	err = en.Encode(c)

	if err != nil {
		return Err(ErrSaveGdx, err.Error())
	}

	return nil
}
