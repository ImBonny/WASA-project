package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type changeRequest struct {
	NewUsername string `json:"newUsername"`
}

type changeResponse struct {
	Message string `json:"message"`
}

// Handler for changing username
func (rt *_router) setMyUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Create a new change request
	var changeReq changeRequest
	// Decode the request body into changeReq
	if err := json.NewDecoder(r.Body).Decode(&changeReq); err != nil {
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

	changeReq.NewUsername = ps.ByName("NewUsername")
	err := rt.db.SetMyUsername(token, changeReq.NewUsername)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	changeResponse := changeResponse{Message: "Successfully changed username"}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(changeResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
