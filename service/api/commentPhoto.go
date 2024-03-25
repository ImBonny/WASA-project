package api

import (
	"bytes"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
	"strconv"
)

type commentRequest struct {
	CommentText string `json:"CommentText"`
}

type commentResponse struct {
	CommentId uint64 `json:"CommentId"`
}

// Handler for commenting on a post
func (rt *_router) commentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Create a new comment request
	var commentReq commentRequest
	// Decode the request body into commentReq
	body, _ := ioutil.ReadAll(r.Body)

	r.Body = ioutil.NopCloser(bytes.NewBuffer(body))

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
	PostId, err0 := strconv.ParseUint(ps.ByName("postId"), 10, 64)

	if err0 != nil {
		http.Error(w, err0.Error(), http.StatusBadRequest)
	}
	if err := json.NewDecoder(r.Body).Decode(&commentReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	commentId, _ := rt.db.CommentPhoto(token, PostId, commentReq.CommentText)

	// Find the post put by the specified user

	commentRes := commentResponse{CommentId: commentId}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err2 := json.NewEncoder(w).Encode(commentRes)
	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusBadRequest)
		return
	}
}
