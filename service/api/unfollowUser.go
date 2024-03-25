package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type unfollowRequest struct {
	Username string `json:"username"`
}

// Unfollow a user
func (rt *_router) unfollowUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("content-type", "application/json")
	token := getToken(r.Header.Get("Authorization"))
	auth, err := rt.db.CheckAuthorization(token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if !auth {
		http.Error(w, "token is invalid", http.StatusBadRequest)
		return
	}

	var request unfollowRequest
	request.Username = ps.ByName("username")
	unfollowId, err1 := rt.db.GetIdFromUsername(request.Username)
	if err1 != nil {
		http.Error(w, err1.Error(), http.StatusBadRequest)
		return
	}
	err2 := rt.db.UnfollowUser(token, unfollowId)
	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusBadRequest)
		return
	}
}
