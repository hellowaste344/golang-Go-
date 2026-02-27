package main

import (
	"fmt"
	"log"
	"net/http"
)

func httpServer() {
	// create an endpoint
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Greetings, %s!", r.URL.Path[1:])
	})
	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "What's up, %s!", r.URL.Path[1:])
	})
	http.HandleFunc("/bye", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Goodbye, %s!", r.URL.Path[1:])
	})

	port := ":8080"
	fmt.Println("Server is running on port", port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
