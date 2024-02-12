package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type uncommentRequest struct {
	CommentId uint64 `json:"commentId"`
}

// Handler for uncommenting a post
func (rt *_router) uncommentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Create a new uncomment request
	var uncommentReq uncommentRequest
	var err error

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

	uncommentReq.CommentId, err = strconv.ParseUint(ps.ByName("commentId"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = rt.db.UncommentPhoto(uncommentReq.CommentId)
}
