package domain

import (
	"fmt"
	"regexp"
	"strings"
)

type Contact struct {
	ID          int
	FirstName   string
	MiddleName  string
	LastName    string
	PhoneNumber string
}

func NewContact(id int, phoneNumber, firstName, middleName, lastName string) (*Contact, error) {
	if !isVaildPhoneNumber(phoneNumber) {
		return nil, fmt.Errorf("invalid phone number %s", phoneNumber)
	}

	return &Contact{
		ID:          id,
		FirstName:   firstName,
		MiddleName:  middleName,
		LastName:    lastName,
		PhoneNumber: phoneNumber,
	}, nil
}

func (c *Contact) FullName() string {
	return strings.Join([]string{c.FirstName, c.MiddleName, c.LastName}, " ")
}

func isVaildPhoneNumber(phoneNumber string) bool {
	matched, _ := regexp.MatchString("^[0-9]+$", phoneNumber)
	return matched
}
