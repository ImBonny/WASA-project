package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type uncommentRequest struct {
	PostId           int    `json:"postId"`
	CommentId        int    `json:"commentId"`
	uncommentingUser string `json:"uncommentingUser"`
}

type uncommentResponse struct {
	Message string `json:"message"`
}

// Handler for uncommenting a post
func uncommentPhoto(w http.ResponseWriter, r *http.Request) {
	// Create a new uncomment request
	var uncommentReq uncommentRequest
	// Decode the request body into uncommentReq
	if err := json.NewDecoder(r.Body).Decode(&uncommentReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	removeCommentByID(uncommentReq.PostId, uncommentReq.CommentId, uncommentReq.uncommentingUser)
	uncommentResponse := uncommentResponse{Message: "Successfully uncommented the post"}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(uncommentResponse)
}

// Remove a comment from a post by ID
func removeCommentByID(postID int, commentID int, uncommentUser string) error {
	// Check if the post exists
	post, exists := posts[postID]
	if !exists {
		return fmt.Errorf("Post with ID %d not found", postID)
	}

	// Find the index of the comment with the given ID
	commentIndex := -1
	for i, comment := range post.Comments {
		if comment.CommentOwner == uncommentUser {
			commentIndex = i
			break
		}
	}

	// If the comment put by the specified user is found, remove it
	if commentIndex != -1 {
		post.Comments = append(post.Comments[:commentIndex], post.Comments[commentIndex+1:]...)
		post.NComments--
		posts[postID] = post // Update the post in the posts map
		return nil
	}

	return fmt.Errorf("Comment with owner %s not found in post %d", uncommentUser, postID)
}
