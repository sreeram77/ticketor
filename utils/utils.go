package utils

import "net/mail"

func CheckEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return false
	}

	return true
}
