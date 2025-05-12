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
	
}