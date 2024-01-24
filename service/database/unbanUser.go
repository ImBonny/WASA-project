package database

func (db *appdbimpl) UnbanUser(id int, to_unban_id int) error {
	_, err := db.c.Exec("DELETE FROM banned WHERE Userid=? AND Username=? AND BannedUserid=? AND BannedUsername=?", id, to_unban_id)
	return err
}
