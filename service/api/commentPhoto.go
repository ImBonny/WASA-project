package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
	"time"
)

type commentRequest struct {
	PostId      int    `json:"postId"`
	Username    string `json:"username"`
	CommentText string `json:"commentText"`
}

type commentResponse struct {
	Comment Comment `json:"comment"`
}

// Handler for commenting on a post
func (rt *_router) commentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Create a new comment request
	var commentReq commentRequest
	// Decode the request body into commentReq
	if err := json.NewDecoder(r.Body).Decode(&commentReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var user User
	token := getToken(r.Header.Get("Authorization"))
	user.UserId = token

	//TODO: IMPLEMENT SECURITY ONCE I HAVE DB

	commentReq.Username = getCurrentUser().Username
	var err error
	commentReq.PostId, err = strconv.Atoi(ps.ByName("postId"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	body, err := json.Marshal(commentReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	commentReq.CommentText = string(body)
	// Find the post put by the specified user
	postIndex := -1

	for i, post := range posts {
		if post.PostId == commentReq.PostId {
			postIndex = i
			break
		}
	}
	myUsername := getCurrentUser().Username
	// If the post put by the specified user is found, add the comment
	var newComment Comment
	if postIndex != -1 {
		newComment = Comment{
			CommentOwner: myUsername,
			CommentText:  commentReq.CommentText,
			CreationTime: time.Now().Format("2006-01-02 15:04:05"),
			CommentId:    len(posts[commentReq.PostId].Comments),
		}
		posts[commentReq.PostId] = Post{
			PostOwner:    posts[commentReq.PostId].PostOwner,
			Image:        posts[commentReq.PostId].Image,
			Comments:     append(posts[commentReq.PostId].Comments, newComment),
			NComments:    posts[commentReq.PostId].NComments + 1,
			Likes:        posts[commentReq.PostId].Likes,
			NLikes:       posts[commentReq.PostId].NLikes,
			CreationTime: posts[commentReq.PostId].CreationTime,
			PostId:       posts[commentReq.PostId].PostId,
		}
	}
	commentResponse := commentResponse{Comment: newComment}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(commentResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
