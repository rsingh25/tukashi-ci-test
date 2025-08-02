package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	port := os.Getenv("PORT")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested:::: %s\n", r.URL.Path)
	})

	fmt.Printf("Server is starting on port %s...\n", port)
	http.ListenAndServe(":"+port, nil)
}
