package database

func (db *appdbimpl) DoLogin(username string) (uint64, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT Username FROM userDb WHERE Username = '" + username + "')"
	result, err := db.c.Query(query)
	result.Scan(&exists)
	if err != nil {
		return 0, err
	}
	if !exists {
		_, err := db.c.Exec("INSERT INTO userDb (Username) VALUES ('" + username + "')")
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
