package main

import (
	"fmt"
	"github.com/EgaPrianto/gojebug"
	"net/http"
	"strings"
)

type Compare struct {
	Field1 int
	Field2 string
}

func main() {
	object1 := Compare{
		Field1: 23,
		Field2: "some words",
	}
	object2 := Compare{
		Field1: 23,
		Field2: "some words",
	}
	//object3 := Compare{
	//	Field1: 33,
	//	Field2: "some words",
	//}
	gojebug.JsonPrint(object1)
	gojebug.PrettyJsonPrint(object1)
	fmt.Println("Equal : ", gojebug.Equal(object1, object2))
	//fmt.Println("Equal : ", gojebug.Equal(object2, object3))

	r, err := http.NewRequest(http.MethodGet, "http://localhost:806/path/to/do?somequery=123&andanotherone=123", strings.NewReader("this supposed to be body"))
	gojebug.CheckErr(err)

	gojebug.PrintRequest(*r)
}
