package repository

import (
	"Api-Aula1-golang/models"
	"database/sql"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (u UserRepo) Create(user models.User) (int64, error) {
	query := `INSERT INTO users (name, email, password, cpf) VALUES (?, ?, ?, ?)`

	statement, err := u.db.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Email, user.Senha, user.CPF)
	if err != nil {
		return 0, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastId, nil
}

func (u UserRepo) FindAll() ([]models.User, error) {
	query := `SELECT id, name, email, cpf FROM users`

	rows, err := u.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.CPF); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (u UserRepo) FindByID(id int64) (models.User, error) {
	query := `SELECT id, name, email, cpf FROM users WHERE id = ?`

	row := u.db.QueryRow(query, id)

	var user models.User
	if err := row.Scan(&user.ID, &user.Name, &user.Email, &user.CPF); err != nil {
		return models.User{}, err
	}

	return user, nil
}
