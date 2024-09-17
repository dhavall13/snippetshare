package models

import (
	"database/sql"
	"time"
)

type Snippet struct {
	ID        int
	Title     string
	Content   string
	CreatedAt time.Time
	ExpiresAt time.Time
}

// SnippetModel this wraps sql.Db pool
type SnippetModel struct {
	DB *sql.DB
}

// Insert This will insert snippet into model
func (m *SnippetModel) Insert(title string, content string, expiresAt int) (int, error) {
	return 0, nil
}

func (m *SnippetModel) Get(id int) (*Snippet, error) {
	return nil, nil
}

func (m *SnippetModel) Latest() ([]*Snippet, error) {
	return nil, nil
}
