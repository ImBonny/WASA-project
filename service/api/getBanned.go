package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type isBannedRequest struct {
	UsernameToCheck string `json:"UsernameToCheck"`
	UsernameBanning string `json:"UsernameBanning"`
}

type isBannedResponse struct {
	Banned bool `json:"Banned"`
}

// Handler for checking if a user is banned
func (rt *_router) getBanned(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Create a new isBanned request
	var isBannedReq isBannedRequest
	// Decode the request body into isBannedReq
	isBannedReq.UsernameToCheck = r.URL.Query().Get("UsernameToCheck")

	isBannedReq.UsernameBanning = r.URL.Query().Get("UsernameBanning")

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
	id2, err3 := rt.db.GetIdFromUsername(isBannedReq.UsernameToCheck)
	if err3 != nil {
		http.Error(w, err3.Error(), http.StatusBadRequest)
		return
	}
	id1, err2 := rt.db.GetIdFromUsername(isBannedReq.UsernameBanning)
	if err2 != nil {
		http.Error(w, err2.Error(), http.StatusBadRequest)
		return

	}
	print("\nid1: ", id1, " id2: ", id2, "\n")
	banned, err4 := rt.db.IsBanned(id1, id2)
	if err4 != nil {
		http.Error(w, err4.Error(), http.StatusBadRequest)
		return
	}
	isBannedResponse := isBannedResponse{Banned: banned}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err5 := json.NewEncoder(w).Encode(isBannedResponse)
	if err5 != nil {
		http.Error(w, err5.Error(), http.StatusBadRequest)
		return
	}
}
