package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {
	target := "127.0.0.1"
	var wg sync.WaitGroup
	fmt.Printf("Scanning %s\n", target)
	for i := 1; i <= 65535; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			address := fmt.Sprintf("%s:%d", target, j)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				// port closed or filtered
				return
			}
			conn.Close()
			fmt.Printf("%d open\n", j)
		}(i)
	}
	wg.Wait()
	fmt.Printf("\nScan finnished\n")
}
