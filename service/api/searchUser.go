package api

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type searchUserRequest struct {
	Username string `json:"username"`
}

type searchUserResponse struct {
	User User `json:"username"`
}

// searchUser returns a list of users that match the search query
func (rt *_router) searchUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	user := r.URL.Query().Get("username")
	var request searchUserRequest
	request.Username = user
	var response searchUserResponse
	var err error
	response.User, err = searchUser(request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// searchUser returns a list of users that match the search query
func searchUser(request searchUserRequest) (User, error) {
	if user, exists := users[request.Username]; exists {
		return user, nil
	}
	return User{}, fmt.Errorf("User '%s' not found", request.Username)
}
