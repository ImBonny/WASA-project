package database

import "errors"

func (db *appdbimpl) DeletePhoto(postId uint64) error {
	var exists bool
	err := db.c.QueryRow("SELECT EXISTS(SELECT 1 FROM postDb WHERE postId = ? )", postId).Scan(&exists)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("not such Post")
		return nil
	}

	_, err = db.c.Exec("DELETE FROM postDb WHERE postId = ?", postId)
	return err
}
