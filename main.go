package main

import (
	"log"
	"strconv"
	"fmt"
"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippet Box")) //--> Displays the text on home page
}

func snippetView(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.PathValue("id")) //--> extract wildcard Id value and conv to integer

	if err != nil || id < 1 { //--> if err is not nill and id is not greater than 1 return NotFound
		http.NotFound(w,r)
		return
	}

	msg := fmt.Sprintf("Display a specific snippet with ID %d", id)

	w.Write([]byte (msg))//--> Displays a snippet with a specific ID
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte ("Create a snippet")) //--> Displays Create a snippet
}

func main() {

	mux:= http.NewServeMux() //--> creates a routing system 

	mux.HandleFunc("/{$}", home) //--> maps / path to home handler

	mux.HandleFunc("/snippet/view/{id}", snippetView) //--> maps /snippet/view to snippetView handler, uses {id} wildcard segment

	mux.HandleFunc("/snippet/create", snippetCreate) // --> maps /snippet/create to snippetCreate handler

	log.Print("starting server on :4000") //--> log starting server on port :4000 to the terminal

	err:= http.ListenAndServe(":4000", mux) //--> check for errors

	log.Fatal(err) //--> if there is an error log it to the terminal

}