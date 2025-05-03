package main

import (
	"log"
"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippet Box")) //--> Displays the text on home page
}

func snippetView(w htttp.ResponseWriter, r *http.Request) {
	w.Write([]byte ("Display a specific snippet")) //--> Displays a specific snippet
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte ("Create a snippet")) //--> Displays Create a snippet
}

func main() {

	mux:= http.NewServeMux() //--> creates a routing system 

	mux.HandleFunc("/", home) //--> maps a particular path to a specific route

	log.Print("starting server on :4000") //--> log starting server on port :4000 to the terminal

	err:= http.ListenAndServe(":4000", mux) //--> check for errors

	log.Fatal(err) //--> if there is an error log it to the terminal

}