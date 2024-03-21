package database

func (db *appdbimpl) GetImage(imageId string) (image Database_photo, err error) {
	var out Database_photo
	err = db.c.QueryRow("SELECT postOwner, image, description, nComments, nLikes, creationTime, postId FROM postDb WHERE postId = ?", imageId).Scan(&out.PostOwner, &out.Image, &out.Description, &out.nComments, &out.nLikes, &out.creationTime, &out.PostId)
	if err != nil {
		return out, err
	}
	return out, nil
}
