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
	if err != nil{
		app.serverError(w, r, err)
		return
	}

	data := app.newTemplateData(r)
	data.Snippets = snippets

	app.render(w, r, http.StatusOK, "home.tmpl", data)

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

	data := app.newTemplateData(r)
	data.Snippets = snippets

	app.render(w, r, http.StatusOK, "view.tmpl", data)

	// files:= []string{
	// 	"./ui/html/base.tmpl.html",
	// 	"./ui/html/partials/nav.tmpl.html",
	// 	"./ui/html/pages/home.tmpl.html",
	// 	"./ui/html/pages/view.tmpl.html",
	//  } //--> path to template files

	//  ts, err := template.ParseFiles(files...)
	//  if err != nil {
	// 	app.serverError(w, r, err)
	// 	return
	//  }

	//  data:= templateData{
	// 	Snippet: snippet,
	//  }

	//  err = ts.ExecuteTemplate(w, "base", data)//--> last parameter (snippet) represents dynamic data
	//  if err != nil {
	// 	app.serverError(w, r, err)
	//  }
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
