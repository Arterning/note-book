package main

import (
	"demo/calc"
	"fmt"
)

// go run main.go
func main() {
	sum := calc.Add(10, 2)
	fmt.Println(sum)
}
