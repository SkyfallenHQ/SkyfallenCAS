package structures

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

type User struct {

	Name string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
	Username string `json:"username" form:"username"`
	Role string `json:"role"`
	Password string `json:"password" form:"password"`

}

func(u *User) Validate() error {

	if u.Role == "" {
		u.Role = "member"
	}

	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	if u.Name == "" {

		return errors.New("the name is required")

	}

	if u.Username == "" {

		return errors.New("the username is required")

	}

	if u.Password == "" {

		return errors.New("the password is required")

	}

	if !emailRegex.MatchString(u.Email) {

		return errors.New("the email is invalid")

	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("password could not be hashed")
	}

	u.Password = string(hashedPassword)

	return nil

}