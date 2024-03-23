package database

func (db *appdbimpl) IsBanned(id1 uint64, id2 uint64) (bool, error) {
	var banned bool
	err := db.c.QueryRow("SELECT EXISTS(SELECT * FROM bannedDb WHERE userBanningId = ? AND userToBanId = ?)", id1, id2).Scan(&banned)
	if err != nil {
		return false, err
	}
	return banned, nil
}
