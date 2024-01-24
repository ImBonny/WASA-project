package database

// creates a new User in the database
func (db *appdbimpl) addUser(username string, id string) error {
	_, err := db.c.Exec("INSERT INTO example_table (Userid, Username) VALUES (?, ?)", id, username)
	return err
}
