package api

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type followRequest struct {
	Username string  `json:"username"`
	Profile  Profile `json:"profile"`
}

type followResponse struct {
	Username string  `json:"username"`
	Profile  Profile `json:"profile"`
}

// Follow a user
func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("content-type", "application/json")
	myUsername := getCurrentUser().Username

	var request followRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	request.Username = ps.ByName("username")
	request.Profile = users[request.Username].Profile
	err = followThisUser(myUsername, request)
	if err != nil {
		return
	}
	response := followResponse{
		Username: request.Username,
		Profile:  request.Profile,
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}
}

// Follow a user
func followThisUser(myUsername string, request followRequest) error {
	if _, exists := users[request.Username]; exists {
		users[myUsername] = User{
			Username: myUsername,
			Profile: Profile{
				Username:       users[myUsername].Username,
				Posts:          users[myUsername].Profile.Posts,
				NumberOfPhotos: users[myUsername].Profile.NumberOfPhotos,
				Followers:      append(users[myUsername].Profile.Followers, request.Username),
				Following:      users[myUsername].Profile.Following,
			},
			BannedUsers: users[myUsername].BannedUsers,
		}
		return nil
		updateUser(users[myUsername])
	}
	return fmt.Errorf("User with %s username not found", request.Username)
}
