/*
Create a struct “customErr” which implements the builtin.error interface. Create a func “foo” that
has a value of type error as a parameter. Create a value of type “customErr” and pass it into
“foo”. If you need a hint, here is one. https://go.dev/play/p/L5QWV8-p11
*/

package main

import "fmt"

type customErr struct {
	info string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("error info: %v", ce.info)
}

func foo(e error) {
	fmt.Println("foo:", e)

}

func main() {
	e := customErr{info: "something went wrong"}

	foo(e)

}
