package api

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type unlikeRequest struct {
	PostId    int    `json:"postId"`
	PostOwner string `json:"username"`
	LikeId    int    `json:"likeId"`
}
type unlikeResponse struct {
	Message string `json:"message"`
}

// Handler for unliking a post
func (rt *_router) unlikePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Create a new unlike request
	var unlikeReq unlikeRequest
	unlikeReq.PostOwner = ps.ByName("username")
	var err error
	unlikeReq.PostId, err = strconv.Atoi(ps.ByName("postId"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	unlikeReq.LikeId, err = strconv.Atoi(ps.ByName("likeId"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Decode the request body into unlikeReq
	if err := json.NewDecoder(r.Body).Decode(&unlikeReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	removeLikeByID(unlikeReq)
	unlikeResponse := unlikeResponse{Message: "Successfully unliked the post"}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(unlikeResponse)
}

// Remove a like from a post by ID
func removeLikeByID(request unlikeRequest) error {
	// Check if the post exists
	post, exists := posts[request.PostId]
	if !exists {
		return fmt.Errorf("Post with ID %d not found", request.PostId)
	}

	// Find the index of the like with the given ID
	likeIndex := -1
	for i, like := range post.Likes {
		if like.LikeId == request.LikeId {
			likeIndex = i
			break
		}
	}

	// If the like put by the specified user is found, remove it
	if likeIndex != -1 {
		post.Likes = append(post.Likes[:likeIndex], post.Likes[likeIndex+1:]...)
		post.NLikes--
		posts[request.PostId] = post // Update the post in the posts map
		return nil
	}

	return fmt.Errorf("Like with id %s not found in post with id %d", request.LikeId, request.PostId)
}
