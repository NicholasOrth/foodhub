package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Test for contains function
func TestContains(t *testing.T) {
	slice := []uint{1, 2, 3, 4, 5}
	val := uint(3)
	if !Contains(slice, val) {
		t.Errorf("Contains function does not work")
	}
}

func TestRemoveFromSlice(t *testing.T) {
	slice := []uint{1, 2, 3, 4, 5}
	val := uint(3)
	slice = RemoveFromSlice(slice, val)
	if Contains(slice, val) {
		t.Errorf("RemoveFromSlice function does not work")
	}
}

// test for block and unblock user function
func TestBlockUser(t *testing.T) {
	user := User{}
	targetID := uint(1)
	user.Blocked = BlockUser(user, targetID)
	if !Contains(user.Blocked, targetID) {
		t.Errorf("BlockUser function does not work")
	}

	user.Blocked = BlockUser(user, targetID)
	BlockUser(user, targetID)
	if Contains(user.Blocked, targetID) {
		t.Errorf("BlockUser function does not work")
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

// handeler for testing POST request
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

// test POST for creation of users
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

/*
func TestAuthUser(t *testing.T) {
	// initialize Gin router and database
	//r := gin.Default()
	db, err := gorm.Open(sqlite.Open("test.sqlite"), &gorm.Config{})

	// add test user to database
	user := User{
		Name: "Nick",
	}
	db.Create(&user)

	// set up test context with JWT cookie
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		ID: user.ID,
	})
	jwtCookie, _ := token.SignedString(jwtKey)
	req, _ := http.NewRequest("GET", "/", nil)
	req.AddCookie(&http.Cookie{Name: "jwt", Value: jwtCookie})
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// call AuthUser function with test context and database
	authUser, claims, err := AuthUser(c, db)

	if err == nil {
		t.Fatalf("Error authenticating user: %v", err)
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
*/

// Test for Signup function
func TestSignup(t *testing.T) {
	// Setup test environment
	r := gin.Default()
	db, _ := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	db.AutoMigrate(&User{})

	// Add route handler
	r.POST("/auth/signup", func(c *gin.Context) {
		p := &params{
			memory:      64 * 1024,
			iterations:  3,
			parallelism: 2,
			saltLength:  16,
			keyLength:   32,
		}

		// Bind request body to User struct
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		// Hash the user's password
		password, err := generateFromPassword(user.Password, p)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		// Store the user in the database
		res := db.Create(&User{
			Name:     user.Name,
			Email:    user.Email,
			Password: password,
		})

		if res.Error != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		c.JSON(http.StatusOK, nil)
	})

	// Create test user
	testUser := User{
		Name:     "Test User",
		Email:    "test@example.com",
		Password: "testpassword",
	}

	// Send request to create user
	w := httptest.NewRecorder()
	reqBody, _ := json.Marshal(testUser)
	req, _ := http.NewRequest("POST", "/auth/signup", bytes.NewReader(reqBody))
	r.ServeHTTP(w, req)

	// Check response status code
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, w.Code)
	}

	// Check that user was created in the database
	var user User
	res := db.First(&user, "email = ?", testUser.Email)
	if res.Error != nil {
		t.Errorf("Expected no error but got %v", res.Error)
	}
	if user.Name != testUser.Name {
		t.Errorf("Expected user name to be %s but got %s", testUser.Name, user.Name)
	}
}
