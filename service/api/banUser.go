package api

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type banRequest struct {
	BannedUser string `json:"bannedUser"`
}

type banResponse struct {
	Username string `json:"username"`
}

func (rt *_router) banUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Create a new ban request
	var banReq banRequest
	// Decode the request body into banReq
	if err := json.NewDecoder(r.Body).Decode(&banReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var user User
	token := getToken(r.Header.Get("Authorization"))
	user.UserId = token

	//TODO: IMPLEMENT SECURITY ONCE I HAVE DB

	myUsername := getCurrentUser().Username
	banReq.BannedUser = ps.ByName("bannedUser")
	err := banThisUser(myUsername, banReq)
	if err != nil {
		return
	}
	banResponse := banResponse{Username: banReq.BannedUser}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(banResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
func banThisUser(banningUser string, request banRequest) error {
	if _, exists := users[request.BannedUser]; !exists {
		return fmt.Errorf("User with %s username not found", request.BannedUser)
	}
	users[banningUser] = User{
		Username:    banningUser,
		Profile:     users[banningUser].Profile,
		BannedUsers: append(users[banningUser].BannedUsers, users[request.BannedUser].Username),
	}
	updateUser(users[banningUser])
	return nil
}
