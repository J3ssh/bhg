package main

import (
	"fmt"
	"net"
)

func main() {
	for i := 1; i <= 1024; i++ { //for loop for port genreation
		address := fmt.Sprintf("scanme.nmap.org:%d", i) //define address
		conn, err := net.Dial("tcp", address)           //define type
		if err != nil {                                 //check error, should probbaly log error here? need to import logging
			//port is closed or filtered
			continue
		}
		conn.Close() //close connection
		fmt.Printf("%d open\n", i)
	}
}
