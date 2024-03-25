package database

import (
	"database/sql"
	"errors"
)

func (db *appdbimpl) GetFollowing(id uint64) (*[]Database_user, error) {
	following := []Database_user{}
	rows, err := db.c.Query("SELECT userToFollowId FROM followersDb WHERE userFollowingId = ?", id)
	if errors.Is(err, sql.ErrNoRows) {
		return &following, nil
	}
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var user Database_user
		err1 := rows.Scan(&user.UserId)
		if err1 != nil {
			return nil, err
		}
		var err2 error
		user.Username, err2 = db.GetUsernameFromId(user.UserId)
		if err2 != nil {
			return nil, err2
		}

		following = append(following, user)
	}
	err3 := rows.Err()
	if err3 != nil {
		return nil, err3
	}
	return &following, nil
}
