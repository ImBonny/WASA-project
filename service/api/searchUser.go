package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type searchUserRequest struct {
	Username string `json:"username"`
}

type searchUserResponse struct {
	Username string `json:"username"`
}

// searchUser returns a list of users that match the search query
func (rt *_router) searchUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var request searchUserRequest
	var response searchUserResponse
	token := getToken(r.Header.Get("Authorization"))
	auth, e := rt.db.CheckAuthorization(token)
	if e != nil {
		http.Error(w, e.Error(), http.StatusBadRequest)
		return
	}
	if !auth {
		http.Error(w, "token is invalid", http.StatusBadRequest)
		return
	}

	request.Username = r.URL.Query().Get("username")
	user, err := rt.db.SearchUser(request.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response.Username = user

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err1 := json.NewEncoder(w).Encode(response)
	if err1 != nil {
		http.Error(w, err1.Error(), http.StatusBadRequest)
		return
	}
}
