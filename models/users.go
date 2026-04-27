package models

//aqui que coloco as regras de negocios que eu quero que as coisas funcionem

type User struct {
	id    int8
	name  string
	cpf   string
	email string
	senha string
}

//Crud
//1-Criar usuário

func (u *User) prepare() error {
	//chamamos o validate()
	//chamamos o format()

	return nil

}

func (u *User) validate(step string) error {
	//aqui vamos verificar se os campos recebidos do usuario nao estao vazios
	//validar. se o cpf e valido

	return nil
}

func (u *User) format(step string) error {
	//vamos formatar as strings, para remover espacos
	//depois vamos aplicar hash na senha

	return nil
}
