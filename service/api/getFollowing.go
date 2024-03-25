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

// Get followings
func (rt *_router) getFollowing(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var request followingsRequest
	request.Username = ps.ByName("username")
	//Check if username1 follows username2
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
	print("\n")
	print("Followings: ")
	//print every following
	for _, following := range *following {
		print(following.Username)
		print("\n")
	}
	//Create response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
