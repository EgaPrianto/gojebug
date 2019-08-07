# GOJEBUG

## Description
Gojebug is a tool for golang debugging to stdout using standard JSON format. 

## Example

```go
    package main
    
    import "github.com/EgaPrianto/gojebug"
    
    type Compare struct {
    	Field1 int
    	Field2 string
    }
    
    func main(){
    	object := Compare{
    		Field1:23,
    		Field2: "some words",
    	}
    	object2 := Compare{
    		Field1:23,
    		Field2: "some words",
    	}
    	gojebug.JsonPrint(object)
    	gojebug.PrettyJsonPrint(object)
    	gojebug.Equal(object,object2)
    }
 ```