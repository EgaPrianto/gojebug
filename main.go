package gojebug

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/stretchr/testify/assert"
)

var (
	should = assert.New(TLogger{})
)

func CheckErr(err error) {
	should.NoError(err)
}

func print(in string) string {
	fmt.Println(in)
	return in
}

type TLogger struct{}

func (t TLogger) Errorf(format string, args ...interface{}) {
	panic(fmt.Sprintf(format, args...))
}

func prettyJsonPrint(something interface{}) string {
	j, err := json.MarshalIndent(something, "", "  ")
	CheckErr(err)
	return string(j)
}

func PrettyJsonPrint(something interface{}) string {
	return print(prettyJsonPrint(something))
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

func printReaderContent(reader io.Reader) string {
	b, err := ioutil.ReadAll(reader)
	CheckErr(err)
	return string(b)
}

func PrintReaderContent(reader io.Reader) string {
	return print(printReaderContent(reader))
}

func printReaderContentJSON(reader io.Reader) string {
	var res map[string]interface{}
	bodyString := printReaderContent(reader)
	if bodyString != "" {
		err := json.Unmarshal([]byte(bodyString), &res)
		CheckErr(err)
		return prettyJsonPrint(res)
	}
	return ""
}

func PrintReaderContentJSON(reader io.Reader) string {
	return print(printReaderContentJSON(reader))
}

func PrintRequest(r http.Request) string {
	return print(printRequest(r))
}

func printRequest(r http.Request) string {
	var res string
	res += "METHOD = " + r.Method + "\n"
	res += fmt.Sprintf("======%s================================\n", "URL")
	res += r.URL.String() + "\n"
	if len(r.Header) > 0 {
		res += fmt.Sprintf("======%s================================\n", "HEADERS")
		res += prettyJsonPrint(r.Header) + "\n"
	}
	if len(r.URL.Query()) > 0 {
		res += fmt.Sprintf("======%s================================\n", "QUERY PARAMS")
		res += prettyJsonPrint(r.URL.Query()) + "\n"
	}
	res += fmt.Sprintf("======%s================================\n", "BODY")
	res += printReaderContentJSON(r.Body) + "\n"
	return res
}
