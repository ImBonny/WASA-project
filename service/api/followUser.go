package api

import (
	"encoding/json"
	"fmt"
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
func followUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	var request followRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	followThisUser(request.Username, request.Profile)
	response := followResponse{
		Username: request.Username,
		Profile:  request.Profile,
	}
	json.NewEncoder(w).Encode(response)
}

// Follow a user
func followThisUser(username string, profileToFollow Profile) error {
	if _, exists := users[profileToFollow.Username]; exists {
		users[username] = User{
			Username: username,
			Profile: Profile{
				Username:       users[username].Username,
				Posts:          users[username].Profile.Posts,
				NumberOfPhotos: users[username].Profile.NumberOfPhotos,
				Followers:      append(users[username].Profile.Followers, profileToFollow.Username),
				Following:      users[username].Profile.Following,
			},
			BannedUsers: users[username].BannedUsers,
		}
		return nil
	}
	return fmt.Errorf("User with %s username not found", username)
}
