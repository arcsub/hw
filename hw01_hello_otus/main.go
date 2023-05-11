package main

import (
	"fmt"

	util "golang.org/x/example/stringutil"
)

func main() {
	text := "Hello, OTUS!"
	fmt.Println(util.Reverse(text))
}
