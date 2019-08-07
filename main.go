package gojebug

import (
	"encoding/json"
	"fmt"

	"github.com/stretchr/testify/assert"
)

var (
	should = assert.New(TLogger{})
)

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}

type TLogger struct{}

func (t TLogger) Errorf(format string, args ...interface{}) {
	panic(fmt.Sprintf(format, args...))
}

func PrettyJsonPrint(something interface{}) {
	j, err := json.MarshalIndent(something, "", "\t")
	CheckErr(err)
	fmt.Println(string(j))
}

func JsonPrint(something interface{}) {
	j, err := json.Marshal(something)
	CheckErr(err)
	fmt.Println(string(j))
}

func Equal(expected interface{}, actual interface{}, msgAndArgs ...interface{}) bool {
	return should.Equal(expected, actual, msgAndArgs...)
}
