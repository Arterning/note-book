/*
不要通过共享内存来通信，而应该通过通信来共享内存
这里的共享内存就是ports这个带有缓存的管道。。也就是说go是通过管道通信来共享ports内存。。ports内存必须是要保证是线程安全的
因为同时会有多个线程读取ports
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
