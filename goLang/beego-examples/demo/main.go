/*
 * @Author: error: git config user.name && git config user.email & please set dead value or install git
 * @Date: 2022-11-14 09:08:27
 * @LastEditors: error: git config user.name && git config user.email & please set dead value or install git
 * @LastEditTime: 2022-11-14 09:18:02
 * @FilePath: \code\goLang\beego-examples\demo\main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
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

// 接口只是定义方法 实现方法需要结构体
type Usber interface {
	start(string, string) string
	stop()
}

type Phone struct {
	Name string
}

// 结构体必须实现接口的所有方法
func (p Phone) start() {
	fmt.Println(p.Name, "start")
}

func (p Phone) stop() {
	fmt.Println("stop")
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

	p := Phone{Name: "黄宁手机"}
	p.start()
	p.stop()
	var o = p
	o.stop()
}
