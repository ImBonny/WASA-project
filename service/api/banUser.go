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
	if err := json.NewDecoder(r.Body).Decode(&banReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token := getToken(r.Header.Get("Authorization"))

	//TODO: IMPLEMENT SECURITY ONCE I HAVE DB

	banReq.BannedUser = ps.ByName("bannedUser")

	bannedId, err := rt.db.GetIdFromUsername(banReq.BannedUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = rt.db.BanUser(token, bannedId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	banResponse := banResponse{Username: banReq.BannedUser}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(banResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
