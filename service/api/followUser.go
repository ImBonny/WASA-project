package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type followRequest struct {
	Username string `json:"username"`
}

type followResponse struct {
	Username string `json:"username"`
}

// Follow a user
func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("content-type", "application/json")

	var request followRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

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

	request.Username = ps.ByName("username")
	followId, err := rt.db.GetIdFromUsername(request.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = rt.db.FollowUser(token, followId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response := followResponse{
		Username: request.Username,
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
