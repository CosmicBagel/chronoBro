package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	println("hello world")
	println("lookie")

	/*
		- http serve
		- hello world html?
			- serve page html? (templates?)
			- serve static js/css/html?
		- routes for get and post
		- send post to db
		- get from db
	*/

	// mux := http.NewServeMux()
	// mux.Handle("", )

	s := http.Server{
		Addr:         ":8080",
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 90 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      http.HandlerFunc(hello),
	}

	err := s.ListenAndServe()
	if err != nil {
		if err != http.ErrServerClosed {
			panic(err)
		}
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
    _, err := fmt.Fprintf(w, "hello world %s\n", r.Method)
    if err != nil {
        panic(err)
    }
}
// domain/service/objectReceivingAction/action/version
