package main

import (
	"fmt"
	"net"
	"sync" //import for WaiteGroup, thread-safe way to control concurrency
)

func main() {
	var wg sync.WaitGroup        //creating WaitGroup
	for i := 1; i <= 1024; i++ { //for loop port genereation
		wg.Add(1)        //increase internal counter by 1
		go func(j int) { //create goroutine
			defer wg.Done() //decrements the counter by one
			address := fmt.Sprintf("scanme.nmap.org:%d", j)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				return
			}
			conn.Close()
			fmt.Printf("%d open\n", j)
		}(i)
	}
	wg.Wait() //blocks exeuction of the goroutine in which it's called, will not allow further execution untill counter reaches zero
}
