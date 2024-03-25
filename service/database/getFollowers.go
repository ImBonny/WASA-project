package database

func (db *appdbimpl) GetFollowers(userid uint64) (*[]Database_user, error) {
	var followers []Database_user
	rows, err := db.c.Query("SELECT username, userid FROM userDb WHERE userid IN (SELECT userFollowingId FROM followersDb WHERE userToFollowId = ?)", userid)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var follower Database_user
		err1 := rows.Scan(&follower.Username, &follower.UserId)
		if err1 != nil {
			panic(err1)
		}
		followers = append(followers, follower)
	}
	err2 := rows.Err()
	if err2 != nil {
		return nil, err2
	}
	return &followers, nil
}
