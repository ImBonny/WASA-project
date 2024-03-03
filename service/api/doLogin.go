package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"regexp"
)

// LoginRequest represents the request body for user details
type LoginRequest struct {
	Name string `json:"name"`
}

// LoginResponse represents the response body for user details
type LoginResponse struct {
	Identifier uint64 `json:"Identifier"`
}

// doLogin handles the login request
func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var loginReq LoginRequest
	_ = ps
	// Decode the request body into loginReq
	var err error
	if err = json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Validate the name using regular expression pattern
	validName := regexp.MustCompile(`^.*$`)
	if !validName.MatchString(loginReq.Name) || len(loginReq.Name) < 3 || len(loginReq.Name) > 16 {
		http.Error(w, "Invalid username format", http.StatusBadRequest)
		return
	}

	var userId uint64
	token := getToken(r.Header.Get("Authorization"))
	var err0 error
	userId, err0 = rt.db.DoLogin(loginReq.Name, token)
	if err0 != nil {
		http.Error(w, err0.Error(), http.StatusBadRequest)
		return
	}
	_, err1 := rt.db.CheckAuthorization(token)
	if err1 != nil {
		http.Error(w, err1.Error(), http.StatusBadRequest)
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
