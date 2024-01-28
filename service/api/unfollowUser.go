package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type unfollowRequest struct {
	Username string  `json:"username"`
	Profile  Profile `json:"profile"`
}

type unfollowResponse struct {
	Username string `json:"username"`
}

// Unfollow a user
func (rt *_router) unfollowUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("content-type", "application/json")
	var user User
	token := getToken(r.Header.Get("Authorization"))
	user.UserId = token

	//TODO: IMPLEMENT SECURITY ONCE I HAVE DB

	var request unfollowRequest
	request.Username = ps.ByName("username")
	request.Profile = users[ps.ByName("username")].Profile
	unfollowId, err := rt.db.GetIdFromUsername(request.Username)
	err = rt.db.UnfollowUser(token, unfollowId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := unfollowResponse{
		Username: request.Username,
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
