package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
	"time"
)

type LikeRequest struct {
	targetPost int    `json:"postId"`
	postOwner  string `json:"username"`
}

type LikeResponse struct {
	like Like `json:"like"`
}

// Handler for liking a post
func (rt *_router) likePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Create a new like request
	var likeReq LikeRequest
	likeReq.postOwner = ps.ByName("username")
	var err error
	likeReq.targetPost, err = strconv.Atoi(ps.ByName("postId"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var user User
	token := getToken(r.Header.Get("Authorization"))
	user.UserId = token

	//TODO: IMPLEMENT SECURITY ONCE I HAVE DB

	// Decode the request body into likeReq
	if err = json.NewDecoder(r.Body).Decode(&likeReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	myUsername := getCurrentUser().Username
	var newLike = Like{
		// Create a new like
		LikeOwner:    myUsername,
		CreationTime: time.Now().Format("2006-01-02 15:04:05"),
		LikeId:       len(posts[likeReq.targetPost].Likes) + 1,
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
	likeResponse := LikeResponse{like: newLike}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(likeResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
