package models

import (
	"errors"
	"fmt"
	"strings"
)

const MinPasswordLength = 8
const MinLoginLength = 6

type InputSignUp struct {
	Login    string `bind:"required"`
	Password string `bind:"required"`
	Name     string `bind:"required"`
	Email    string `bind:"required"`
}

func (i *InputSignUp) IsValid() error {
	if len(i.Login) < MinLoginLength {
		return errors.New(fmt.Sprintf("Login min length %d symbols", MinPasswordLength))
	}
	if strings.Contains(i.Login, "@") {
		return errors.New(`login have forbidden symbol "@" `)
	}
	if len(i.Password) < MinPasswordLength {
		return errors.New(fmt.Sprintf("password min length %d symbols", MinPasswordLength))
	}
	return nil
}

func (i *InputSignIn) Validate() error {
	if i.Identifier == "" {
		return errors.New("login required")
	}
	if i.Password == "" {
		return errors.New("password required")
	}
	return nil
}

type OutputSignUp struct {
	Session string `json:"session"`
}

type InputSignIn struct {
	Identifier string
	Password   string
}

type OutputSignIn struct {
	Token   string  `json:"token"`
	Account Account `json:"account"`
}
