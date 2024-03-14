package database

// creates a new User in the database
func (db *appdbimpl) DoLogin(username string) (uint64, error) {
	// check if user exists
	var exists bool
	err := db.c.QueryRow("SELECT EXISTS(SELECT Username FROM userDb WHERE Username = ?)", username).Scan(&exists)
	if err != nil {
		return 0, err
	}
	if !exists {
		_, err := db.c.Exec("INSERT INTO userDb (Username) VALUES (?)", username)
		if err != nil {
			return 0, err
		}
		token, err1 := db.GetIdFromUsername(username)
		if err1 != nil {
			panic(err1)
		}
		_, err2 := db.c.Exec("INSERT INTO authorizedDb (UserId) VALUES (?)", token)
		if err2 != nil {
			panic(err2)
		}
	}
	token, err3 := db.GetIdFromUsername(username)
	if err3 != nil {
		panic(err3)
	}
	return token, nil
}
