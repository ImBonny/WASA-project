package database

func (db *appdbimpl) SetMyUsername(id uint64, username string) error {
	_, err := db.c.Exec("UPDATE userDb SET username = ? WHERE userId = ?", username, id)
	if err != nil {
		return err
	}
	return err
}
