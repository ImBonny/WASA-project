package api

import (
	"encoding/json"
	"github.com/ImBonny/WASA-project.git/service/database"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type getMyStreamRequest struct {
	Username string `json:"username"`
}

type getMyStreamResponse struct {
	Posts []database.Database_photo `json:"posts"`
}

// getMyStream returns a list of posts from the user's stream
func (rt *_router) getMyStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Get username from session
	var request getMyStreamRequest
	request.Username = ps.ByName("username")

	token := getToken(r.Header.Get("Authorization"))

	//TODO: IMPLEMENT SECURITY ONCE I HAVE DB

	myStream, err := rt.db.GetMyStream(token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Create response
	response := getMyStreamResponse{Posts: *myStream}
	// Send response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}
