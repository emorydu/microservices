package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	// requests to the path /goodbye with be handled by this function
	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("Goodbye World")
	})

	// any other request will be handled by this function
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Running Hello World")

		// read the body
		b, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println("Error reading body", err)

			http.Error(w, "Unable to read request body", http.StatusBadRequest)
			return
		}

		// write the response
		fmt.Fprintf(w, "Hello, %s\n", b)
	})

	// Listen for connections on all ip address (0.0.0.0)
	// port 9000
	log.Println("Starting Server")
	err := http.ListenAndServe(":9000", nil)
	log.Fatal(err)
}
