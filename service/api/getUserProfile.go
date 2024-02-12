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
	var err error
	if err = json.NewDecoder(r.Body).Decode(&getProfileRequest{}); err != nil {
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
