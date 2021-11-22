/*
Project: configen parse_json.go
Created: 2021/11/19 by Landers
*/

package configen

func parseJson(c interface{}, file string) error {
	if !checkStructPtr(c) {
		return Err(ErrNotPointer)
	}
	data, err := loadFile(file)
	if err != nil {
		return Err(ErrJson, err.Error())
	}

	err = json.Unmarshal(data, c)
	if err != nil {
		return Err(ErrJson, err.Error())
	}

	return nil
}

func saveJson(c interface{}, file string) error {
	if !checkStructPtr(c) && !checkStruct(c) {
		return Err(ErrNotPointer, ErrNotStruct)
	}
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return Err(ErrSaveJson, err.Error())
	}

	err = saveFile(data, file)
	if err != nil {
		return Err(ErrSaveJson, err.Error())
	}

	return nil
}
