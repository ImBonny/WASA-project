package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type unbanUserRequest struct {
	bannedUser    string `json:"bannedUser"`
	unbanningUser string `json:"username"`
}

type unbanUserResponse struct {
	Message string `json:"message"`
}

// Handler for unbanning a user
func unbanUser(w http.ResponseWriter, r *http.Request) {
	// Create a new unban request
	var unbanReq unbanUserRequest
	// Decode the request body into unbanReq
	if err := json.NewDecoder(r.Body).Decode(&unbanReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	unbanThisUser(unbanReq.bannedUser, unbanReq.unbanningUser)
	unbanResponse := unbanUserResponse{Message: "Successfully unbanned the user"}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(unbanResponse)
}

// Unban a user
func unbanThisUser(bannedUser string, unbanningUser string) error {
	if _, exists := users[bannedUser]; !exists {
		return fmt.Errorf("User with %s username not found", bannedUser)
	}
	userIndex := -1
	for i, user := range users[unbanningUser].BannedUsers {
		if user == bannedUser {
			userIndex = i
			break
		}
	}
	if userIndex != -1 {
		CurrentUser.BannedUsers = append(users[unbanningUser].BannedUsers[:userIndex], users[unbanningUser].BannedUsers[userIndex+1:]...)
		users[unbanningUser] = CurrentUser
		return nil
	}

	return fmt.Errorf("User with %s username not found in banned users of %s", bannedUser, unbanningUser)
}
