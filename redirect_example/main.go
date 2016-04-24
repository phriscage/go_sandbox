package main

import (
	"flag"
	"github.com/gorilla/handlers"
	"log"
	"net/http"
	"os"
)

func redirect(w http.ResponseWriter, r *http.Request) {

	http.Redirect(w, r, "http://example.com", 301)
}

func main() {
	//verbose := flag.Bool("v", false, "verbose logging to stdout")
	addr := flag.String("addr", ":8080", "http listen address")
	flag.Parse()
	loggedRouter := handlers.LoggingHandler(os.Stdout, http.DefaultServeMux)
	http.HandleFunc("/", redirect)
	log.Fatal(http.ListenAndServe(*addr, loggedRouter))
}
