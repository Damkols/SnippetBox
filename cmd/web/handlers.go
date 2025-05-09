package main

import (
    "strconv"
    "fmt"
    "net/http"
	"html/template"
	"errors"

	"snippetbox.usmkols.net/internal/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) { //--> method of type application

    w.Header().Add("Server", "Go") //--> setting response header map, header name: Server, header value: Go

	snippets, err := app.snippets.Latest() //--> get snippets from Latest() method
	

	files:= []string{
		"./ui/html/base.tmpl.html",
		"./ui/html/pages/home.tmpl.html",
		"./ui/html/pages/title.tmpl.html",
		"./ui/html/partials/nav.tmpl.html",
	 } //--> path to template files

	ts, err := template.ParseFiles(files...) //--> Parse home tmpl file

	if err != nil {
		app.serverError(w, r, err)
		return
	} //-->  catch error if any

	err = ts.ExecuteTemplate(w, "base", nil) //--> execute parsed tmpl file

	if err != nil {
		app.serverError(w, r, err)
	} //--> catch error if any

}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {

    id, err := strconv.Atoi(r.PathValue("id")) //--> extract wildcard Id value and conv to integer

    if err != nil || id < 1 { //--> if err is not nil and id is not greater than 1 return NotFound
        http.NotFound(w,r)
        return
    }

	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			http.NotFound(w,r)
		} else {
			app.serverError(w, r, err)
		}
		return
	}

    fmt.Fprintf(w, "%+v", snippet) //--> write snippetdata as a plain-text HTTP response body
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {

	//--> Dummy data
	title := "0 snail"
	content := "0 snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n- Kobayashi Issa"
	expires := 7

	id, err := app.snippets.Insert(title, content, expires) //--> Pass dummy data to SnippetModel.Insert() method and get ID back
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
   
}
