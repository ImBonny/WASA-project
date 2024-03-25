package database

import (
	"database/sql"
	"errors"
)

func (db *appdbimpl) SearchUser(username string) (string, error) {
	var user string
	err := db.c.QueryRow("SELECT username FROM userDb WHERE username = ?", username).Scan(&user)
	if errors.Is(err, sql.ErrNoRows) {

		return "", nil // Utente non trovato, restituisci un errore nullo
	}
	if err != nil {
		return "", err
	}
	return user, nil
}
