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
	var err error
	if err = json.NewDecoder(r.Body).Decode(&deleteReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

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
	var err1 error
	deleteReq.postId, err1 = strconv.Atoi(ps.ByName("postId"))
	if err1 != nil {
		http.Error(w, err1.Error(), http.StatusBadRequest)
	}
	err2 := rt.db.DeletePhoto(uint64(deleteReq.postId))
	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusBadRequest)
		return
	}
	deleteResponse := deleteResponse{Message: "Successfully deleted the post"}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err3 := json.NewEncoder(w).Encode(deleteResponse)
	if err3 != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
