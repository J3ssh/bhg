package main

import (
	"fmt"
	"sync"
)

func worker(ports chan int, wg *sync.WaitGroup) {
	for p := range ports {
		fmt.Println(p)
		wg.Done()
	}
}

func main() {
	ports := make(chan int, 100) //creating channel, channel limit cap = 100, can only hold 100 itmes
	var wg sync.WaitGroup
	for i := 0; i < cap(ports); i++ { //for loop to create number of worker goroutines
		go worker(ports, &wg)
	}
	for i := 1; i <= 1024; i++ { //for loop port generation
		wg.Add(1)
		ports <- i
	}
	wg.Wait()
	close(ports) //closing channel
}
