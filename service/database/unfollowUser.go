package database

import "errors"

func (db *appdbimpl) UnfollowUser(followerId uint64, followingId uint64) error {
	// Check if the follower is already following the user
	var alreadyFollowing bool
	err := db.c.QueryRow("SELECT EXISTS(SELECT 1 FROM followersDb WHERE userFollowingId = ? AND userToFollowId = ?)", followerId, followingId).Scan(&alreadyFollowing)
	if err != nil {
		return err
	}
	if !alreadyFollowing {
		return errors.New("not following")
	}

	// Insert the follow into the database
	_, err = db.c.Exec("DELETE FROM followersDb WHERE userFollowingId = ? AND userToFollowId = ?", followerId, followingId)
	if err != nil {
		return err
	}

	return nil
}
