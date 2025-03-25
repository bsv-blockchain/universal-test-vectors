package testabilities

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	emailRegex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
)

// Paymail wraps a string paymail address to provide additional common methods.
type Paymail string

// String is necessary to fulfill the fmt.Stringer interface.
func (p Paymail) String() string {
	return string(p)
}

// Address returns the address of this paymail.
func (p Paymail) Address() string {
	return string(p)
}

// PublicName returns the public name of this paymail (for testing purposes, it's just the alias in uppercase).
func (p Paymail) PublicName() string {
	return strings.ToUpper(p.Alias())
}

// Domain returns the domain of this paymail.
func (p Paymail) Domain() string {
	_, domain, _ := sanitizeEmail(p.String())
	return domain
}

// Alias returns the alias of this paymail.
func (p Paymail) Alias() string {
	alias, _, _ := sanitizeEmail(p.String())
	return alias
}

// DefaultPaymail returns the default paymail of this user.
func (f *User) DefaultPaymail() Paymail {
	if len(f.Paymails) == 0 {
		return ""
	}
	return f.Paymails[0]
}

func sanitizeEmail(email string) (string, string, error) {
	email = strings.TrimSpace(email)

	if email == "" {
		return "", "", fmt.Errorf("email is empty")
	}

	email = strings.ToLower(email)

	if !emailRegex.MatchString(email) {
		return "", "", fmt.Errorf("invalid email format")
	}

	parts := strings.Split(email, "@")
	alias, domain := parts[0], parts[1]

	return alias, domain, nil
}
