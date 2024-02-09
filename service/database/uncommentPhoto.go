package database

import "errors"

func (db *appdbimpl) UncommentPhoto(commentId uint64) error {
	var exists bool
	err := db.c.QueryRow("SELECT EXISTS(SELECT 1 FROM commentDb WHERE commentId = ? )", commentId).Scan(&exists)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("not such comment")
	}
	var postId uint64
	err = db.c.QueryRow("SELECT postId FROM commentDb WHERE commentId = ?", commentId).Scan(&postId)
	if err != nil {
		panic(err)
	}
	_, err = db.c.Exec("DELETE FROM commentDb WHERE commentId = ?", commentId)
	if err != nil {
		return err
	}
	_, err = db.c.Exec("UPDATE postDb SET nComments = nComments - 1 WHERE postId = ?", postId)
	if err != nil {
		return err
	}
	return err
}
