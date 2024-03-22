package database

import "errors"

func (db *appdbimpl) UnlikePhoto(userID uint64, postID uint64) error {
	var alreadyLiked bool
	err := db.c.QueryRow("SELECT EXISTS(SELECT 1 FROM likesDb WHERE userId = ? AND postId = ?)", userID, postID).Scan(&alreadyLiked)
	if err != nil {
		return err
	}
	if !alreadyLiked {
		return errors.New("not already liked")
	}

	_, err = db.c.Exec("DELETE FROM likesDb WHERE postId = ? AND userId = ?", postID, userID)
	if err != nil {
		return err
	}
	_, err = db.c.Exec("UPDATE postDb SET nLikes = nLikes - 1 WHERE postId = ?", postID)
	if err != nil {
		return err
	}
	return err
}
