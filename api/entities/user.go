package entities

import (
	"errors"
	v "go_clean_api/api/shared/validators"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"

	uuid "github.com/satori/go.uuid"
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

func (user *User) Prepare() error {
	err := user.validate()
	if err != nil {
		// log.Fatalf("Error during the user validation: %v", err)
		return err
	}

	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		log.Fatalf("Error during the password generation: %v", err)
		return err
	}

	user.ID = uuid.NewV4().String()
	user.CreatedAt = time.Now()
	user.Password = string(password)
	user.Token = uuid.NewV4().String()

	return nil
}

func (u *User) validate() error {
	if v.IsEmpty(u.Name) {
		return errors.New("name is empty")
	}

	if !v.IsValidEmail(u.Email) {
		return errors.New("email is invalid")
	}

	if v.IsValidPassword(u.Password) {
		return errors.New("password is invalid")
	}

	return nil
}
