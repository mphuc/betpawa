package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Success!") // 200
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"

	}

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "OK") // 200
	})
	log.Fatal(http.ListenAndServe("0.0.0.0:"+port, nil))
}
