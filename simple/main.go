package main

import (
	"fmt"
	"net/http"
	"os"
)

func server(w http.ResponseWriter, r *http.Request) {

	data, err := os.ReadFile("index.html")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Set the content type header (important for the browser)
	w.Header().Set("Content-Type", http.DetectContentType(data))

	// Write the file contents to the response using fmt.Fprintf
	_, err = fmt.Fprintf(w, "%s", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	// Create a handler function
	http.HandleFunc("/", server)

	// Start the server
	fmt.Println("Server listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
