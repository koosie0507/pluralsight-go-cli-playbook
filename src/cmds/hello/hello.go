package main

import (
	"fmt"
	"hello"
)
import "C"

func main() {
	hello.SayHello()
}

//export Hello
func Hello() {
	fmt.Println("Hello")
}
