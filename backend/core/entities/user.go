package entities

import (
	"errors"
	"log"
	"net/mail"
	"time"
)

/*
developers are used to name the tables depending
on the language that interact the database in our
case we should be dealing whit Pascal/CamelCase
since GoLang use camelCase
-----------------------------------------------
snake_case rend les noms plus faciles à manipuler
et à comprendre, tout en évitant des pièges liés
à la casse et à la compatibilité entre différents
systèmes ansi que le mapping (ORM).
               Référence :
* Martin Fowler, Refactoring: Improving the Design
of Existing Code, Addison-Wesley, 2018.

* PostgreSQL Documentation sur la convention de
nommage des identifiants : PostgreSQL Naming
*/

type User struct {
	ID          int       `json:"id" db:"id"`
	Username    string    `json:"username" db:"username"`
	FirstName   string    `json:"first_name" db:"first_name"`
	LastName    string    `json:"last_name" db:"last_name"`
	Email       string    `json:"email" db:"email"`
	Password    string    `json:"password" db:"password"`
	DateOfBirth string    `json:"date_of_birth" db:"date_of_birth"`
	AboutMe     string    `json:"about_me" db:"about_me"`
	City        string    `json:"city" db:"city"`
	CoverPic    string    `json:"cover_pic" db:"cover_pic"` //#yagni
	ProfilePic  string    `json:"profile_pic" db:"profile_pic"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"` //#yagni
	Friend      []Friends `json:"friends"`
}

func (u User) validateEmail() error {
	_, err := mail.ParseAddress(u.Email)
	if err != nil {
		return errors.New("email invalide:" + string(err.Error()))
	}
	return nil
}

func (u *User) validateNickname() error {
	if len(u.FirstName) > 18 {
		return errors.New("le nickname ne peut pas dépasser 18 caractères")
	}
	return nil
}

func (u *User) Validate() error {
	if err := u.validateEmail(); err != nil {
		log.Println(err)
		return err
	}
	if err := u.validateNickname(); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
