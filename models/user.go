package models

import (
	"time"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int64
	FullName  string
	Email     string
	Password  string
	PhoneNum  string
	BirthDate time.Time
}

func (u *User) HashPassword() error {
	bytes_pass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil{
		return err
	}

	u.Password = string(bytes_pass)
	return nil
}


func (u *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(providedPassword))
	return err
}