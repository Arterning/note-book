/*
 * @Author: ning.huang ning.huang
 * @Date: 2022-11-14 10:52:05
 * @LastEditors: ning.huang ning.huang
 * @LastEditTime: 2022-11-14 10:58:14
 * @FilePath: \code\goLang\tcp-scanner\single.go
 * @Description: 非并发版本的Tcp扫描器
 */
package main

import (
	"fmt"
	"net"
)

func main() {
	for i := 21; i < 120; i++ {

		address := fmt.Sprintf("108.160.139.217:%d", i)

		conn, err := net.Dial("tcp", address)
		if err != nil {
			fmt.Printf("%s close \n", address)
			continue
		}
		conn.Close()
		fmt.Printf("%s open\n", address)
	}
}
