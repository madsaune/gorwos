package main

import (
	"strings"
	"testing"
	"unicode"
)

func IsUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func IsLower(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func TestCount(t *testing.T) {
	password, err := newRandomString(&options{
		uppercase: true,
		lowercase: true,
		digits:    true,
		symbols:   true,
		length:    16,
		count:     1,
	})

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if len(password) != 1 {
		t.Errorf("Password Count is %d; want 1", len(password))
	}
}

func TestOnlyUppercase(t *testing.T) {
	password, err := newRandomString(&options{
		uppercase: true,
		length:    16,
		count:     1,
	})

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !IsUpper(password[0]) {
		t.Errorf("Password is %s; want %s", password[0], strings.ToUpper(password[0]))
	}
}

func TestOnlyLowercase(t *testing.T) {
	password, err := newRandomString(&options{
		lowercase: true,
		length:    16,
		count:     1,
	})

	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !IsLower(password[0]) {
		t.Errorf("Password is %s; want %s", password[0], strings.ToLower(password[0]))
	}
}
