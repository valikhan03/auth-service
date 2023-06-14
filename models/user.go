package models

import (
	//"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int64     `db:"id"`
	FullName  string    `db:"full_name"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	PhoneNum  string    `db:"phone_num"`
	//BirthDate time.Time `db:"birth_date"`
}

func (u *User) HashPassword() error {
	bytes_pass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(bytes_pass)
	return nil
}

func (u *User) CheckPassword(providedPassword string) error {
	bytes_pass, err := bcrypt.GenerateFromPassword([]byte(providedPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), bytes_pass)
	return err
}
