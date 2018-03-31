package test_test

import (
	"fmt"

	"github.com/cdongyang/library/utils"
)

func ExampleLoadCurrentDirJSONFile() {
	var config = struct {
		Test bool `json:"test"`
	}{}
	err := utils.LoadCurrentDirJSONFile("test.json", &config)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(config.Test)
	//Output:
	//true
}
