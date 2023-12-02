package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type banRequest struct {
	bannedUser  string `json:"bannedUser"`
	banningUser string `json:"username"`
}

type banResponse struct {
	Message string `json:"message"`
}

// Handler for banning a user
func banUser(w http.ResponseWriter, r *http.Request) {
	// Create a new ban request
	var banReq banRequest
	// Decode the request body into banReq
	if err := json.NewDecoder(r.Body).Decode(&banReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	banThisUser(banReq.bannedUser, banReq.banningUser)
	banResponse := banResponse{Message: "Successfully banned the user"}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(banResponse)
}

// Ban a user
func banThisUser(bannedUser string, banningUser string) error {
	if _, exists := users[bannedUser]; !exists {
		return fmt.Errorf("User with %s username not found", bannedUser)
	}
	users[banningUser] = User{
		Username:    banningUser,
		Profile:     users[banningUser].Profile,
		BannedUsers: append(users[banningUser].BannedUsers, users[bannedUser].Username),
	}
	return nil
}
