/*
*
项目里面只能由一个main方法和main包
*
*/
package main

//T 包的别名
import (
	"demo/calc"
	_ "demo/test"
	T "demo/tools"
	"fmt"

	"github.com/shopspring/decimal"
)

/*
*
main init 是最后执行的
如果main include B , B include C
那么C init -> B init -> main init
*
*/
func init() {
	fmt.Println("main init...")
}

// go run main.go
func main() {
	sum := calc.Add(10, 2)
	fmt.Println(sum)

	fmt.Println(calc.Age)

	fmt.Println(T.Mul(8, 9))

	T.Fuck()

	var quantity = decimal.NewFromInt(3)
	fmt.Println(quantity)
}
