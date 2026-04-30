package models

import (
	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

//aqui que coloco as regras de negocios que eu quero que as coisas funcionem

type User struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	CPF   string `json:"cpf"`
	Email string `json:"email"`
	Senha string `json:"senha"`
}

//Crud
//1-Criar usuário

func (u *User) Prepare(step string) error {
	//chamamos o validate()
	//chamamos o format()
	if err := u.Validate(step); err != nil {
		return err
	}

	if err := u.Format(step); err != nil {
		return err
	}

	return nil

}

func (u *User) Validate(step string) error {
	//aqui vamos verificar se os campos recebidos do usuario nao estao vazios
	//validar. se o cpf e valido

	if u.Name == "" {
		return errors.New("nome é obrigatório")
	}

	if u.Email == "" {
		return errors.New("email é obrigatório")
	}

	if u.CPF == "" {
		return errors.New("cpf é obrigatório")
	}

	if step == "create" && u.Senha == "" {
		return errors.New("senha é obrigatória")
	}
	if !strings.Contains(u.Email, "@") {
		return errors.New("email inválido")
	}

	return nil

}

func (u *User) Format(step string) error {
	u.Name = strings.TrimSpace(u.Name)
	u.Email = strings.TrimSpace(u.Email)
	u.CPF = strings.TrimSpace(u.CPF)

	if step == "create" {
		hashed, err := hashPassword(u.Senha)
		if err != nil {
			return err
		}
		u.Senha = string(hashed)
	}

	if u.Senha != "" {
		hashed, err := hashPassword(u.Senha)
		if err != nil {
			return err
		}
		u.Senha = string(hashed)
	}

	return nil
}

func hashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}
