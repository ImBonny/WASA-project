package database

func (db *appdbimpl) GetBanned(userid uint64) (*[]Database_user, error) {
	var bannedUsers []Database_user
	rows, err := db.c.Query("SELECT username, userid FROM userDb WHERE userid IN (SELECT userToBanId FROM bannedDb WHERE userBanningId = ?)", userid)
	if err != nil {
		panic(err)
	}
	err = rows.Close()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var banned Database_user
		err = rows.Scan(&banned.Username, &banned.UserId)
		if err != nil {
			panic(err)
		}
		bannedUsers = append(bannedUsers, banned)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return &bannedUsers, nil
}
