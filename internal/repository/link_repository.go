package repository

import (
	"database/sql"
	"fmt"

	"github.com/fririz/URLShortener/domain"
	_ "github.com/mattn/go-sqlite3"
)

type LinkRepository struct {
	db *sql.DB
}

func NewLinkRepository(storagePath string) (*LinkRepository, error) {
	db, err := sql.Open("sqlite3", storagePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// 	type Link struct {
	// 	ID        int    `json:"id"`
	// 	URL       string `json:"url"`
	// 	Slug      string `json:"slug"`
	// 	CreatedAt string `json:"created_at"`
	// 	Visits    int    `json:"visits"`
	// }

	query := `
    CREATE TABLE IF NOT EXISTS links (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        url TEXT UNIQUE NOT NULL,
        slug TEXT NOT NULL,
        created_at TEXT NOT NULL,
        visits INTEGER NOT NULL DEFAULT 0
    );`

	_, err = db.Exec(query)
	if err != nil {
		return nil, fmt.Errorf("failed to create table: %w", err)
	}

	return &LinkRepository{db: db}, nil
}

func (lr *LinkRepository) AddLink(link *domain.Link) error {
	query := "INSERT INTO links (id, url, slug, created_at) VALUES (?, ?, ?, ?)"
	_, err := lr.db.Exec(query, link.ID, link.URL, link.Slug, link.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (lr *LinkRepository) GetLinkById(id int) (*domain.Link, error) {
	link := &domain.Link{}
	query := `SELECT id, url, slug, created_at, visits FROM links WHERE id = ?`
	row := lr.db.QueryRow(query, id)
	err := row.Scan(&link.ID, &link.URL, &link.Slug, &link.CreatedAt, &link.Visits)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("link with id %d not found", id)
		}
		return nil, err
	}
	return link, nil
}

func (lr *LinkRepository) GetLastId() (int, error) {
	var lastId int
	query := "SELECT COALESCE(MAX(id), 0) FROM links"

	err := lr.db.QueryRow(query).Scan(&lastId)
	if err != nil {
		return 0, err
	}
	return lastId, nil
}
