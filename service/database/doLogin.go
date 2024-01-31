package database

// creates a new User in the database
func (db *appdbimpl) DoLogin(username string, token uint64) (uint64, error) {
	// check if user exists
	var exists bool
	err := db.c.QueryRow("SELECT EXISTS(SELECT Username FROM userDb WHERE Username = ?)", username).Scan(&exists)
	if err != nil {
		return 0, err
	}
	if !exists {
		_, err := db.c.Exec("INSERT INTO authorizedDb (UserId) VALUES (?)", token)
		if err != nil {
			return 0, err
		}
	}
	return token, nil
}
