package api

import (
	"encoding/json"
	"github.com/ImBonny/WASA-project.git/service/database"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type followingsRequest struct {
	Username string `json:"Username"`
}

type followingsResponse struct {
	Followings []database.Database_user `json:"Followings"`
}

func (rt *_router) getFollowing(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var request followingsRequest
	request.Username = ps.ByName("username")

	userId, err := rt.db.GetIdFromUsername(request.Username)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	following, err1 := rt.db.GetFollowing(userId)
	if err1 != nil {
		http.Error(w, err1.Error(), http.StatusBadRequest)
		return
	}
	response := followingsResponse{
		Followings: *following,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
