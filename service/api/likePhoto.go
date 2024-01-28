package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type LikeRequest struct {
	targetPost uint64 `json:"postId"`
	postOwner  string `json:"username"`
}

type LikeResponse struct {
	like Like `json:"like"`
}

// Handler for liking a post
func (rt *_router) likePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Create a new like request
	var likeReq LikeRequest
	likeReq.postOwner = ps.ByName("username")
	var err error
	likeReq.targetPost, err = strconv.ParseUint(ps.ByName("postId"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var user User
	token := getToken(r.Header.Get("Authorization"))
	user.UserId = token

	//TODO: IMPLEMENT SECURITY ONCE I HAVE DB
	rt.db.LikePhoto(likeReq.targetPost, user.UserId)

	// Create a new like
	var newLike Like
	newLike.PostId = likeReq.targetPost
	newLike.LikeOwner = user.UserId

	likeResponse := LikeResponse{like: newLike}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(likeResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
