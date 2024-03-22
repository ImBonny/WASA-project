package database

func (db *appdbimpl) GetComments(postId uint64) (*[]Database_comment, error) {
	var comments []Database_comment
	rows, err := db.c.Query("SELECT * FROM commentDb WHERE postId = ?", postId)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var comment Database_comment
		err = rows.Scan(&comment.CommentId, &comment.CommentOwner, &comment.PostId, &comment.CommentText, &comment.CreationTime)
		if err != nil {
			panic(err)
		}
		comments = append(comments, comment)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return &comments, nil
}
