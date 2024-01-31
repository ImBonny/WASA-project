package database

import "errors"

func (db *appdbimpl) CheckAuthorization(token uint64) (bool, error) {
	if token == 0 {
		return false, errors.New("token is invalid")
	} else {
		var id uint64
		err := db.c.QueryRow("SELECT UserId FROM authorizedDb WHERE UserId = ?", token).Scan(&id)
		if err != nil {
			return false, err
		}
		return false, nil
	}
}
