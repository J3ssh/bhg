package main

import (
	"fmt" //for format
	"log" //logging package log.Fatalln("text %d data")
	"os"  //for read write  "stdin/stdout"
)

//FooReader defines an io.Reader to read from stdin
type FooReader struct{} //define custom type
//read reads data from stdin.
func (fooReader *FooReader) Read(b []byte) (int, error) { //read byte Read([]byte)
	fmt.Print("in>")
	return os.Stdin.Read(b) //read from Stdin
}

//
type FooWriter struct{} //define custom type
//
func (fooWriter *FooWriter) Write(b []byte) (int, error) { //write byte Write([]byte)
	fmt.Print("out>")
	return os.Stdout.Write(b) //write to Stdout
}

//this entire function can be replaced with
// func Copy(dst.io.Writer, src io.Reader) (written int64, error)
func main() { //create input slice
	//Instatiate reader and writer
	var (
		reader FooReader
		writer FooWriter
	)

	//Create buffer to hol d input/output
	input := make([]byte, 4096)

	//Use reader to read input
	s, err := reader.Read(input)
	if err != nil {
		log.Fatalln("Unable to read data")
	}
	fmt.Printf("read %d bytes from stdin\n", s)

	//use writer to write output
	s, err = writer.Write(input)
	if err != nil {
		log.Fatalln("unable to write data")
	}
	fmt.Printf("Wrote %d bytes to stdout\n", s)
}
