package database

import "time"

func (db *appdbimpl) UploadPhoto(id uint64, photo string, caption string) (uint64, error) {
	res, err := db.c.Exec("INSERT INTO postDb (postOwner, image, description, creationTime) VALUES (?, ?, ?, ?)", id, photo, caption, time.Now())
	if err != nil {
		return 0, err
	}
	lastInsertId, err := res.LastInsertId()
	return uint64(lastInsertId), nil
}
