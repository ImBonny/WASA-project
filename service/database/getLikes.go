package database

func (db *appdbimpl) GetLikes(postId uint64) (*[]Database_like, error) {
	var likes []Database_like
	rows, err := db.c.Query("SELECT * FROM likesDb WHERE postId = ?", postId)
	if err != nil {
		panic(err)
	}
	err = rows.Close()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var like Database_like
		err = rows.Scan(&like.postId, &like.likeOwner, &like.creationTime)
		if err != nil {
			panic(err)
		}
		likes = append(likes, like)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return &likes, nil
}
