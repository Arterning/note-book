package calc

import "fmt"

//private var belong to package
var aa = "abc"

//public variable
var Age = 90

// 首字母大写表示public func
func Add(x, y int) int {
	return x + y
}

func Sub(x, y int) int {
	return x - y
}

//packge init method would be executed..
func init() {
	fmt.Println("calc pacakge init...")
}
