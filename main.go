package main

import (
	"fmt"

	sample "github.com/elliotforbes/test-package"
)

func Hello() string {
	return "Bla"
}
func main() {
	fmt.Println("Hello Worldsss")
	sample.MySampleFunc()
}
