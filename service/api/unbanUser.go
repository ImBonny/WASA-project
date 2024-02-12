package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type unbanUserRequest struct {
	BannedUser    string `json:"bannedUser"`
	UnbanningUser string `json:"username"`
}

type unbanUserResponse struct {
	Message string `json:"message"`
}

// Handler for unbanning a user
func (rt *_router) unbanUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Create a new unban request
	var unbanReq unbanUserRequest
	unbanReq.UnbanningUser = ps.ByName("username")
	unbanReq.BannedUser = ps.ByName("bannedUser")
	// Decode the request body into unbanReq
	var err error
	if err = json.NewDecoder(r.Body).Decode(&unbanReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token := getToken(r.Header.Get("Authorization"))

	auth, err := rt.db.CheckAuthorization(token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if !auth {
		http.Error(w, "token is invalid", http.StatusBadRequest)
		return
	}

	bannedId, err := rt.db.GetIdFromUsername(unbanReq.BannedUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = rt.db.UnbanUser(token, bannedId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	unbanResponse := unbanUserResponse{Message: "Successfully unbanned the user"}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(unbanResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
