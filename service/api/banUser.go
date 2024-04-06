package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type banRequest struct {
	BannedUser string `json:"bannedUser"`
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

	auth, e := rt.db.CheckAuthorization(token)
	if e != nil {
		http.Error(w, e.Error(), http.StatusBadRequest)
		return
	}
	if !auth {
		http.Error(w, "token is invalid", http.StatusBadRequest)
		return
	}

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
}
