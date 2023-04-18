package test

import (
	"backend"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// test the hash function
func TestHashStr(t *testing.T) {
	passwordTest := "TestPassword11@@"
	hashed := backend.HashStr(passwordTest)
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(passwordTest))
	if err != nil {
		t.Errorf("Hashed password does not match")
	}
}

// Test for contains function
// func TestContains(t *testing.T) {
// 	slice := []uint{1, 2, 3, 4, 5}
// 	val := uint(3)
// 	if !Contains(slice, val) {
// 		t.Errorf("Contains function does not work")
// 	}
// }

// func TestRemoveFromSlice(t *testing.T) {
// 	slice := []uint{1, 2, 3, 4, 5}
// 	val := uint(3)
// 	slice = RemoveFromSlice(slice, val)
// 	if Contains(slice, val) {
// 		t.Errorf("RemoveFromSlice function does not work")
// 	}
// }

// test for block and unblock user function
// func TestBlockUser(t *testing.T) {
// 	user := User{}
// 	targetID := uint(1)
// 	user.Blocked = BlockUser(user, targetID)
// 	if !Contains(user.Blocked, targetID) {
// 		t.Errorf("BlockUser function does not work")
// 	}

// 	user.Blocked = BlockUser(user, targetID)
// 	BlockUser(user, targetID)
// 	if Contains(user.Blocked, targetID) {
// 		t.Errorf("BlockUser function does not work")
// 	}
// }

// handeler for testing GET request
func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	// retrieve list of users from database
	users := []backend.User{{Name: "Nick"}, {Name: "Larry"}}

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
	var users []backend.User
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

var users []backend.User

// handeler for testing POST request
func createUserHandler(w http.ResponseWriter, r *http.Request) {
	// decode JSON request body into User object
	var user backend.User
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

// test POST for creation of users
func TestCreateUser(t *testing.T) {
	// create test server with createUserHandler as handler for /users endpoint
	server := httptest.NewServer(http.HandlerFunc(createUserHandler))
	defer server.Close()

	// create new user object
	newUser := backend.User{Name: "Charlie"}

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
	var createdUser backend.User
	err = json.NewDecoder(resp.Body).Decode(&createdUser)
	if err != nil {
		t.Fatalf("Error decoding response JSON: %v", err)
	}

	// verify that the created user matches the expected user
	expectedUser := backend.User{Name: "Nick"}
	if createdUser.ID != expectedUser.ID {
		t.Fatalf("Expected created user to be %+v, but got %+v", expectedUser.ID, createdUser.ID)
	}
}
func TestAuthUser(t *testing.T) {
	// initialize Gin router and database
	//r := gin.Default()
	db, err := gorm.Open(sqlite.Open("test.sqlite"), &gorm.Config{})

	// add test user to database
	user := backend.User{
		Name: "Nick",
	}
	db.Create(&user)

	// set up test context with JWT cookie
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, backend.Claims{
		ID: user.ID,
	})
	jwtCookie, _ := token.SignedString(main.JwtKey)
	req, _ := http.NewRequest("GET", "/", nil)
	req.AddCookie(&http.Cookie{Name: "jwt", Value: jwtCookie})
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// call AuthUser function with test context and database
	authUser, claims, err := backend.AuthUser(c, db)

	if err != nil {
		//t.Fatalf("Error authenticating user: %v", err)
	}

	// check for correct user informationif user.ID != authUser.ID {
	if user.ID != authUser.ID {
		t.Fatalf("User ID incorrect")
	}

	// check for correct claims information
	if user.ID != claims.ID {
		t.Fatalf("Claims ID incorrect")
	}
}
