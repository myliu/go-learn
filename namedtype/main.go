// An example shows the difference between Named Type and Unnamed Type.
//
// Type
//  Named Type is different from Unnamed Type.
//
// Assignability
//  Variable of Named Type is assignable to variable of Unnamed Type, vice versa.
//  Variable of different Named Type is not assignable to each other.
//
// Source
//  http://golang.org/ref/spec#Type_identity
package main

import (
	"fmt"
	"reflect"
)

type T1 []string
type T2 []string

func main() {
	foo0 := []string{}
	foo1 := T1{}
	foo2 := T2{}
	fmt.Println(reflect.TypeOf(foo0))
	fmt.Println(reflect.TypeOf(foo1))
	fmt.Println(reflect.TypeOf(foo2))

	// Output:
	// []string
	// main.T1
	// main.T2

	// foo0 can be assigned to foo1, vice versa
	foo1 = foo0
	foo0 = foo1

	// foo2 cannot be assigned to foo1
	// prog.go:28: cannot use foo2 (type T2) as type T1 in assignment
	// foo1 = foo2
}
