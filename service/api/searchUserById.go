package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type searchUserByIdRequest struct {
	UserId uint64 `json:"UserId"`
}

type searchUserByIdResponse struct {
	Username string `json:"Username"`
}

// searchUserById searches for a user by their user id
func (rt *_router) searchUserById(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Get the user id
	var request searchUserByIdRequest
	var err error
	request.UserId, err = strconv.ParseUint(r.URL.Query().Get("UserId"), 10, 64)
	// Get the user

	if err != nil {
		http.Error(w, "Invalid user id", http.StatusBadRequest)
		return
	}
	user, err1 := rt.db.GetUsernameFromId(request.UserId)
	if err1 != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	var response searchUserByIdResponse
	response.Username = user
	// Send the user
	w.Header().Set("Content-Type", "application/json")
	err2 := json.NewEncoder(w).Encode(response)
	if err2 != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

}
