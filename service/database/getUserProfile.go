package database

func (db *appdbimpl) GetUserProfile(username string) (*Database_profile, error) {
	userid, err := db.GetIdFromUsername(username)
	if err != nil {
		return nil, err
	}
	posts, err := db.GetPosts(userid)
	if err != nil {
		return nil, err
	}
	images, err1 := db.GetImages(&posts)
	if err1 != nil {
		return nil, err1
	}
	profile := Database_profile{Username: username, Posts: *images, NumberOfPhotos: len(posts)}
	return &profile, nil
}
