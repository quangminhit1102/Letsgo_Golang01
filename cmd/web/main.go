package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	addr := flag.String("addr", ":4000", "HTTP network address") // Get Flag Variable
	flag.Parse()                                                 // Parse the Flag Variable

	fileServer := http.FileServer(http.Dir("./ui/static/")) // Static File Server
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)                  // Info Log
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile) // Error Log

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	infoLog.Printf("Starting server on :%s", *addr) // flag return a Pointer
	err := http.ListenAndServe(":4000", mux)

	errorLog.Fatal(err)
}
