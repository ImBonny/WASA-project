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
	Identifier uint64 `json:"Identifier"`
}

// users is a map of username to User
var users = map[string]User{}

// getCurrentUser() is the currently logged in user
var currentUser User

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
	err := error(nil)

	currentUser.UserId, err = rt.db.DoLogin(loginReq.Name)

	// Create the response body
	response := LoginResponse{Identifier: currentUser.UserId}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
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

// getCurrentUser() returns the currently logged in user
func getCurrentUser() User {
	return currentUser
}

func setCurrentUser(user User) {
	currentUser = user
}
func setUsername(username string) {
	currentUser.Username = username
}

func addPost(id int) {
	currentUser.Profile.Posts = append(currentUser.Profile.Posts, id)
	currentUser.Profile.NumberOfPhotos++
	users[currentUser.Username] = currentUser
}

func deletePost(id int) {
	for i, postId := range currentUser.Profile.Posts {
		if postId == id {
			currentUser.Profile.Posts = append(currentUser.Profile.Posts[:i], currentUser.Profile.Posts[i+1:]...)
		}
	}
	return
}

func updateUser(user User) {
	currentUser = user
}
