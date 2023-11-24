package main

import (
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
	w.Write([]byte("fsdafsadfasfd\n"))
}
