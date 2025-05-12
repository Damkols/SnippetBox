package main
import (
	"html/template"
	"path/filepath"
	"snippetbox.usmkols.net/internal/models"
)


type templateData struct {
	Snippet models.Snippet
	Snippets []models.Snippet
}


func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{} //--> Initialize a new map to act as cache

	pages, err := filepath.Glob("./ui/html/pages/*.tmpl.html") //--> get all tmpl.html files in pages directory
	if err != nil {
		return nil, err
	}

	
}