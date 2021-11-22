/*
Project: configen parse_ini.go
Created: 2021/11/19 by Landers
*/

package configen

import (
	"gopkg.in/ini.v1"
)

func parseINI(c interface{}, file string) error {
	if !checkStructPtr(c) {
		return Err(ErrNotPointer)
	}
	cfg, err := ini.Load(file)
	if err != nil {
		return Err(ErrIni, err.Error())
	}
	err = cfg.MapTo(c)
	if err != nil {
		return Err(ErrIni, err.Error())
	}

	return nil
}

func saveINI(c interface{}, file string) error {
	if !checkStructPtr(c) {
		return Err(ErrNotPointer)
	}
	cfg := ini.Empty()
	err := ini.ReflectFrom(cfg, c)
	if err != nil {
		return Err(ErrSaveIni, err.Error())
	}

	err = cfg.SaveTo(file)
	if err != nil {
		return Err(ErrSaveIni, err.Error())
	}

	return nil
}
