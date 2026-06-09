package repository

import (
	"Api-Aula1-golang/models"
	"database/sql"
)

type BookRepo struct {
	db *sql.DB
}

func NewBookRepo(db *sql.DB) *BookRepo {
	return &BookRepo{db: db}
}

func (b *BookRepo) Create(book models.Book) (int64, error) {
	query := `INSERT INTO books (user_id, title, author, description, publisher, year)
	          VALUES (?, ?, ?, ?, ?, ?)`

	stmt, err := b.db.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(book.UserID, book.Title, book.Author, book.Description, book.Publisher, book.Year)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func (b *BookRepo) FindByUser(userID int64) ([]models.Book, error) {
	query := `SELECT id, user_id, title, author, description, publisher, year, created_at
	          FROM books WHERE user_id = ? ORDER BY created_at DESC`

	rows, err := b.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var book models.Book
		if err = rows.Scan(&book.ID, &book.UserID, &book.Title, &book.Author,
			&book.Description, &book.Publisher, &book.Year, &book.CreatedAt); err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}

func (b *BookRepo) FindByID(bookID, userID int64) (models.Book, error) {
	query := `SELECT id, user_id, title, author, description, publisher, year, created_at
	          FROM books WHERE id = ? AND user_id = ?`

	var book models.Book
	err := b.db.QueryRow(query, bookID, userID).Scan(
		&book.ID, &book.UserID, &book.Title, &book.Author,
		&book.Description, &book.Publisher, &book.Year, &book.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return models.Book{}, nil
	}
	return book, err
}

func (b *BookRepo) Update(book models.Book) error {
	query := `UPDATE books SET title = ?, author = ?, description = ?, publisher = ?, year = ?
	          WHERE id = ? AND user_id = ?`

	stmt, err := b.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(book.Title, book.Author, book.Description, book.Publisher, book.Year, book.ID, book.UserID)
	return err
}

func (b *BookRepo) Delete(bookID, userID int64) error {
	query := `DELETE FROM books WHERE id = ? AND user_id = ?`

	stmt, err := b.db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(bookID, userID)
	return err
}
