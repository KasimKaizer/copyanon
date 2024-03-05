package models

import (
	"database/sql"
	"errors"
	"time"
)

type Gist struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type GistModel struct {
	DB *sql.DB
}

func (g *GistModel) Insert(title, content string, expires int) (int, error) {

	stmt := `INSERT INTO gists (title, content, created, expires)
	VALUES (?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	result, err := g.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (g *GistModel) Get(id int) (Gist, error) {
	stmt := `SELECT id, title, content, created, expires FROM gists
	WHERE expires > UTC_TIMESTAMP() AND id = ?`

	row := g.DB.QueryRow(stmt, id)

	var n Gist

	err := row.Scan(&n.ID, &n.Title, &n.Content, &n.Created, &n.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Gist{}, ErrNoRecord
		} else {
			return Gist{}, err
		}
	}
	return n, nil
}

func (g *GistModel) Latest() ([]Gist, error) {
	stmt := `SELECT id, title, content, created, expires FROM gists
	WHERE expires > UTC_TIMESTAMP() ORDER BY id DESC LIMIT 10`

	rows, err := g.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var gists []Gist

	for rows.Next() {
		var n Gist

		err := rows.Scan(&n.ID, &n.Title, &n.Content, &n.Created, &n.Expires)
		if err != nil {
			return nil, err
		}
		gists = append(gists, n)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return gists, nil
}
