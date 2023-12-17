package api

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type unbanUserRequest struct {
	BannedUser    string `json:"bannedUser"`
	UnbanningUser string `json:"username"`
}

type unbanUserResponse struct {
	Message string `json:"message"`
}

// Handler for unbanning a user
func (rt *_router) unbanUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Create a new unban request
	var unbanReq unbanUserRequest
	unbanReq.UnbanningUser = ps.ByName("username")
	unbanReq.BannedUser = ps.ByName("bannedUser")
	// Decode the request body into unbanReq
	if err := json.NewDecoder(r.Body).Decode(&unbanReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var user User
	token := getToken(r.Header.Get("Authorization"))
	user.UserId = token

	//TODO: IMPLEMENT SECURITY ONCE I HAVE DB

	err := unbanThisUser(unbanReq.BannedUser, unbanReq.UnbanningUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	unbanResponse := unbanUserResponse{Message: "Successfully unbanned the user"}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(unbanResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

// Unban a user
func unbanThisUser(bannedUser string, unbanningUser string) error {
	if _, exists := users[bannedUser]; !exists {
		return fmt.Errorf("user with %s username not found", bannedUser)
	}
	userIndex := -1
	for i, user := range users[unbanningUser].BannedUsers {
		if user == bannedUser {
			userIndex = i
			break
		}
	}
	if userIndex != -1 {
		currentUser := users[unbanningUser]
		currentUser.BannedUsers = append(users[unbanningUser].BannedUsers[:userIndex], users[unbanningUser].BannedUsers[userIndex+1:]...)
		users[unbanningUser] = currentUser
		updateUser(users[unbanningUser])
		return nil
	}

	return fmt.Errorf("user with %s username not found in banned users of %s", bannedUser, unbanningUser)
}
