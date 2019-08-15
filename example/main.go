package main

import (
	"fmt"
	"github.com/EgaPrianto/gojebug"
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
	object3 := Compare{
		Field1: 33,
		Field2: "some words",
	}
	gojebug.JsonPrint(object1)
	gojebug.PrettyJsonPrint(object1)
	fmt.Println("Equal : ", gojebug.Equal(object1, object2))
	fmt.Println("Equal : ", gojebug.Equal(object2, object3))
}
