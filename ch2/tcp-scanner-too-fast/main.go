package main

import (
	"fmt"
	"net"
)

func main() { //start main function
	for i := 1; i <= 1024; i++ { //for loop port genreation
		go func(j int) { //goroutine without wait for bad example
			address := fmt.Sprintf("scanme.nmap.org%d", j)
			conn, err := net.Dial("tcp", address)
			if err != nil { //error checking
				return
			}
			conn.Close()               //connection close
			fmt.Printf("%d open\n", j) //response
		}(i)
	}
} //this code launches a single goroutine per connection, the main goroutine doesnt know to wait for the connection to take place
