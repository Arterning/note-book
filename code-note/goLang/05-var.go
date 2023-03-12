/*
 * @Author: ning.huang ning.huang
 * @Date: 2022-11-14 09:08:29
 * @LastEditors: ning.huang ning.huang
 * @LastEditTime: 2022-11-14 10:50:21
 * @FilePath: \code\goLang\examples\p5_var.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

func test0() {
	var intVal int
	intVal := 1 // 这时候会产生编译错误，因为 intVal 已经声明，不需要重新声明
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

// 短声明变量的好处
func test3() {

	switch num := rand.Intn(10); num {

	case 0:
		fmt.Println("0000")
	case 1:
		fmt.Println("1111")
	default:
		fmt.Println("default")
	}

}
