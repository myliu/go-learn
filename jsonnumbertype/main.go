// An example shows the default number type after JSON decoding
//
// Following JSON decoding, if the concrete type of the interface is number,
// the default type is float64.  In order to be decoded to int, the field type
// has to be explicitly set to int.
//
// Source
//  http://golang.org/pkg/encoding/json/#Unmarshal
package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Animal struct {
	Id    interface{}
	IntId int
}

func main() {
	a := Animal{
		Id:    1,
		IntId: 2,
	}

	fmt.Println(reflect.ValueOf(a.Id).Kind())
	fmt.Println(reflect.ValueOf(a.IntId).Kind())

	// Output:
	// int
	// int

	p, err := json.Marshal(a)

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(p, &a)

	if err != nil {
		panic(err)
	}

	fmt.Println(reflect.ValueOf(a.Id).Kind())
	fmt.Println(reflect.ValueOf(a.IntId).Kind())

	// Output:
	// float64
	// int
}
