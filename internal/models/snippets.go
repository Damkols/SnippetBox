package models

import (
	"database/sql"
	"time"
)

type Snippet struct { //--> Snippet type to hold data for an individual snippet
	ID int
	Title string
	Content string
	Created time.Time
	Expires time.Time
}

type SnippetModel struct {
	DB *sql.DB
}

func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	return 0, nil
} //--> func inserts a new snippet into the database

func (m *SnippetModel) Get(id int) (Snippet, error) {
	return Snippet{}, nil
} //--> func returns a specific snippet based on its ID

func (m *SnippetModel) Latest() ([]Snippet, error) {
	return nil, nil
} //--> func returns 10 most recently created snippets