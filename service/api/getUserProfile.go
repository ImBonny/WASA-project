package api

import (
	"encoding/json"
	"github.com/ImBonny/WASA-project.git/service/database"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type getProfileRequest struct {
	Username string `json:"username"`
}

type getProfileResponse struct {
	Profile database.Database_profile `json:"profile"`
}

func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var profReq getProfileRequest
	// Decode the request body into banReq
	if err := json.NewDecoder(r.Body).Decode(&getProfileRequest{}); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token := getToken(r.Header.Get("Authorization"))
	if token == 0 {
		http.Error(w, "no token provided", http.StatusBadRequest)
		return
	}
	//TODO: IMPLEMENT SECURITY ONCE I HAVE DB

	profReq.Username = ps.ByName("username")
	profile, err := rt.db.GetUserProfile(profReq.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	profResponse := getProfileResponse{Profile: *profile}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(profResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
