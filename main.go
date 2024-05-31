package main

import (
	"fmt"
	"net/http"
)

func server(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, `<!DOCTYPE html>
	<html>
	<head>
		 <title>Simple Go Web Page</title>
	</head>
	<body>
		 <h1>Hello from Go!</h1>
		 <p>This is a basic web page served by a Go server.</p>
	</body>
	</html>`)
}

func main() {
	// Create a handler function
	http.HandleFunc("/", server)

	// Start the server
	fmt.Println("Server listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
