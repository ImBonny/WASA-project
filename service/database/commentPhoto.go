package database

import "time"

func (db *appdbimpl) CommentPhoto(userid uint64, photoid uint64, commenttext string) (uint64, error) {
	res, err := db.c.Exec("INSERT INTO commentDb (commentOwner, postId, content, creationTime) VALUES (?, ?, ?, ?)", userid, photoid, commenttext, time.Now())
	if err != nil {
		return 0, err
	}
	lastInsertId, err := res.LastInsertId()
	db.c.Exec("UPDATE postDb SET nComments = nComments + 1 WHERE postId = ?", photoid)
	return uint64(lastInsertId), nil
}
