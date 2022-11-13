package main

import "fmt"

func main() {
    var i int
    var f float64
    var b bool
    var s string
    fmt.Printf("%v %v %v %q\n", i, f, b, s)
}


func test0() {
	var intVal int
	intVal :=1 // 这时候会产生编译错误，因为 intVal 已经声明，不需要重新声明
}

// 可以将 var f string = "Runoob" 简写为 f := "Runoob"：
func test1() {
	f := "Runoob" // var f string = "Runoob"
    fmt.Println(f)
}

// 根据值自行判定变量类型。
func test2() {
	var d = true
    fmt.Println(d)
}
