package utils

import (
	"errors"
	"strconv"
	"strings"
)

func CPFvalidator(cpf string) error {
	cpf = onlyDigits(cpf)

	if len(cpf) != 11 {
		return errors.New("cpf deve ter 11 dígitos")
	}

	if CheckAllEqual(cpf) {
		return errors.New("Números iguais, CPF inválido")
	}

	if !CalculaVerificador1(cpf) {
		return errors.New("Primeiro dígito verificador inválido")
	}

	if !CalculaVerificador2(cpf) {
		return errors.New("Segundo dígito verificador inválido")
	}

	return nil
}

func CalculaVerificador1(cpf string) bool {
	digits := strings.Split(cpf, "")
	if len(digits) < 10 {
		return false
	}

	soma := 0
	for i := 0; i < 9; i++ {
		n, err := strconv.Atoi(digits[i])
		if err != nil {
			return false
		}
		soma += n * (10 - i)
	}

	dv := (soma * 10) % 11
	if dv == 10 || dv == 11 { // 👈 corrigido
		dv = 0
	}

	return strconv.Itoa(dv) == digits[9]
}

func CalculaVerificador2(cpf string) bool {
	digits := strings.Split(cpf, "")
	if len(digits) < 11 {
		return false
	}

	soma := 0
	for i := 0; i < 10; i++ {
		n, err := strconv.Atoi(digits[i])
		if err != nil {
			return false
		}
		soma += n * (11 - i)
	}

	dv := (soma * 10) % 11
	if dv == 10 || dv == 11 { // 👈 corrigido
		dv = 0
	}
	return strconv.Itoa(dv) == digits[10]
}

// verifica se o cpf nao e igual
func CheckAllEqual(cpf string) bool {
	if len(cpf) == 0 {
		return false
	}
	first := cpf[0]
	for i := 1; i < len(cpf); i++ {
		if cpf[i] != first {
			return false
		}
	}
	return true
}

// remove qualquer coisa que nao seja numero
func onlyDigits(s string) string {
	var b strings.Builder

	b.Grow(len(s))

	for _, r := range s {
		if r >= '0' && r <= '9' {
			b.WriteRune(r)
		}
	}
	return b.String()
}
