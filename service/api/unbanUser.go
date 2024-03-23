package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type unbanUserRequest struct {
	BannedUser string `json:"bannedUser"`
}

type unbanUserResponse struct {
	Message string `json:"message"`
}

// Handler for unbanning a user
func (rt *_router) unbanUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Create a new unban request
	var unbanReq unbanUserRequest
	// Decode the request body into unbanReq
	unbanReq.BannedUser = r.URL.Query().Get("BannedUser")

	token := getToken(r.Header.Get("Authorization"))

	auth, err1 := rt.db.CheckAuthorization(token)
	if err1 != nil {
		http.Error(w, err1.Error(), http.StatusBadRequest)
		return
	}
	if !auth {
		http.Error(w, "token is invalid", http.StatusBadRequest)
		return
	}

	bannedId, err2 := rt.db.GetIdFromUsername(unbanReq.BannedUser)
	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusBadRequest)
		return
	}
	err3 := rt.db.UnbanUser(token, bannedId)
	if err3 != nil {
		http.Error(w, err3.Error(), http.StatusBadRequest)
		return
	}
	unbanResponse := unbanUserResponse{Message: "Successfully unbanned the user"}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err4 := json.NewEncoder(w).Encode(unbanResponse)
	if err4 != nil {
		http.Error(w, err4.Error(), http.StatusBadRequest)
		return
	}
}
