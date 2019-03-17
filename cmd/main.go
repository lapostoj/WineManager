package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/lapostoj/winemanager/service/presentation"
)

const defaultPort = "8080"

// Init is called before the application starts.
func main() {
	router := presentation.NewRouter()

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./app/")))
	http.Handle("/api/", router)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}
