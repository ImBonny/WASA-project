package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type followsRequest struct {
	Username1 string `json:"username"`
	Username2 string `json:"username"`
}

type followsResponse struct {
	Result bool `json:"result"`
}

func (rt *_router) getFollows(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var request followsRequest
	request.Username1 = ps.ByName("username")
	request.Username2 = ps.ByName("usernameFollowed")

	userId, err := rt.db.GetIdFromUsername(request.Username1)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	following, err1 := rt.db.GetFollowing(userId)
	if err1 != nil {
		http.Error(w, err1.Error(), http.StatusBadRequest)
		return
	}
	follows := false
	for _, user := range *following {
		if user.Username == request.Username2 {
			follows = true
			break
		}
	}
	response := followsResponse{
		Result: follows,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
