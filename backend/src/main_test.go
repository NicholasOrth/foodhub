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

func TestAddFollower(t *testing.T) {
	user := User{}
	target := User{}

	user.Following = append(user.Following, target.ID)
	target.Followers = append(target.Followers, user.ID)

	if len(user.Following) != 1 {
		t.Errorf("User following list not updated, length: %d", len(user.Following))
	}
	if len(target.Followers) != 1 {
		t.Errorf("Target followers list not updated, length: %d", len(target.Followers))
	}
}
