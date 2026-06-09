package models

import (
	"errors"
	"strings"
	"time"
)

type Book struct {
	ID          int64     `json:"id"`
	UserID      int64     `json:"user_id"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	Description string    `json:"description,omitempty"`
	Publisher   string    `json:"publisher,omitempty"`
	Year        int       `json:"year,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
}

func (b *Book) Prepare() error {
	if err := b.Validate(); err != nil {
		return err
	}
	b.Format()
	return nil
}

func (b *Book) Validate() error {
	if strings.TrimSpace(b.Title) == "" {
		return errors.New("título é obrigatório")
	}
	if strings.TrimSpace(b.Author) == "" {
		return errors.New("autor é obrigatório")
	}
	return nil
}

func (b *Book) Format() {
	b.Title = strings.TrimSpace(b.Title)
	b.Author = strings.TrimSpace(b.Author)
	b.Description = strings.TrimSpace(b.Description)
	b.Publisher = strings.TrimSpace(b.Publisher)
}
