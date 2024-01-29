package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type unfollowRequest struct {
	Username string `json:"username"`
}

type unfollowResponse struct {
	Username string `json:"username"`
}

// Unfollow a user
func (rt *_router) unfollowUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("content-type", "application/json")
	token := getToken(r.Header.Get("Authorization"))
	if token == 0 {
		http.Error(w, "no token provided", http.StatusBadRequest)
		return
	}

	//TODO: IMPLEMENT SECURITY ONCE I HAVE DB

	var request unfollowRequest
	request.Username = ps.ByName("username")
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
