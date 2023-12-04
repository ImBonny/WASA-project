package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"regexp"
)

// LoginRequest represents the request body for user details
type LoginRequest struct {
	Name string `json:"name"`
}

// LoginResponse represents the response body for user details
type LoginResponse struct {
	Identifier string `json:"Identifier"`
}

// users is a map of username to User
var users = map[string]User{}

// CurrentUser is the currently logged in user
var CurrentUser User

// doLogin handles the login request
func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var loginReq LoginRequest
	// Decode the request body into loginReq
	if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Validate the name using regular expression pattern
	validName := regexp.MustCompile(`^.*$`)
	if !validName.MatchString(loginReq.Name) || len(loginReq.Name) < 3 || len(loginReq.Name) > 16 {
		http.Error(w, "Invalid username format", http.StatusBadRequest)
		return
	}
	// Create a new user if it doesn't exist
	createUser(loginReq.Name)
	CurrentUser = users[loginReq.Name]
	//TODO: check the logic behind the identifier
	identifier := "abcdef012345"
	// Create the response body
	response := LoginResponse{Identifier: identifier}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// createUser creates a new user if it doesn't exist
func createUser(username string) {
	// Check if the user already exists
	if _, exists := users[username]; !exists {
		// Create a new user
		newUser := User{
			Username: username,
			Profile: Profile{
				Posts:          []int{},
				NumberOfPhotos: 0,
				Followers:      []string{},
				Following:      []string{},
			},
			BannedUsers: []string{},
		}
		users[username] = newUser
	}
}

// getCurrentUser returns the currently logged in user
func getCurrentUser() User {
	return CurrentUser
}
