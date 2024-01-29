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
	profile := Database_profile{Username: username, Posts: posts, NumberOfPhotos: len(posts)}
	return &profile, nil
}
