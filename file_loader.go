/*
Project: configen file_loader.go
Created: 2021/11/19 by Landers
*/

package configen

import (
	"io/ioutil"
	"sync"
)

// 统一的文件加载器
func loadFile(f string) ([]byte, error) {
	lock := sync.Mutex{}
	lock.Lock()
	data, err := ioutil.ReadFile(f)
	defer lock.Unlock()

	if err != nil {
		return nil, Err(ErrLoad, err.Error())
	}

	return data, nil
}
