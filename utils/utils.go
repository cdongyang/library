package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
	"runtime"
)

func Lowbit(x int) int {
	return x & (-x)
}

func LoadJSONFile(filepath string, data interface{}) (err error) {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, data)
}

func LoadCurrentDirJSONFile(filename string, data interface{}) (err error) {
	_, file, _, _ := runtime.Caller(1)
	return LoadJSONFile(path.Join(path.Dir(file), filename), data)
}
