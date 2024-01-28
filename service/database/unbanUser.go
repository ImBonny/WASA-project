package database

import "errors"

func (db *appdbimpl) UnbanUser(id uint64, toUnbanId uint64) error {
	var alreadyBanned bool
	err := db.c.QueryRow("SELECT EXISTS(SELECT 1 FROM bannedDb WHERE userBaningId = ? AND userToBanId = ?)", id, toUnbanId).Scan(&alreadyBanned)
	if err != nil {
		return err
	}
	if !alreadyBanned {
		return errors.New("already banned")
	}

	_, err = db.c.Exec("DELETE FROM bannedDb WHERE userBanningId=? AND userToBanId=?", id, toUnbanId)
	return err
}
