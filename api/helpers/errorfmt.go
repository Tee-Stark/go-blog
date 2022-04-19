package helpers

import (
	"errors"
	"strings"
)

func FormatError(err string) error {
	if strings.Contains(err, "username") {
		return errors.New("Username already taken")
	}

	if strings.Contains(err, "email") {
		return errors.New("Email Already Taken")
	}

	if strings.Contains(err, "title") {
		return errors.New("Title Already Taken")
	}
	if strings.Contains(err, "hashedPassword") {
		return errors.New("Incorrect Password")
	}
	return errors.New("Incorrect Details")
}
