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
		err = rows.Scan(&follower.Username, &follower.UserId)
		if err != nil {
			panic(err)
		}
		followers = append(followers, follower)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return &followers, nil
}
