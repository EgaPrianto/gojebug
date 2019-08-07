package gojebug

import (
	"encoding/json"
	"fmt"
)

func PrettyJsonPrint(something interface{}) {
	j, _ := json.MarshalIndent(something, "", "\t")
	fmt.Println(string(j))
}
