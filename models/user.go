package models

import (
	"github.com/alexandrevicenzi/unchained"
)

type User struct {
	Id        int
	Username  string `db:"username"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Password  string `db:"password"`
	Operator  `db:"op"`
}

func (u *User) CheckPassword(pass string) bool {
	valid, _ := unchained.CheckPassword(pass, u.Password)
	if valid {
		return true
	} else {
		return false
	}
}
