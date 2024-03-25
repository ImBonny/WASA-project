package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type unlikeRequest struct {
	TargetPost uint64 `json:"TargetPost"`
	LikeOwner  uint64 `json:"LikeOwner"`
}

// Handler for unliking a post
func (rt *_router) unlikePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Create a new unlike request
	var unlikeReq unlikeRequest
	var err error

	err = json.NewDecoder(r.Body).Decode(&unlikeReq)

	if err != nil {

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token := getToken(r.Header.Get("Authorization"))

	auth, err1 := rt.db.CheckAuthorization(token)
	if err1 != nil {
		http.Error(w, err1.Error(), http.StatusBadRequest)
		return
	}
	if !auth {
		http.Error(w, "token is invalid", http.StatusBadRequest)
		return
	}

	err2 := rt.db.UnlikePhoto(unlikeReq.LikeOwner, unlikeReq.TargetPost)
	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusBadRequest)
		return
	}

}
