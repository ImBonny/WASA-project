package database

import "time"

func (db *appdbimpl) UploadPhoto(id uint64, photo []byte, caption string) (uint64, error) {
	res, err := db.c.Exec("INSERT INTO postDb (postOwner, image, description, creationTime, nComments, nLikes) VALUES (?, ?, ?, ?,0,0)", id, photo, caption, time.Now())
	if err != nil {
		return 0, err
	}
	lastInsertId, err := res.LastInsertId()
	if err != nil {
		panic(err)
		return 0, err
	}
	return uint64(lastInsertId), nil
}
