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
	user := User{
		/*
			//ID: 1,
			Name:      "TestUser1",
			Email:     "email1@email.com",
			Password:  "password1",
			Following: []uint{},
			Followers: []uint{},
		*/
	}
	target := User{
		/*
			//ID: 2,
			Name:      "TestUser2",
			Email:     "email2@email.com",
			Password:  "password2",
			Following: []uint{},
			Followers: []uint{},
		*/
	}
	//user.ID = 1
	//target.ID = 2
	//user.Following = append(user.Following, target.ID)
	AddFollower(&user, &target)
	if len(user.Following) != 1 {
		t.Errorf("User following list not updated, length: %d", len(user.Following))
	}
	if len(target.Followers) != 1 {
		t.Errorf("Target followers list not updated, length: %d", len(target.Followers))
	}
}
