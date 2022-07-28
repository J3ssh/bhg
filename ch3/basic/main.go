package main

import (
	"fmt"       //import for format
	"io/ioutil" //import for input output control
	"log"       //import for logging
	"net/http"  //import for http connections
	"net/url"   //import for handling url and url parsing etc
	"strings"   //import for handling strings
)

func main() { //create main function

	resp, err := http.Get("https://google.com/robots.txt") //create get request
	if err != nil {                                        //error check
		log.Panicln(err) //log error
	}
	//print HTTP status
	fmt.Println(resp.Status) //output response

	//read and display response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil { //check error
		log.Panicln(err) //log error
	}

	fmt.Println(string(body)) //print response body and strings
	resp.Body.Close()         //close connection

	resp, err = http.Head("https://google.com/robots.txt") //create head request
	if err != nil {                                        //error check
		log.Panicln(err) //error handle
	}
	resp.Body.Close()        //close connection
	fmt.Println(resp.Status) //print status

	form := url.Values{}    //create form
	form.Add("food", "bar") // populate form
	resp, err = http.Post(  //create post request followed by data
		"https://google.com/robots.txt",
		"application/x-www-form-urlencoded",
		strings.NewReader(form.Encode()), //strings and encoding
	)
	if err != nil { //error check
		log.Panicln(err) //log use Paicln for http
	}
	resp.Body.Close()        //close connection
	fmt.Println(resp.Status) //print status

	req, err := http.NewRequest("DELETE", "https://google.com/robots.txt", nil) //create request
	if err != nil {                                                             //check error
		log.Panicln(err) //http log using Panicln
	}
	var client http.Client     //create httpClient
	resp, err = client.Do(req) //create response
	if err != nil {            //error check
		log.Panicln(err) //log error
	}
	resp.Body.Close()        //close connection
	fmt.Println(resp.Status) /// print response

	req, err = http.NewRequest("PUT", "https://google.com/robots.txt", strings.NewReader(form.Encode())) //create put request

	if err != nil { //check error
		log.Panicln(err) //log error
	}
	resp.Body.Close()        //close connection
	fmt.Println(resp.Status) //print status
}
