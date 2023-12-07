package api

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type changeRequest struct {
	Username    string `json:"username"`
	NewUsername string `json:"newUsername"`
}

type changeResponse struct {
	Message string `json:"message"`
}

// Handler for changing username
// TODO: check how to handle request body
func (rt *_router) setMyUsername(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Create a new change request
	var changeReq changeRequest
	// Decode the request body into changeReq
	if err := json.NewDecoder(r.Body).Decode(&changeReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	changeReq.Username = ps.ByName("Username")
	changeReq.NewUsername = ps.ByName("NewUsername")
	changeToNewUsername(changeReq)
	changeResponse := changeResponse{Message: "Successfully changed username"}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(w).Encode(changeResponse)
	if err != nil {
		return
	}
}

// Change the username of the current user
func changeToNewUsername(request changeRequest) error {
	if _, exists := users[request.NewUsername]; !exists {
		delete(users, request.Username)
		setUsername(request.NewUsername)
		users[request.NewUsername] = getCurrentUser()
		return nil
	}

	return fmt.Errorf("User with username %s already exists", request.NewUsername)
}
