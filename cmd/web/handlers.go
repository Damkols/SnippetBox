package main

import (
    "strconv"
    "fmt"
    "net/http"
	"html/template"
	"log"
)

func home(w http.ResponseWriter, r *http.Request) {

    w.Header().Add("Server", "Go") //--> setting response header map, header name: Server, header value: Go

	files:= []string{
		"./ui/html/base.tmpl.html",
		"./ui/html/pages/home.tmpl.html",
		"./ui/html/pages/title.tmpl.html",
	 } //--> path to template files

	ts, err := template.ParseFiles(files...) //--> Parse home tmpl file

	if err != nil {
		log.Print(err.Error())
		http.Error(w,"Internal Server Error", http.StatusInternalServerError)
		return
	} //-->  catch error if any

	err = ts.ExecuteTemplate(w, "base", nil) //--> execute parsed tmpl file

	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	} //--> catch error if any

}

func snippetView(w http.ResponseWriter, r *http.Request) {

    id, err := strconv.Atoi(r.PathValue("id")) //--> extract wildcard Id value and conv to integer

    if err != nil || id < 1 { //--> if err is not nill and id is not greater than 1 return NotFound
        http.NotFound(w,r)
        return
    }

    fmt.Fprintf(w, "Display a specific snippet with ID %d", id)
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {

    w.WriteHeader(http.StatusCreated) //--> using WriteHeader to send status codes

    w.Write([]byte ("Create a snippet")) //--> Displays Create a snippet
}
