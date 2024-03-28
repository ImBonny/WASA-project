package api

import (
	"encoding/json"
	"github.com/ImBonny/WASA-project.git/service/database"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type getMyStreamRequest struct {
	Username string `json:"Username"`
}

type getMyStreamResponse struct {
	Posts []database.Database_photo `json:"Posts"`
}

// getMyStream returns a list of posts from the user's stream
func (rt *_router) getMyStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Get username from session
	var request getMyStreamRequest
	request.Username = ps.ByName("username")

	// Get token from header

	token := getToken(r.Header.Get("Authorization"))

	auth, e := rt.db.CheckAuthorization(token)
	if e != nil {
		http.Error(w, e.Error(), http.StatusBadRequest)
		return
	}
	if !auth {
		http.Error(w, "token is invalid", http.StatusBadRequest)
		return
	}

	myStream, err := rt.db.GetMyStream(token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	myPosts, err1 := rt.db.GetImages(myStream)
	if err1 != nil {
		http.Error(w, err1.Error(), http.StatusBadRequest)
		return

	}
	// Create response
	response := getMyStreamResponse{Posts: *myPosts}
	// Send response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err2 := json.NewEncoder(w).Encode(response)
	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusBadRequest)
		return
	}

}
