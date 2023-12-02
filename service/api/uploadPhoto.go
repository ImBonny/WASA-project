package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"time"
)

var posts = map[int]Post{}

// Upload a photo to the current user's profile
func uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("content-type", "application/json")
	currentUsername := getCurrentUser().Username
	uploadedImage := r.URL.Query().Get("image")
	id := int(len(posts))
	posts[id] = Post{
		PostOwner:    currentUsername,
		Image:        uploadedImage,
		Comments:     []Comment{},
		NComments:    0,
		Likes:        []Like{},
		NLikes:       0,
		CreationTime: time.Now().Format("2006-01-02 15:04:05"),
		PostId:       id,
	}
	CurrentUser.Profile.Posts = append(getCurrentUser().Profile.Posts, id)
	CurrentUser.Profile.NumberOfPhotos++
	users[currentUsername] = CurrentUser
	json.NewEncoder(w).Encode(id)
}
