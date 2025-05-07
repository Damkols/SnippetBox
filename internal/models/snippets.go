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