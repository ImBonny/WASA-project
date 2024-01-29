package database

func (db *appdbimpl) SearchUser(username string) (Database_user, error) {
	var user Database_user
	err := db.c.QueryRow("SELECT * FROM userDb WHERE username = ?", username).Scan(&user.Username, &user.UserId)
	if err != nil {
		return Database_user{}, err
	}
	return user, nil
}
