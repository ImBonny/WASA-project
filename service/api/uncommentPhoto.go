package api

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type uncommentRequest struct {
	PostId    int    `json:"postId"`
	CommentId int    `json:"commentId"`
	PostOwner string `json:"username"`
}

type uncommentResponse struct {
}

// Handler for uncommenting a post
func (rt *_router) uncommentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Create a new uncomment request
	var uncommentReq uncommentRequest
	uncommentReq.PostOwner = ps.ByName("username")
	var err error
	uncommentReq.PostId, err = strconv.Atoi(ps.ByName("postId"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	uncommentReq.CommentId, err = strconv.Atoi(ps.ByName("commentId"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Decode the request body into uncommentReq
	if err := json.NewDecoder(r.Body).Decode(&uncommentReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = removeCommentByID(uncommentReq.PostId, uncommentReq.CommentId)
	if err != nil {
		return
	}
	var uncommentResponse uncommentResponse
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(uncommentResponse)
}

// Remove a comment from a post by ID
func removeCommentByID(postID int, commentID int) error {
	// Check if the post exists
	post, exists := posts[postID]
	if !exists {
		return fmt.Errorf("Post with ID %d not found", postID)
	}

	// Find the index of the comment with the given ID
	commentIndex := -1
	for i, comment := range post.Comments {
		if comment.CommentId == commentID {
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

	return fmt.Errorf("Comment with id %d not found in post %d", commentID, postID)
}
