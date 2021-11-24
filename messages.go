/*
Project: configen messages.go
Created: 2021/11/19 by Landers
*/

package configen

import (
	"errors"
	"fmt"
	"strings"
)

// errors
const (
	ErrIni      = "failed to parse ini"
	ErrJson     = "failed to parse json"
	ErrYaml     = "failed to parse yaml"
	ErrGdx      = "failed to parse gdx"
	ErrSaveIni  = "failed to save ini"
	ErrSaveJson = "failed to save json"
	ErrSaveYaml = "failed to save yaml"
	ErrSaveGdx  = "failed to save gdx"
	ErrSuffix   = "file type not support"
	ErrLoad     = "failed to load file"
	ErrSave     = "failed to save file"
	ErrGenTemp  = "failed to generate template"

	ErrNotPointer = "input is not a pointer"
	ErrNotStruct  = "input is not a struct"
	ErrValidate   = "failed to reflect type"
)

// infos
const (
	INI    = "ini"
	JSON   = "json"
	Config = "config" // config is an alias of json
	Yaml   = "yaml"
	Gdx    = "gdx"
)

// 类型枚举
const (
	INIType = iota
	JSONType
	YamlType
	GdxType
)

var support = []string{INI, JSON, Config, Yaml, Gdx}

func Err(e string, args ...string) error {
	errs := append(append([]string{}, e), args...)

	return errors.New(fmt.Sprint(
		strings.Join(errs, ",")))
}
