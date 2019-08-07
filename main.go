package gojebug

import (
	"encoding/json"
	"fmt"
)

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

func PrettyJsonPrint(something interface{}) {
	j, err := json.MarshalIndent(something, "", "\t")
	CheckErr(err)
	fmt.Println(string(j))
}
