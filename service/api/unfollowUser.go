package api

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type unfollowRequest struct {
	Username string  `json:"username"`
	Profile  Profile `json:"profile"`
}

type unfollowResponse struct {
	Username string `json:"username"`
}

// Unfollow a user
func (rt *_router) unfollowUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("content-type", "application/json")

	var request unfollowRequest
	request.Username = ps.ByName("username")
	request.Profile = users[ps.ByName("username")].Profile
	err := unfollowThisUser(getCurrentUser().Username, request)
	if err != nil {
		return
	}
	response := unfollowResponse{
		Username: request.Username,
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}
}

// Unfollow a user
func unfollowThisUser(myUsername string, request unfollowRequest) error {
	if _, exists := users[myUsername]; exists {
		users[myUsername] = User{
			Username: myUsername,
			Profile: Profile{
				Username:       users[myUsername].Username,
				Posts:          users[myUsername].Profile.Posts,
				NumberOfPhotos: users[myUsername].Profile.NumberOfPhotos,
				Followers:      users[myUsername].Profile.Followers,
				Following:      removeFollowing(users[myUsername].Profile.Following, request.Username),
			},
			BannedUsers: users[myUsername].BannedUsers,
		}
		updateUser(users[myUsername])
		return nil
	}
	return fmt.Errorf("User with %s username not found", request.Username)
}

func removeFollowing(following []string, username string) []string {
	for i := range following {
		if following[i] == username {
			following = append(following[:i], following[i+1:]...)
			break
		}
	}
	return following
}
