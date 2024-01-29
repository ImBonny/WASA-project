package database

import "errors"

func (db *appdbimpl) DeleteUser(id uint64) error {
	var exists bool
	err := db.c.QueryRow("SELECT EXISTS(SELECT 1 FROM userDb WHERE userid = ?)", id).Scan(&exists)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("user does not exist")
	}
	_, err = db.c.Exec("DELETE FROM userDb WHERE userid = ?", id)
	if err != nil {
		return err
	}
	return nil
}
