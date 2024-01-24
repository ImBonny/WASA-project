package database

func (db *appdbimpl) GetIdFromUsername(username string) (uint64, error) {
	var id uint64
	err := db.c.QueryRow("SELECT Userid FROM user WHERE Username=?", username).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
