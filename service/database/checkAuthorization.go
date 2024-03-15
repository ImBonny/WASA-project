package database

import "errors"

func (db *appdbimpl) CheckAuthorization(token uint64) (bool, error) {
	if token == 0 {
		return false, errors.New("token is invalid!")
	} else {
		_, err := db.c.Exec("SELECT UserId FROM authorizedDb WHERE UserId = ?", token)

		if err != nil {
			return false, err
		}
		return true, nil
	}
}
