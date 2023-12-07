package api

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type deleteRequest struct {
	postId   int    `json:"postId"`
	username string `json:"username"`
}

type deleteResponse struct {
	Message string `json:"message"`
}

// Handler for deleting a post
func (rt *_router) deletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Create a new delete request
	var deleteReq deleteRequest
	// Decode the request body into deleteReq
	if err := json.NewDecoder(r.Body).Decode(&deleteReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	deleteReq.username = ps.ByName("username")
	var err error
	deleteReq.postId, err = strconv.Atoi(ps.ByName("postId"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	removePostByID(deleteReq.postId, deleteReq.username)
	deleteResponse := deleteResponse{Message: "Successfully deleted the post"}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(deleteResponse)
	if err != nil {
		return
	}
}

func removePostByID(postID int, deleteUsername string) error {
	// Find the index of the post with the given ID
	postIndex := 1
	for i, post := range posts {
		if post.PostOwner == deleteUsername {
			postIndex = i
			break
		}
	}
	// If the post put by the specified user is found, remove it
	if postIndex != -1 {
		delete(posts, postIndex)
		deletePost(postIndex)
		return nil
	}

	return fmt.Errorf("Post with owner %s not found in post %d", deleteUsername, postID)
}
