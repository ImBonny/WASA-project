package database

func (db *appdbimpl) GetImages(postIds *[]uint64) (*[]Database_photo, error) {
	var images []Database_photo
	for _, postId := range *postIds {
		var image Database_photo
		err := db.c.QueryRow("SELECT postOwner, image, description, nComments, nLikes, creationTime, postId FROM postDb WHERE postId = ?", postId).Scan(&image.PostOwner, &image.Image, &image.Description, &image.NComments, &image.NLikes, &image.CreationTime, &image.PostId)
		if err != nil {
			return nil, err
		}
		images = append(images, image)
	}
	return &images, nil
}
