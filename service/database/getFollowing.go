package database

func (db *appdbimpl) GetFollowing(id uint64) (*[]Database_user, error) {
	following := []Database_user{}
	rows, err := db.c.Query("SELECT * FROM followersDb WHERE userFollowingId = ?", id)
	if err != nil {
		return nil, err
	}
	err = rows.Close()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var user Database_user
		err = rows.Scan(&user.Username, &user.UserId)
		if err != nil {
			return nil, err
		}
		following = append(following, user)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return &following, nil
}
