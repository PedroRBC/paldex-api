// This is a development test server for the Paldex API
package main

import (
	"fmt"
	"log"
	"net/http"

	api "github.com/PedroRBC/paldex-api/api"
	_ "github.com/PedroRBC/paldex-api/api/_pkg" // Import for init function
)

func main() {
	// Set up routes using the handler package
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		api.Handler(w, r)
	})

	// Start server
	port := ":8080"
	fmt.Printf("Server starting on port %s...\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
