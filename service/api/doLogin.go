package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// LoginRequest represents the request body for user details
type LoginRequest struct {
	Username string `json:"username"`
}

// LoginResponse represents the response body for user details
type LoginResponse struct {
	Identifier uint64 `json:"Identifier"`
}

// doLogin handles the login request
func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var loginReq LoginRequest
	// Decode the request body into loginReq
	if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userId, err0 := rt.db.DoLogin(loginReq.Username)
	if err0 != nil {
		http.Error(w, err0.Error(), http.StatusBadRequest)
		return
	}

	// Create the response body
	response := LoginResponse{Identifier: userId}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err2 := json.NewEncoder(w).Encode(response)
	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusBadRequest)
		return
	}
}
