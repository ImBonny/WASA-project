package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type unlikeRequest struct {
	PostId       int    `json:"postId"`
	unlikingUser string `json:"unlikingUser"`
}
type unlikeResponse struct {
	Message string `json:"message"`
}

// Handler for unliking a post
func unlikePhoto(w http.ResponseWriter, r *http.Request) {
	// Create a new unlike request
	var unlikeReq unlikeRequest
	// Decode the request body into unlikeReq
	if err := json.NewDecoder(r.Body).Decode(&unlikeReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	removeLikeByID(unlikeReq.PostId, unlikeReq.unlikingUser)
	unlikeResponse := unlikeResponse{Message: "Successfully unliked the post"}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(unlikeResponse)
}

// Remove a like from a post by ID
func removeLikeByID(postID int, unlikeUser string) error {
	// Check if the post exists
	post, exists := posts[postID]
	if !exists {
		return fmt.Errorf("Post with ID %d not found", postID)
	}

	// Find the index of the like with the given ID
	likeIndex := -1
	for i, like := range post.Likes {
		if like.LikeOwner == unlikeUser {
			likeIndex = i
			break
		}
	}

	// If the like put by the specified user is found, remove it
	if likeIndex != -1 {
		post.Likes = append(post.Likes[:likeIndex], post.Likes[likeIndex+1:]...)
		post.NLikes--
		posts[postID] = post // Update the post in the posts map
		return nil
	}

	return fmt.Errorf("Like with owner %s not found in post %d", unlikeUser, postID)
}
