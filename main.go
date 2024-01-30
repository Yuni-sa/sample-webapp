package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// Get the port from the environment variable or use a default value (e.g., 3000)
	port := getPort()

	// Serve the index.html file
	http.Handle("/", http.FileServer(http.Dir("static")))

	// Endpoint to serve the port information
	http.HandleFunc("/port", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s", port)
	})

	http.HandleFunc("/env", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s", getEnv())
	})

	// Start the server
	fmt.Printf("Server is running on :%s\n", port)
	http.ListenAndServe(":"+port, nil)
}

func getPort() string {
	// Get the port from the environment variable, or use a default value (e.g., 3000)
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	return port
}

func getEnv() string {
	// Get the port from the environment variable, or use a default value (e.g., 3000)
	env := os.Getenv("ENV")
	if env == "" {
		env = "dev"
	}
	return env
}
