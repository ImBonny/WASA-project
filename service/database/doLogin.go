package database

import (
	"database/sql"
	"errors"
)

// creates a new User in the database
func (db *appdbimpl) DoLogin(username string) (uint64, error) {
	res, err := db.c.Exec("INSERT INTO userDb (Username) VALUES (?, ?)", username)
	if err != nil {
		userId := uint64(0)
		if err = db.c.QueryRow("SELECT Username, UserId FROM userDb WHERE Username = ?", username).Scan(&username, &userId); err != nil {
			if err == sql.ErrNoRows {
				return userId, errors.New("user does not exist")
			}
		}
		return userId, nil
	}
	lastInsertId, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint64(lastInsertId), nil
}
