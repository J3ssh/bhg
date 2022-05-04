package main

import (
	"fmt"
	"net"
	"sort" //sort function to filter output in order
)

func worker(ports, results chan int) { //worker function accepts two channels
	for p := range ports {
		address := fmt.Sprintf("scanme.nmap.org:%d", p) //setting target connection
		conn, err := net.Dial("tcp", address)           //connection
		if err != nil {                                 //error checking
			results <- 0 //if port closed send 0
			continue
		}
		conn.Close() //connection close
		results <- p //if open send port
	}
}
func main() {
	ports := make(chan int, 100) //create channel limit to 100 items
	results := make(chan int)    //create new channel to communicate results from worker to main thread
	var openports []int          //slice

	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}
	go func() { //sending worker to seperate go routines
		for i := 1; i <= 1024; i++ {
			ports <- i
		}
	}()

	for i := 0; i < 1024; i++ { //result gathering loop
		port := <-results
		if port != 0 { //check port value
			openports = append(openports, port) //if port is bigger than zero add to slice
		}
	}
	close(ports)
	close(results)
	sort.Ints(openports) //sort ports and then loop over slice and print open ports
	for _, port := range openports {
		fmt.Printf("%d open\n", port)
	}
}
