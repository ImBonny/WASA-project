package database

import "errors"

func (db *appdbimpl) BanUser(id uint64, toBanId uint64) error {
	var alreadyBanned bool
	err := db.c.QueryRow("SELECT EXISTS(SELECT 1 FROM bannedDb WHERE userBanningId = ? AND userToBanId = ?)", id, toBanId).Scan(&alreadyBanned)
	if err != nil {
		return err
	}
	if alreadyBanned {
		return errors.New("already banned")
	}

	_, err = db.c.Exec("INSERT INTO bannedDb (userBanningId, userToBanId) VALUES (?, ?)", id, toBanId)
	if err != nil {
		return err
	}
	_, err = db.c.Exec("DELETE FROM followersDb WHERE userFollowingId = ? AND userToFollowId = ?", toBanId, id)
	if err != nil {
		return err
	}
	return err
}
