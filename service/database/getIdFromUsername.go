package database

func (db *appdbimpl) GetIdFromUsername(username string) (uint64, error) {
	var id int
	err := db.c.QueryRow("SELECT UserId FROM userDb WHERE username = ?", username).Scan(&id)
	if err != nil {
		return 0, err
	}
	return uint64(id), nil
}
