package main

import (
    "strconv"
    "fmt"
    "net/http"
	"errors"
	"strings"
	"unicode/utf8"

	"snippetbox.usmkols.net/internal/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) { //--> method of type application

	snippets, err := app.snippets.Latest() //--> get snippets from Latest() method
	if err != nil{
		app.serverError(w, r, err)
		return
	}

	data := app.newTemplateData(r)
	data.Snippets = snippets

	app.render(w, r, http.StatusOK, "home.tmpl.html", data)

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
	data.Snippet = snippet

	app.render(w, r, http.StatusOK, "view.tmpl.html", data)
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	app.render(w, r, http.StatusOK, "create.tmpl.html", data)
}

func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm() //--> r.ParseForm adds any data in POST request body to r.PostForm map
		if err != nil {
		app.clientError(w, r, http.StatusBadRequest)
		return
	}

	title := r.PostForm.Get("title") //--> use GET method on r.PostForm to get title from PostForm map
	content := r.PostForm.Get("content") //--> use GET method on r.PostForm to get content from PostForm map

	expires, err := strconv.Atoi(r.PostForm.Get("expires")) //--> use GET method on r.PostForm to get expires from PostForm map and convert to int
	if err != nil {
		app.clientError(w, r, http.StatusBadRequest)
		return
	}

	fieldErrors := make(map[string]string) //--> initialize a map to hold any validation errors


	id, err := app.snippets.Insert(title, content, expires) //--> Pass dummy data to SnippetModel.Insert() method and get ID back
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
   
}
