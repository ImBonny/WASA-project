package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type followRequest struct {
	Username string `json:"username"`
}

// Follow a user
func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	var request followRequest
	request.Username = ps.ByName("username")

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

	followId, err1 := rt.db.GetIdFromUsername(request.Username)
	if err1 != nil {
		http.Error(w, err1.Error(), http.StatusBadRequest)
		return
	}
	err2 := rt.db.FollowUser(token, followId)
	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusBadRequest)
		return
	}
}
