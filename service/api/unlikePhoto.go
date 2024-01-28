package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type unlikeRequest struct {
	PostId uint64 `json:"postId"`
	LikeId uint64 `json:"likeId"`
}
type unlikeResponse struct {
	Message string `json:"message"`
}

// Handler for unliking a post
func (rt *_router) unlikePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Create a new unlike request
	var unlikeReq unlikeRequest
	var err error
	unlikeReq.PostId, err = strconv.ParseUint(ps.ByName("postId"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var user User
	token := getToken(r.Header.Get("Authorization"))
	user.UserId = token

	//TODO: IMPLEMENT SECURITY ONCE I HAVE DB

	unlikeReq.LikeId, err = strconv.ParseUint(ps.ByName("likeId"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	rt.db.UnlikePhoto(unlikeReq.PostId, unlikeReq.LikeId)
	unlikeResponse := unlikeResponse{Message: "Successfully unliked the post"}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(unlikeResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
