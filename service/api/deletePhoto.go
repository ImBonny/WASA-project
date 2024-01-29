package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type deleteRequest struct {
	postId int `json:"postId"`
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

	token := getToken(r.Header.Get("Authorization"))
	if token == 0 {
		http.Error(w, "no token provided", http.StatusBadRequest)
		return
	}

	//TODO: IMPLEMENT SECURITY ONCE I HAVE DB

	var err error
	deleteReq.postId, err = strconv.Atoi(ps.ByName("postId"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	err = rt.db.DeletePhoto(uint64(deleteReq.postId))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	deleteResponse := deleteResponse{Message: "Successfully deleted the post"}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(deleteResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
