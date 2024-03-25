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
	profReq.Username = r.URL.Query().Get("username")

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
	profile, err1 := rt.db.GetUserProfile(profReq.Username)
	if err1 != nil {
		http.Error(w, err1.Error(), http.StatusBadRequest)
		return
	}
	profResponse := getProfileResponse{Profile: *profile}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err3 := json.NewEncoder(w).Encode(profResponse)
	if err3 != nil {
		http.Error(w, err3.Error(), http.StatusBadRequest)
		return
	}
}
