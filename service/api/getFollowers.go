package api

import (
	"encoding/json"
	"github.com/ImBonny/WASA-project.git/service/database"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type followersRequest struct {
	Username string `json:"Username"`
}

type followersResponse struct {
	Followers []database.Database_user `json:"Followers"`
}

func (rt *_router) getFollowers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var request followersRequest
	request.Username = ps.ByName("username")
	//Get followers
	userId, err := rt.db.GetIdFromUsername(request.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	followers, err1 := rt.db.GetFollowers(userId)
	print("Followers: ")
	//print every follower
	for _, follower := range *followers {
		print(follower.Username)
		print("/n")
	}
	if err1 != nil {
		http.Error(w, err1.Error(), http.StatusBadRequest)
		return
	}
	response := followersResponse{
		Followers: *followers,
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
