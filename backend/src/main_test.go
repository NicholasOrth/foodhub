package main

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestHashStr(t *testing.T) {
	passwordTest := "TestPassword11@@"
	hashed := HashStr(passwordTest)
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(passwordTest))
	if err != nil {
		t.Errorf("Hashed password does not match")
	}
}
