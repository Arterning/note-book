package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	for _, arg := range os.Args[1:] {
		fmt.Println(arg)
	}

	fmt.Println(strings.Join(os.Args[1:], " "))

}
