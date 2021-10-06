package entities

import (
	"errors"
	ERROR "go_clean_api/api/shared/constants"
	"go_clean_api/api/shared/utils"
	v "go_clean_api/api/shared/validators"
	"time"
)

/**
* As tags json e bson indicam que um User pode ser serializado nestes formatos
**/

type User struct {
	Base     `valid:"required"`
	Name     string `json:"name" bson:"name"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"-" bson:"-"`
	Token    string `json:"token" bson:"token"`
}

func BuildUser(name string, email string, password string) (*User, error) {
	user := &User{
		Name:     name,
		Email:    email,
		Password: password,
	}

	err := user.Prepare()
	if err != nil {
		return nil, err
	}

	return user, nil
}

/*
1 - O trecho user *User antes do nome da função representa uma amarração entre a função
Prepare e a struc User, ou seja, é como se User fosse uma classe e a função Prepare
fosse um método público.
2 - É usado o * no User, pois todas as vezes que o user for alterado ele será atualizado
em todos os objetos pois esta utilizando o mesmo local na memória
3 - O retorno da função é um error que pode ter valor nil, ou seja, se o erro for nil quer dizer
que a função funcionou corretamente. Esta é uma forma de validação.
4 - A função Prepare começa com letra maiuscula pois é um método publico. Funções que começam com
letra minuscula são funcões privadas.
*/

func (u *User) Prepare() error {

	err := u.validate()
	if err != nil {
		return err
	}

	u.ID = NewID()
	u.CreatedAt = time.Now()
	u.Password = string(u.Password)
	//Esse token deve ser gerado no bussinees
	u.Token = utils.GenerateUUIDV4()

	return nil
}

func (u *User) validate() error {
	if v.IsEmpty(u.Name) {
		return errors.New(ERROR.NAME_IS_EMPTY)
	}

	if !v.IsValidEmail(u.Email) {
		return errors.New(ERROR.INVALID_EMAIL)
	}

	if !v.IsValidPassword(u.Password) {
		return errors.New(ERROR.INVALID_PASSWORD)
	}

	return nil
}
