package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	// Use the Header().Add() method to add a 'Server: Go' header to the
	// response header map. The first parameter is the header name, and
	// the second parameter is the header value.
	w.Header().Add("Server", "Go")

	w.Write([]byte("Hello from Snippetbox"))
}

// Add a snippetView handler function.
func snippetView(w http.ResponseWriter, r *http.Request) {
	// Extract the value of the id wildcard from the request using r.PathValue()
	// and try to convert it to an integer using the strconv.Atoi() function. If
	// it can't be converted to an integer, or the value is less than 1, we
	// return a 404 page not found response.
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	// Use the fmt.Sprintf() function to interpolate the id value with a
	// message, then write it as the HTTP response.
	//msg := fmt.Sprintf("Display a specific snippet with ID %d...", id)
	//w.Write([]byte(msg))

	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

// Add a snippetCreate handler function.
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet..."))
}

// Add a snippetCreatePost handler function.
func snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)

	w.Write([]byte("Save a new snippet..."))
}

// register their routes with something called the default
// servemux. This is just a regular servemux like we’ve already been using, but which is
// initialized automatically by Go and stored in the http.DefaultServeMux global variable.
func main() {
	mux := http.NewServeMux()
	//mux.HandleFunc("/{$}", home) // Restrict this route to exact matches on / only.
	//mux.HandleFunc("/snippet/view/{id}", snippetView)
	//mux.HandleFunc("/snippet/create", snippetCreate)
	// Prefix the route patterns with the required HTTP method (for now, we will
	// restrict all three routes to acting on GET requests).
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	// Create the new route, which is restricted to POST requests only.
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	log.Print("starting server on :4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

	// http://localhost:4000/
	// http://localhost:4000/foo/bar (will not work after adding the $ to the trailing slash)
	// http://localhost:4000/snippet/view
	// http://localhost:4000/snippet/create
}
