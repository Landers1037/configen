/*
Project: configen file_saver.go
Created: 2021/11/22 by Landers
*/

package configen

import (
	"io/ioutil"
	"sync"
)

func saveFile(data []byte, f string) error {
	lock := sync.Mutex{}
	lock.Lock()

	defer lock.Unlock()
	err := ioutil.WriteFile(f, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
