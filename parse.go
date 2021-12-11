/*
Project: configen parse.go
Created: 2021/11/20 by Landers
*/

package configen

import (
	"github.com/json-iterator/go"
)

// 加载配置文件 映射到结构体中

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// ParseConfig 指定配置文件类型解析
// c must be a pointer
// t only support messages INI, Json, Yaml, Gdx
func ParseConfig(c interface{}, t, file string) error {
	for _, v := range support {
		if t == v {
			return parseForType(c, t, file)
		}
	}

	return Err(ErrSuffix)
}

// ParseConfigMap 指定配置文件类型解析为字典
// t only support messages INI, Json, Yaml, Gdx
func ParseConfigMap(t, file string) (*map[string]interface{}, error) {
	for _, v := range support {
		if t == v {
			var m map[string]interface{}
			err := parseForType(&m, t, file)
			if err != nil {
				return nil, err
			}

			return &m, nil
		}
		break
	}

	return nil, Err(ErrSuffix)
}

// SaveConfig 保存数据到文件中
// c must be a pointer
// t is type must be in supports
// file is the output file
func SaveConfig(c interface{}, t, file string) error {
	for _, v := range support {
		if t == v {
			return saveForType(c, t, file)
		}
	}

	return Err(ErrSuffix)
}

// c should be a pointer
func parseForType(c interface{}, t, file string) error {
	switch t {
	case INI:
		return parseINI(c, file)
	case JSON:
		return parseJson(c, file)
	case Pig:
		return parseJson(c, file)
	case Yaml:
		return parseYaml(c, file)
	case Gdx:
		return parseGdx(c, file)
	default:
		return Err(ErrSuffix)
	}
}

// 基于枚举的条件选择
func parseByEnum() error {
	return nil
}

func saveForType(c interface{}, t, file string) error {
	switch t {
	case INI:
		return saveINI(c, file)
	case JSON:
		return saveJson(c, file)
	case Pig:
		return saveJson(c, file)
	case Yaml:
		return saveYaml(c, file)
	case Gdx:
		return saveGdx(c, file)
	default:
		return Err(ErrSuffix)
	}
}
