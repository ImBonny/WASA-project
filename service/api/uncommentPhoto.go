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
	var err1 error
	uncommentReq.CommentId, err1 = strconv.ParseUint(ps.ByName("commentId"), 10, 64)
	if err1 != nil {
		http.Error(w, err1.Error(), http.StatusBadRequest)
		return
	}
	err2 := rt.db.UncommentPhoto(uncommentReq.CommentId)
	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusBadRequest)
		return
	}
}
