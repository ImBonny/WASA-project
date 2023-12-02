package api

import (
	"encoding/json"
	"net/http"
	"time"
)

type LikeRequest struct {
	targetPost int    `json:"postId"`
	likingUser string `json:"username"`
}

type LikeResponse struct {
	// Define the structure of the response if needed
	// This could include the liked post details, confirmation message, etc.
	Message string `json:"message"`
}

// Handler for liking a post
func likePhoto(w http.ResponseWriter, r *http.Request) {
	// Create a new like request
	var likeReq LikeRequest
	// Decode the request body into likeReq
	if err := json.NewDecoder(r.Body).Decode(&likeReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var newLike = Like{
		// Create a new like
		LikeOwner:    likeReq.likingUser,
		CreationTime: time.Now().Format("2006-01-02 15:04:05"),
		LikeId:       int(len(posts[likeReq.targetPost].Likes)),
		// Update the post with the new like
	}
	posts[likeReq.targetPost] = Post{
		PostOwner:    posts[likeReq.targetPost].PostOwner,
		Image:        posts[likeReq.targetPost].Image,
		Comments:     posts[likeReq.targetPost].Comments,
		NComments:    posts[likeReq.targetPost].NComments,
		Likes:        append(posts[likeReq.targetPost].Likes, newLike),
		NLikes:       posts[likeReq.targetPost].NLikes + 1,
		CreationTime: posts[likeReq.targetPost].CreationTime,
		PostId:       posts[likeReq.targetPost].PostId,
	}
	// In this example, assuming a successful like, prepare the response
	likeResponse := LikeResponse{Message: "Successfully liked the post"}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(likeResponse)
}
