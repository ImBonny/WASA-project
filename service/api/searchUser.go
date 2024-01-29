package api

import (
	"encoding/json"
	"github.com/ImBonny/WASA-project.git/service/database"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type searchUserRequest struct {
	Username string `json:"username"`
}

type searchUserResponse struct {
	User database.Database_user `json:"user"`
}

// searchUser returns a list of users that match the search query
func (rt *_router) searchUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var request searchUserRequest
	var response searchUserResponse
	var err error

	token := getToken(r.Header.Get("Authorization"))
	if token == 0 {
		http.Error(w, "no token provided", http.StatusBadRequest)
		return
	}

	//TODO: IMPLEMENT SECURITY ONCE I HAVE DB

	request.Username = r.URL.Query().Get("username")
	user, err := rt.db.SearchUser(request.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response.User = user

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
