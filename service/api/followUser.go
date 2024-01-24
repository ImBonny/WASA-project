package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type followRequest struct {
	Username string  `json:"username"`
	Profile  Profile `json:"profile"`
}

type followResponse struct {
	Username string  `json:"username"`
	Profile  Profile `json:"profile"`
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

	var user User
	token := getToken(r.Header.Get("Authorization"))
	user.UserId = token

	//TODO: IMPLEMENT SECURITY ONCE I HAVE DB

	request.Username = ps.ByName("username")
	request.Profile = users[request.Username].Profile
	followId, err := rt.db.GetIdFromUsername(request.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = rt.db.FollowUser(int(token), followId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response := followResponse{
		Username: request.Username,
		Profile:  request.Profile,
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
