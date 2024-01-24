package database

func (db *appdbimpl) BanUser(id int, to_ban_id int) error {
	_, err := db.c.Exec("INSERT INTO banned (Userid, BannedUserid) VALUES (?, ?, ?, ?)", id, to_ban_id)
	return err
}
