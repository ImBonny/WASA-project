package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type getProfileRequest struct {
	Username string `json:"username"`
}

type getProfileResponse struct {
	Profile Profile `json:"profile"`
}

func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var profReq getProfileRequest
	// Decode the request body into banReq
	if err := json.NewDecoder(r.Body).Decode(&getProfileRequest{}); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	profReq.Username = ps.ByName("username")
	profile := getProfile(profReq.Username)
	profResponse := getProfileResponse{Profile: profile}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(profResponse)
}

func getProfile(username string) Profile {
	return users[username].Profile
}
