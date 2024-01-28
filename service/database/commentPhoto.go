package database

import "time"

func (db *appdbimpl) CommentPhoto(userid string, photoid string, commenttext string) (uint64, error) {
	res, err := db.c.Exec("INSERT INTO commentsDb (UserId, postId, content, creationTime) VALUES (?, ?, ?, ?)", userid, photoid, commenttext, time.Now())
	if err != nil {
		return 0, err
	}
	lastInsertId, err := res.LastInsertId()
	return uint64(lastInsertId), nil
}
