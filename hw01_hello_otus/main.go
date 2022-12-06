package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	message := "Hello, OTUS!"
	fmt.Print(stringutil.Reverse(message))
}
