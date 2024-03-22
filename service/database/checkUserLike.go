package database

func (db *appdbimpl) CheckUserLike(userID uint64, postID uint64) (bool, error) {
	var alreadyLiked bool
	err := db.c.QueryRow("SELECT EXISTS(SELECT 1 FROM likesDb WHERE userId = ? AND postId = ?)", userID, postID).Scan(&alreadyLiked)
	if err != nil {
		return false, err
	}
	return alreadyLiked, nil
}
