package models

import (
	"Api-Aula1-golang/security"
	"Api-Aula1-golang/utils"
	"errors"
	"strings"

	"github.com/badoux/checkmail"
)

type User struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	CPF   string `json:"cpf"`
	Email string `json:"email"`
	Senha string `json:"senha"`
}

func (u *User) Prepare(step string) error {
	if err := u.Validate(step); err != nil {
		return err
	}

	if err := u.Format(step); err != nil {
		return err
	}

	return nil
}

func (u *User) Validate(step string) error {
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

	if err := checkmail.ValidateFormat(u.Email); err != nil {
		return err
	}

	if err := utils.CPFvalidator(u.CPF); err != nil {
		return err
	}

	return nil
}

func (u *User) Format(step string) error {
	u.Name = strings.TrimSpace(u.Name)
	u.Email = strings.TrimSpace(u.Email)
	u.CPF = strings.TrimSpace(u.CPF)

	u.Name = strings.ToLower(u.Name)
	u.Email = strings.ToLower(u.Email)

	if step == "create" && u.Senha != "" {
		hashedPassword, err := security.Hash(u.Senha)
		if err != nil {
			return err
		}
		u.Senha = string(hashedPassword)
	}

	return nil
}
