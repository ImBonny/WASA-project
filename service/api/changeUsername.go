package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type changeRequest struct {
	Username string `json:"username"`
}

type changeResponse struct {
	Message string `json:"message"`
}

// Handler for changing username
func changeUsername(w http.ResponseWriter, r *http.Request) {
	// Create a new change request
	var changeReq changeRequest
	// Decode the request body into changeReq
	if err := json.NewDecoder(r.Body).Decode(&changeReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	changeToNewUsername(changeReq.Username)
	changeResponse := changeResponse{Message: "Successfully changed username"}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(changeResponse)
}

// Change the username of the current user
func changeToNewUsername(username string) error {
	if _, exists := users[username]; !exists {
		delete(users, CurrentUser.Username)
		CurrentUser.Username = username
		users[username] = CurrentUser
		return nil
	}

	return fmt.Errorf("User with username %s already exists", username)
}
