package database

func (db *appdbimpl) CommentPhoto(userid string, photoid string, commentid string, commentauthor string, commenttext string) error {
	_, err := db.c.Exec("INSERT INTO comments (Userid, postId, Commentid, CommentAuthor, CommentText) VALUES (?, ?, ?, ?, ?)", userid, photoid, commentid, commentauthor, commenttext)
	return err
}
