package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type commentRequest struct {
	PostId      uint64 `json:"postId"`
	CommentText string `json:"commentText"`
}

type commentResponse struct {
	commentId uint64 `json:"comment"`
}

// Handler for commenting on a post
func (rt *_router) commentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Create a new comment request
	var commentReq commentRequest
	// Decode the request body into commentReq
	if err := json.NewDecoder(r.Body).Decode(&commentReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token := getToken(r.Header.Get("Authorization"))

	//TODO: IMPLEMENT SECURITY ONCE I HAVE DB

	var err error
	commentReq.PostId, err = strconv.ParseUint(ps.ByName("postId"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	body, err := json.Marshal(commentReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	commentReq.CommentText = string(body)
	// Find the post put by the specified user

	commentId, _ := rt.db.CommentPhoto(token, commentReq.PostId, commentReq.CommentText)

	commentResponse := commentResponse{commentId: commentId}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(commentResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
