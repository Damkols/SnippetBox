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

type snippetCreateForm struct {
	Title string
	Content string
	Expires int
	FieldErrors map[string]string
}

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
	
	data.Form = snippetCreateForm{
		Expires: 365
	}
	
	app.render(w, r, http.StatusOK, "create.tmpl.html", data)
}

func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm() //--> r.ParseForm adds any data in POST request body to r.PostForm map
		if err != nil {
		app.clientError(w, r, http.StatusBadRequest)
		return
	}

	expires, err := strconv.Atoi(r.PostForm.Get("expires")) //--> use GET method on r.PostForm to get expires from PostForm map and convert to int
	if err != nil {
		app.clientError(w, r, http.StatusBadRequest)
		return
	}

	form := snippetCreateForm{
		Title: r.PostForm.Get("title")
		Content: r.PostForm.Get("content")
		Expires: expires,
		FieldErrors: map[string]string{}
	}


	if strings.TrimSpace(title) == "" {
		form.FieldErrors["title"] = "This field cannot be blank"
	} else if utf8.RuneCountInString(title) > 100 {
		form.FieldErrors["title"] = "This field cannot be more than 100 characters long"
	}

	if form.Expires != 1 && expires != 7 && expires != 365 {
		form.FieldErrors["expires"] = "This field must equal 1, 7 or 365"
	}

	if len(fieldErrors) > 0 {
		data := app.newTemplateData(r)
		data.Form = form
		app.render(w, r, http.StatusUnprocessableEntity, "create.tmpl", data)
		return
	}

	id, err := app.snippets.Insert(form.Title, form.Content, form.Expires) //--> Pass dummy data to SnippetModel.Insert() method and get ID back
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
   
}
