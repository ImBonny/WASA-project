package database

func (db *appdbimpl) GetPosts(userid uint64) ([]uint64, error) {
	var posts []uint64
	rows, err := db.c.Query("SELECT postId FROM postDb WHERE postOwner = ?", userid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var post uint64
		err = rows.Scan(&post)
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
