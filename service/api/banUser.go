package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type banRequest struct {
	BannedUser string `json:"bannedUser"`
}

type banResponse struct {
	Username string `json:"username"`
}

func (rt *_router) banUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Create a new ban request
	var banReq banRequest
	// Decode the request body into banReq
	var err error
	if err = json.NewDecoder(r.Body).Decode(&banReq); err != nil {
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

	banReq.BannedUser = ps.ByName("bannedUser")

	bannedId, err0 := rt.db.GetIdFromUsername(banReq.BannedUser)
	if err0 != nil {
		http.Error(w, err0.Error(), http.StatusBadRequest)
		return
	}
	err1 := rt.db.BanUser(token, bannedId)
	if err1 != nil {
		http.Error(w, err1.Error(), http.StatusBadRequest)
		return
	}
	banResponse := banResponse{Username: banReq.BannedUser}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err2 := json.NewEncoder(w).Encode(banResponse)
	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusBadRequest)
		return
	}
}
