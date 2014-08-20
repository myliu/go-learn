package main

import "fmt"

// Foo is the interface that needs to be implemented.
type Foo interface {
	Run(s string)
}

// fooFunc is a convenience type to avoid having to declare a
// struct to implement the Foo interface.
type fooFunc func(s string)

func (f fooFunc) Run(s string) {
	f(s)
}

// Static verification
var _ Foo = fooFunc(nil)

func main() {
	f := func(s string) {
		fmt.Println(s)
	}

	foo := fooFunc(f)

	foo.Run("Hello, World")
}
