package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type deleteRequest struct {
	PostId uint64 `json:"postId"`
}

// Handler for deleting a post
func (rt *_router) deletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Create a new delete request
	var deleteReq deleteRequest
	// Decode the request body into deleteReq

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
	var err1 error
	deleteReq.PostId, err1 = strconv.ParseUint(ps.ByName("postId"), 10, 64)
	if err1 != nil {
		http.Error(w, err1.Error(), http.StatusBadRequest)
	}

	err2 := rt.db.DeletePhoto(deleteReq.PostId)
	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusBadRequest)
		return
	}
}
