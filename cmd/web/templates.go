package main
import (
	"html/template"
	"path/filepath"
	"time"
	"snippetbox.usmkols.net/internal/models"
)


type templateData struct {
	CurrentYear int
	Snippet models.Snippet
	Snippets []models.Snippet
	templateCache map[string]*template.Template
}

func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

var functions = template.FuncMap{
	"humanDate": humanDate,
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{} //--> Initialize a new map to act as cache

	pages, err := filepath.Glob("./ui/html/pages/*.tmpl.html") //--> get all tmpl.html files in pages directory
	if err != nil {
		return nil, err
	}

	for _, page := range pages{ //--> map over pages

		name := filepath.Base(page) //--> extract filename e.g home.tmpl.html

		ts, err:= template.New(name).Funcs(functions).ParseFiles("ui/html/base.tmpl.html")
		if err != nil{
			return nil, err
		}

		ts, err = ts.ParseGlob("ui/html/partials/*.tmpl.html")
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseFiles(page)
		if err != nil{
			return nil, err
		}

		cache[name] = ts //-> add the template set to map
	}

	return cache, nil //--> return map
}