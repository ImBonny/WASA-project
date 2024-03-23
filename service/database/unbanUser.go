package database

func (db *appdbimpl) UnbanUser(id uint64, toUnbanId uint64) error {
	var err error
	_, err = db.c.Exec("DELETE FROM bannedDb WHERE userBanningId=? AND userToBanId=?", id, toUnbanId)
	return err
}
