package main

import (
	"fmt"
	"math/rand"
	"strings"
)

// go run hello.go
func main() {
	fmt.Println("Hello, World!")

	const lS = 232244
	var distance = 234324
	distance = 234
	fmt.Println(distance)

	var (
		var1 = "nihao"
		var2 = "hello"
	)
	fmt.Println(var1, var2)

	var num = rand.Intn(10)
	fmt.Println(num)

	var exist = strings.Contains("hello world", "world")
	fmt.Println(exist)
}
