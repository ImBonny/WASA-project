package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type commentRequest struct {
	postId      int    `json:"postId"`
	username    string `json:"username"`
	commentText string `json:"commentText"`
}

type commentResponse struct {
	Message string `json:"message"`
}

// Handler for commenting on a post
func commentPhoto(w http.ResponseWriter, r *http.Request) {
	// Create a new comment request
	var commentReq commentRequest
	// Decode the request body into commentReq
	if err := json.NewDecoder(r.Body).Decode(&commentReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	addComment(commentReq.postId, commentReq.username, commentReq.commentText)
	commentResponse := commentResponse{Message: "Successfully commented the post"}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(commentResponse)
}

func addComment(postID int, commentOwner string, commentText string) error {
	// Check if the post exists
	_, exists := posts[postID]
	if !exists {
		return fmt.Errorf("Post with ID %d not found", postID)
	}

	// Find the index of the post with the given ID
	postIndex := -1
	for i, post := range posts {
		if post.PostOwner == commentOwner {
			postIndex = i
			break
		}
	}

	// If the post put by the specified user is found, add the comment
	if postIndex != -1 {
		var newComment = Comment{
			CommentOwner: commentOwner,
			CommentText:  commentText,
			CreationTime: time.Now().Format("2006-01-02 15:04:05"),
			CommentId:    int(len(posts[postID].Comments)),
		}
		posts[postID] = Post{
			PostOwner:    posts[postID].PostOwner,
			Image:        posts[postID].Image,
			Comments:     append(posts[postID].Comments, newComment),
			NComments:    posts[postID].NComments + 1,
			Likes:        posts[postID].Likes,
			NLikes:       posts[postID].NLikes,
			CreationTime: posts[postID].CreationTime,
			PostId:       posts[postID].PostId,
		}
		return nil
	}

	return fmt.Errorf("Post with owner %s not found in post %d", commentOwner, postID)
}
