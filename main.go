package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	http.HandleFunc("/name", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "MOULI!")
	})

	http.HandleFunc("/regnum", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "7376212IT184!")
	})

	http.HandleFunc("/department", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "INFORMATION TECHNOLOGY!")
	})

	http.HandleFunc("/favouritecolor", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "VIOLET!")
	})

	fmt.Printf("Server running (port=8080), route: http://localhost:8080/helloworld\n")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}
