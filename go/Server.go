package main

import (
	//"html"
	"io"
	"log"
	//"time"
	"net/http"
	"fmt"
)

var Address string = "127.0.0.1"
var Port uint = 8080

// handles all requests
func handle(w http.ResponseWriter, r *http.Request) {
	// parse request
	// construct response

	// figure out the Controller/Action

	// execute Controller->Action and get result

	// write response
	io.WriteString(w, "hello world")
}

func Run() {
	address := fmt.Sprintf("%s:%d", Address, Port)
	s := &http.Server{
		Addr:		    address,
		Handler:		http.HandlerFunc(handle),
		//ReadTimeout:	10 * time.Second,
		//WriteTimeout:   10 * time.Second,
		//MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}

func main() {
	Run()
}

