package main

import (
	"flag"
	"fmt"
	"net/http"

	f "goupie-tracker/server"
)

// The program entry point
func main() {
	// Serve static files with custom handler to prevent directory listing
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", f.NoDirListing(http.StripPrefix("/static/", fs)))

	// Routing handlers
	http.HandleFunc("/", f.HomePage)
	http.HandleFunc("/artistInfo", f.ArtistPage)

	// Parse the server port from the command line flags
	serverPort := flag.String("port", "8080", "port to serve on")
	flag.Parse()

	fmt.Printf("http://localhost:%s - Server started on port\n", *serverPort)

	// Start the server && listen the requests
	http.ListenAndServe(":"+*serverPort, nil)
}
