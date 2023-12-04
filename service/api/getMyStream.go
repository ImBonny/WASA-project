package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type getMyStreamRequest struct {
	Username string `json:"username"`
}

type getMyStreamResponse struct {
	Posts []Post `json:"posts"`
}

// getMyStream returns a list of posts from the user's stream
func (rt *_router) getMyStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Get username from session
	var request getMyStreamRequest
	request.Username = ps.ByName("username")
	// Get posts from user's stream
	posts := getMyStream(request)
	// Create response
	response := getMyStreamResponse{Posts: posts}
	// Send response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// getMyStream returns a list of posts from the user's stream
func getMyStream(request getMyStreamRequest) []Post {
	var stream []Post
	for _, following := range CurrentUser.Profile.Following {
		for _, postId := range users[following].Profile.Posts {
			stream = append([]Post{}, posts[postId])
		}
	} //Sort posts by timestamp
	for i := 0; i < len(stream); i++ {
		for j := i + 1; j < len(stream); j++ {
			if stream[i].CreationTime < stream[j].CreationTime {
				temp := stream[i]
				stream[i] = stream[j]
				stream[j] = temp
			}
		}
	}
	return stream
}
