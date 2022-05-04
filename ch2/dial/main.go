package main

import (
	"fmt"
	"net"
)

func main() {
	_, err := net.Dial("tcp", "scanme.nmap.org:80") // Dial(network, address string)
	if err == nil {                                 //check for error if error nil connection works
		fmt.Println("Connection sucessful") //print working
	}
}
