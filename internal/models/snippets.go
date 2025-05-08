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

	stmt := `INSERT INTO snippets (title, content, created, expires) VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))` //--> SQL statement to insert data to our database

	result, err := m.DB.Exec(stmt, title, content, expires) //--> use Exec method to execute statement

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId() //--> gets the ID of inserted data
	if err != nil {
		return 0, err
	}

	return int(id), nil

} //--> func inserts a new snippet into the database

func (m *SnippetModel) Get(id int) (Snippet, error) {

	stmt := `SELECT id, title, content, created, expires FROM snippets WHERE expires > UTC_TIMESTAMP() AND id = ?` //--> SQL statement to get data from database

	row := m.DB.QueryRow(stmt, id) //--> Using QueryRow method on connection pool to execute the SQL statement

	return Snippet{}, nil
} //--> func returns a specific snippet based on its ID

func (m *SnippetModel) Latest() ([]Snippet, error) {
	return nil, nil
} //--> func returns 10 most recently created snippets