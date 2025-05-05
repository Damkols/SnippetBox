package main

import (
    "log"
    "net/http"
    "flag"
    "log/slog"
    "os"
)

func main() {

    addr := flag.String("addr", ":4000", "HTTP network address") //--> command line flags

    flag.Parse()

    logger:= slog.New(slog.NewTextHandler(os.Stdout, nil)) //--> initializing a structured logger

	fileServer := http.FileServer(http.Dir("./ui/static/")) //--> get static files

    mux:= http.NewServeMux() //--> creates a routing system 
	
	mux.Handle("GET /static/", http.StripPrefix("/static",fileServer)) //--> Get static files and use strip prefix to strip leading /static

    mux.HandleFunc("GET /{$}", home) //--> maps / path to home handler

    mux.HandleFunc("GET /snippet/view/{id}", snippetView) //--> maps /snippet/view to snippetView handler, uses {id} wildcard segment

    mux.HandleFunc("POST /snippet/create", snippetCreate) // --> maps /snippet/create to snippetCreate handler

    logger.Info("starting server", "addr", *addr) //--> log starting server on port :4000 to the terminal

    err:= http.ListenAndServe(*addr, mux) //--> check for errors

    log.Fatal(err) //--> if there is an error log it to the terminal

}
