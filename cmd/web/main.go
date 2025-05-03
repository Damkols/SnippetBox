package main

import (
    "log"
    "net/http"
)

func main() {

    mux:= http.NewServeMux() //--> creates a routing system 

    mux.HandleFunc("GET /{$}", home) //--> maps / path to home handler

    mux.HandleFunc("GET /snippet/view/{id}", snippetView) //--> maps /snippet/view to snippetView handler, uses {id} wildcard segment

    mux.HandleFunc("POST /snippet/create", snippetCreate) // --> maps /snippet/create to snippetCreate handler

    log.Print("starting server on :4000") //--> log starting server on port :4000 to the terminal

    err:= http.ListenAndServe(":4000", mux) //--> check for errors

    log.Fatal(err) //--> if there is an error log it to the terminal

}
