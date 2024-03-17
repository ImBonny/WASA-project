package database

func (db *appdbimpl) GetUsernameFromId(id uint64) (string, error) {
	var username string
	err := db.c.QueryRow("SELECT username FROM userDb WHERE UserId=?", id).Scan(&username)
	if err != nil {
		return "", err
	}
	return username, nil
}
