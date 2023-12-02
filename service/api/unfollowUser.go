package api

import (
	"encoding/json"
	"fmt"
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
func unfollowUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	var request unfollowRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	unfollowThisUser(request.Username, request.Profile)
	response := unfollowResponse{
		Username: request.Username,
	}
	json.NewEncoder(w).Encode(response)
}

// Unfollow a user
func unfollowThisUser(username string, profileToUnfollow Profile) error {
	if _, exists := users[profileToUnfollow.Username]; exists {
		users[username] = User{
			Username: username,
			Profile: Profile{
				Username:       users[username].Username,
				Posts:          users[username].Profile.Posts,
				NumberOfPhotos: users[username].Profile.NumberOfPhotos,
				Followers:      users[username].Profile.Followers,
				Following:      removeFollowing(users[username].Profile.Following, profileToUnfollow.Username),
			},
			BannedUsers: users[username].BannedUsers,
		}
		return nil
	}
	return fmt.Errorf("User with %s username not found", username)
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
