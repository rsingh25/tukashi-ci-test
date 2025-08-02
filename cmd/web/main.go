package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	port := os.Getenv("PORT")
	host := os.Getenv("HOSTNAME")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, this request is served by host %s\n", host)
	})

	fmt.Printf("Server is starting on port %s...\n", port)
	http.ListenAndServe(":"+port, nil)
}
