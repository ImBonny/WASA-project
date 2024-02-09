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
	message string `json:"message"`
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

	err = rt.db.LikePhoto(likeReq.targetPost, token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	likeResponse := LikeResponse{message: "Post liked"}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(likeResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
