package main

import (
	"fmt"

	util "golang.org/x/example/stringutil"
)

func main() {
	message := "Hello, OTUS!"
	fmt.Print(util.Reverse(message))
}
