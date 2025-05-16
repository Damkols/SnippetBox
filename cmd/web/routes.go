package main


import (
    "net/http"
)

func (app *application) routes() http.Handler {
	fileServer := http.FileServer(http.Dir("./ui/static/")) //--> get static files

    mux:= http.NewServeMux() //--> creates a routing system 
	
	mux.Handle("GET /static/", http.StripPrefix("/static",fileServer)) //--> Get static files and use strip prefix to strip leading /static

    mux.HandleFunc("GET /{$}", app.home) //--> app.home creates a method value, maps / path to home handler also ensures when a request comes in home handler is able to use dependencies from the stored address

    mux.HandleFunc("GET /snippet/view/{id}", app.snippetView) //--> maps /snippet/view to snippetView handler, uses {id} wildcard segment

    mux.HandleFunc("GET /snippet/create", app.snippetCreate) // --> maps /snippet/create to snippetCreate handler

    // mux.HandleFunc("POST /snippet/createpost", app.snippetCreatePost) // --> maps /snippet/create to snippetCreate handler

	return app.recoverPanic(app.logRequest(commonHeaders(mux)))

}