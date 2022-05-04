package main

import (
	"io"
	"log"
	"net"
)

//echo is a hanlder function that simply echoes received data
func echo(conn net.Conn) { //connection hanlder function, the function loops indefineitely
	defer conn.Close()

	//Create a buffer to store received data
	b := make([]byte, 512)
	for {
		//receive data via conn.Read into a buffer
		size, err := conn.Read(b[0:]) 
		if err == io.EOF {
			log.Println("Client disconnected")
			break
		}
		if err != nil {
			log.Println("unexepcted error")
			break
		}
		log.Printf("Received %d bytes: %s\n", size, string(b))

		//Send data via conn.Write
		log.Println("writing data")
		if _, err := conn.Write(b[0:size]); err != nil { //using conn.Write to write data into a buffer?
			log.Fatalln("Unable to write data")
		}
	}
}

func main() {
	//bind to TCP port 20080 on all interfaces
	listner, err := net.Listen("tcp", ":20080") //start TCP listner on defined port 
	if err != nil {
		log.Fatalln("unable to bind on port")
	}
	log.Println("Listening on 0.0.0.0:20080")
	for { //infinite loop ensures that the server will contiune to listen for connections even after one has been received
		//Wait for connection create net.Conn on connection established
		conn, err := listner.Accept() // we call listener.Accept
		log.Println("Received connection")
		if err != nil {
			log.Fatalln("Unable to accppt connection")
		}
		//handle the connection using goroutines for concurrency
		go echo(conn) //prefaced with go makes this concurrent
	}
}
