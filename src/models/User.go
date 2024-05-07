package models

import (
	"api/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type User struct {
	ID uint64 `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Username string `json:"username,omitempty"`
	Email string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

func (user *User) Prepare(stage string) error {
	if error := user.validate(stage); error !=nil {
		return error
	}


	if 	error := user.format(stage); error != nil {
		return error
	}
	
	return nil
}


func (user *User) validate(stage string) error {
	if user.Name == "" {
		return errors.New("name cannot be empty")
	}
	if user.Username == "" {
		return errors.New("username cannot be empty")
	}
	if user.Email == "" {
		return errors.New("email cannot be empty")
	}

	if error := checkmail.ValidateFormat(user.Email); error != nil {
		return errors.New("email must be valid")
	}

	if stage == "create" && user.Password == "" {
		return errors.New("password cannot be empty")
	}

	return nil
}

func (user *User) format(stage string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Username = strings.TrimSpace(user.Username)
	user.Email = strings.TrimSpace(user.Email)

	if stage == "register" {
		hashedPassword, error := security.Hash(user.Password) 
		
		if error != nil {
			return error
		}

		user.Password = string(hashedPassword)
	}

	return nil
}