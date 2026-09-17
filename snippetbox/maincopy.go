package main

import (
	"log"
	"net/http"
)

func home1(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}

// Add a snippetView handler function.
func snippetView2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet..."))
}

// Add a snippetCreate handler function.
func snippetCreate2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet..."))
}

func maincopy() {
	http.HandleFunc("/", home1)
	http.HandleFunc("/snippet/view", snippetView2)
	http.HandleFunc("/snippet/create", snippetCreate2)

	log.Print("starting server on :4000")

	err := http.ListenAndServe(":4000", nil)
	log.Fatal(err)

	// http://localhost:4000/
	// http://localhost:4000/foo/bar (will not work after adding the $ to the trailing slash)
	// http://localhost:4000/snippet/view
	// http://localhost:4000/snippet/create
}
