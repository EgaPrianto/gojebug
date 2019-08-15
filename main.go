package gojebug

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"

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

func PrettyJsonPrint(something interface{}) string {
	j, err := json.MarshalIndent(something, "", "\t")
	CheckErr(err)
	res := string(j)
	fmt.Println(res)
	return res
}

func JsonPrint(something interface{}) string {
	j, err := json.Marshal(something)
	CheckErr(err)
	res := string(j)
	fmt.Println(res)
	return res
}

func Equal(expected interface{}, actual interface{}, msgAndArgs ...interface{}) bool {
	return should.Equal(expected, actual, msgAndArgs...)
}

func PrintReaderContent(reader io.Reader) string{
	b, err := ioutil.ReadAll(reader)
	CheckErr(err)
	res := string(b)
	fmt.Println(res)
	return res
}
