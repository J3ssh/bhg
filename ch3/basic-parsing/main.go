package main

import (
	"encoding/json" //encoding and json handling makes it able to parse json response 
	"log"           //log handling
	"net/http"      //http connection  package
)

type Status struct { //create struct contains exected response
	Message string
	Status  string
}

func main() { //create main function to send post request
	res, err := http.Post( //craft post request
		"https://domain.me", //domain
		"application/json",
		nil,
	)
	if err != nil {
		log.Fatalln(err)
	}

	var status Status
	if err := json.NewDecoder(res.Body).Decode(&status); err != nil { //decode response body status
		log.Fatalln(err) //log error
	}

	defer res.Body.Close()                                  //close connection
	log.Printf("%s -> %s\n", status.Status, status.Message) //lquery  the status struct by accessing exported data types

}
