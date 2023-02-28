package main

import (
	"bytes"
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

// handeler for testing GET request
func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	// retrieve list of users from database
	users := []User{{Name: "Nick"}, {Name: "Larry"}}

	// encode users as JSON and write to response
	json.NewEncoder(w).Encode(users)
}

// API tests
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

var users []User

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	// decode JSON request body into User object
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//add user to database
	users = append(users, user)

	// encode user as JSON and write to response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func TestCreateUser(t *testing.T) {
	// create test server with createUserHandler as handler for /users endpoint
	server := httptest.NewServer(http.HandlerFunc(createUserHandler))
	defer server.Close()

	// create new user object
	newUser := User{Name: "Charlie"}

	// encode user as JSON
	requestBody, err := json.Marshal(newUser)
	if err != nil {
		t.Fatalf("Error encoding request JSON: %v", err)
	}

	// make POST request to /users endpoint with JSON request body
	resp, err := http.Post(server.URL+"/users", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatalf("Error making POST request: %v", err)
	}
	defer resp.Body.Close()

	// verify that the response status code is 201 Created
	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("Expected status code 201 Created, but got %v", resp.StatusCode)
	}

	// decode response JSON into User object
	var createdUser User
	err = json.NewDecoder(resp.Body).Decode(&createdUser)
	if err != nil {
		t.Fatalf("Error decoding response JSON: %v", err)
	}

	// verify that the created user matches the expected user
	expectedUser := User{Name: "Nick"}
	if createdUser.ID != expectedUser.ID {
		t.Fatalf("Expected created user to be %+v, but got %+v", expectedUser.ID, createdUser.ID)
	}
}
