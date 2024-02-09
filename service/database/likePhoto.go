package database

import (
	"errors"
	"time"
)

func (db *appdbimpl) LikePhoto(postID uint64, userID uint64) error {
	var alreadyLiked bool
	err := db.c.QueryRow("SELECT EXISTS(SELECT 1 FROM likesDb WHERE userId = ? AND postId = ?)", userID, postID).Scan(&alreadyLiked)
	if err != nil {
		return err
	}
	if alreadyLiked {
		return errors.New("already liked")
	}

	_, err = db.c.Exec("INSERT INTO likesDb (postId, userId, creationTime) VALUES (?, ?, ?)", postID, userID, time.Now())
	_, err = db.c.Exec("UPDATE postDb SET nLikes = nLikes+1 WHERE postId = ?", postID)
	if err != nil {
		return err
	}
	return err
}
