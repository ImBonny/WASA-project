package database

import "errors"

func (db *appdbimpl) followUser(userFollowingId uint64, userToFollowId uint64) error {
	// Check if the follower is already following the user
	var alreadyFollowing bool
	err := db.c.QueryRow("SELECT EXISTS(SELECT 1 FROM followersDb WHERE userFollowingId = ? AND userToFollowId = ?)", userFollowingId, userToFollowId).Scan(&alreadyFollowing)
	if err != nil {
		return err
	}
	if alreadyFollowing {
		return errors.New("already following")
	}

	// Insert the follow into the database
	_, err = db.c.Exec("INSERT INTO followersDb (userFollowingId, userToFollowId) VALUES (?, ?)", userFollowingId, userToFollowId)
	if err != nil {
		return err
	}

	return nil
}
