package database

func (db *appdbimpl) GetPosts(userid uint64) ([]Database_photo, error) {
	var posts []Database_photo
	rows, err := db.c.Query("SELECT * FROM postDb WHERE postOwner = ?", userid)
	if err != nil {
		return nil, err
	}
	err = rows.Close()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var post Database_photo
		err = rows.Scan(&post.postId, &post.postOwner, &post.image, &post.description, &post.nComments, &post.nLikes, &post.creationTime)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return posts, nil
}
