package utils

import "net/mail"

type Validation struct{}

func (v Validation) ValidateEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func (v Validation) ValidatePhoneNumber(phone string) bool {
	return len(phone) >= 10
}

func (v Validation) ValidatePassword(password string) bool {
	return len(password) >= 8
}
