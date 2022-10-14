package main

import (
	"fmt"

	sample "github.com/elliotforbes/test-package"
)

func Hello() string {
	return "Bla"
}
func main() {
	fmt.Println("Hello Worlds")
	sample.MySampleFunc()
}
