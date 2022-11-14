package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func main() {

	start := time.Now()

	var wg sync.WaitGroup
	for i := 21; i < 6005; i++ {

		wg.Add(1)

		go func(port int) {

			defer wg.Done() //condownlatch.countdown()

			address := fmt.Sprintf("108.160.139.217:%d", port)

			conn, err := net.Dial("tcp", address)
			if err != nil {
				//fmt.Printf("%s close \n", address)
				return
			}
			conn.Close()
			fmt.Printf("%s open\n", address)
		}(i)

	}
	wg.Wait()

	elapsed := time.Since(start) / 1e9
	fmt.Printf("\n%d seconds", elapsed)
}
