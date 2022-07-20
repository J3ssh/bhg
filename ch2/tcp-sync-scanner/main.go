package main

import (
	"fmt"
	"sync"
)

func worker(ports chan int, wg *sync.WaitGroup) {
	for p := range ports { //rececive from the port channel, looping till closed
		fmt.Println(p)
		wg.Done()
	}
}

func main() {
	ports := make(chan int, 100) //creating channel, channel limit cap = 100, can only hold 100 itmes
	var wg sync.WaitGroup
	for i := 0; i < cap(ports); i++ { //for loop to create number of workers
		go worker(ports, &wg)
	}
	for i := 1; i <= 1024; i++ { //for loop port generation
		wg.Add(1)
		ports <- i //sending ports to port channel 
	}
	wg.Wait()
	close(ports) //closing channel
}
