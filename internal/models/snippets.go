package models

import (
	"database/sql"
	"errors"
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

	var s Snippet //--> initialize a new Snippet

	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires) //--> error handling
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Snippet{}, ErrNoRecord
		} else {
			return Snippet{}, err
		}
	}

	return s, nil //--> return filled struct

} //--> func returns a specific snippet based on its ID

func (m *SnippetModel) Latest() ([]Snippet, error) {

	stmt := `SELECT id, title, content, created, expires FROM snippets WHERE expires > UTC_TIMESTAMP() ORDER BY id DESC LIMIT 10` //--> sql statemnet to get the last 10 snippets

	rows, err := m.DB.Query(stmt) //--> use Query method on db connection pool to excute sql stmt
	if err != nil {
		return nil, err
	} //--> erorr handling

	defer rows.Close() //--> defer rows.Close before Latest method returns something

	var snippets []Snippet //--> intializing an empty struct to hold snippet struct from DB

	for rows.Next() { //--> iterate through rows

		var s Snippet //--> new snippet struct to copy data into

		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires) //--> rows.Scan to copy values into Snippet struct
		if err != nil {
			return nil, err
		}

	}

} //--> func returns 10 most recently created snippets