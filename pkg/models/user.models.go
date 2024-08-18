package models

import (
	"gorm.io/gorm"
	"strings"
)

type User struct {
	gorm.Model
	Initial   bool `gorm:"-"` // unpersisted field
	Name      string
	Email     string
	AddressID uint
	Address   Address
}

type Address struct {
	gorm.Model
	Name string
}

func (u *User) Validate() (msgs []string) {
	if u.Initial {
		return
	}
	msgs = append(msgs, u.ValidateName()...)
	msgs = append(msgs, u.ValidateEmail()...)
	msgs = append(msgs, u.ValidateAddress()...)
	return msgs
}

func (u *User) ValidateName() (msgs []string) {
	if u.Initial {
		return
	}
	if u.Name == "" {
		msgs = append(msgs, "Name is required")
	}
	return msgs
}

func (u *User) NameHasError() bool {
	return len(u.ValidateName()) > 0
}

func (u *User) ValidateEmail() (msgs []string) {
	if u.Initial {
		return
	}
	//slog.Info("ValidateEmail", "u.Email", u.Email)
	if u.Email == "" || !strings.Contains(u.Email, "@") {
		return append(msgs, "Email is required")
	}
	return msgs
}

func (u *User) EmailHasError() bool {
	return len(u.ValidateEmail()) > 0
}

func (u *User) ValidateAddress() (msgs []string) {
	if u.Initial {
		return
	}
	// TODO(BW-18.08.24): Implement validation
	//if u.Address == "" || !strings.Contains(u.Address, "@") {
	//	return append(msgs, "Address is required")
	//}
	return msgs
}

func (u *User) AddressHasError() bool {
	return len(u.ValidateAddress()) > 0
}
