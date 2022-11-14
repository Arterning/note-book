/*
 * @Author: ning.huang ning.huang
 * @Date: 2022-11-14 11:08:48
 * @LastEditors: ning.huang ning.huang
 * @LastEditTime: 2022-11-14 11:29:46
 * @FilePath: \code\goLang\tcp-scanner\multi2.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"fmt"
	"net"
	"sort"
)

var workerNumber = 50

func worker(ports chan int, results chan int) {
	for p := range ports {
		//fmt.Println(p)
		address := fmt.Sprintf("localhost:%d", p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			//fmt.Printf("%s close \n", address)
			results <- 0
			continue
		}
		conn.Close()
		results <- p
		fmt.Printf("%s open\n", address)
	}
}

func main() {
	//带缓存的channel
	ports := make(chan int, 100)

	//不带缓存的channel
	results := make(chan int)

	var openport []int
	var closeport []int

	for i := 0; i < workerNumber; i++ {
		go worker(ports, results)
	}

	//send port to worker
	go func() {
		for i := 0; i < 1024; i++ {
			ports <- i
		}
	}()

	//get result
	for i := 0; i < 1024; i++ {
		port := <-results
		if port != 0 {
			openport = append(openport, port)
		} else {
			closeport = append(closeport, port)
		}
	}

	sort.Ints(openport)
	sort.Ints(closeport)

	for _, op := range openport {
		fmt.Printf("%d open port", op)
	}

	close(ports)
	close(results)

}
