package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// post (create), get (read), put (update), delete
// https://chrono-api.cosmicbagel.com/v1/events/{id}
// domain/service/objectReceivingAction/action/version

// /events/{id} [get, put, delete]
// /events [post]

// db events table
// id (int64), start (datetime), end (datetime), title (varchar), description (varchar)

func main() {

	println("hello world")
	println("lookie")

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/events/", eventsPathHandler)
	mux.HandleFunc("/", rootPathHandler)

	s := http.Server{
		Addr:         ":8080",
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 90 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      mux,
	}

	err := s.ListenAndServe()
	if err != nil {
		if err != http.ErrServerClosed {
			panic(err)
		}
	}
}

func eventsPathHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/v1/events/")

	if len(idStr) == 0 {
		
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintf(w, "400 bad request")
		return
	}

	_, err = fmt.Fprintf(w, "events %d\n", id)
	if err != nil {
		panic(err)
	}

}

func rootPathHandler(w http.ResponseWriter, r *http.Request) {
	// matches everything that doesn't match a handled path
	// so if path is not literally root "/", then it is 404
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	_, err := fmt.Fprintf(w, "hello world %s\n", r.Method)
	if err != nil {
		panic(err)
	}
}
