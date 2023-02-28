package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

// test the hash function
func TestHashStr(t *testing.T) {
	passwordTest := "TestPassword11@@"
	hashed := HashStr(passwordTest)
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(passwordTest))
	if err != nil {
		t.Errorf("Hashed password does not match")
	}
}

// test the add follower function
// NOTE: not yet implemented as a built in function, instead just use append twice
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

func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	// retrieve list of users from database
	users := []User{{Name: "Nick"}, {Name: "Larry"}}

	// encode users as JSON and write to response
	json.NewEncoder(w).Encode(users)
}

// test GET for users
func TestGetUsers(t *testing.T) {
	// create test server with getUsersHandler as handler for /users endpoint
	server := httptest.NewServer(http.HandlerFunc(getUsersHandler))
	defer server.Close()

	// make GET request to /users endpoint
	resp, err := http.Get(server.URL + "/users")
	if err != nil {
		t.Fatalf("Error making GET request: %v", err)
	}
	defer resp.Body.Close()

	// decode response JSON into slice of User objects
	var users []User
	err = json.NewDecoder(resp.Body).Decode(&users)
	if err != nil {
		t.Fatalf("Error decoding response JSON: %v", err)
	}

	// verify that we received the expected list of users
	if len(users) != 2 {
		t.Fatalf("Expected 2 users, but got %d", len(users))
	}
	if users[0].Name != "Nick" {
		t.Fatalf("Expected first user to be Nick, but got %s", users[0].Name)
	}
	if users[1].Name != "Larry" {
		t.Fatalf("Expected second user to be Larry, but got %s", users[1].Name)
	}
}
