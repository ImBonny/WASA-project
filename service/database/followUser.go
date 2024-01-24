package database

import "errors"

func (db *appdbimpl) followUser(followerId uint64, followingId uint64) error {
	// Check if the follower is already following the user
	var alreadyFollowing bool
	err := db.c.QueryRow("SELECT EXISTS(SELECT 1 FROM following WHERE follower_id = ? AND following_id = ?)", followerId, followingId).Scan(&alreadyFollowing)
	if err != nil {
		return err
	}
	if alreadyFollowing {
		return errors.New("already following")
	}

	// Insert the follow into the database
	_, err = db.c.Exec("INSERT INTO following (follower_id, following_id) VALUES (?, ?)", followerId, followingId)
	if err != nil {
		return err
	}

	return nil
}
