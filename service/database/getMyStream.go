package database

import "sort"

func (db *appdbimpl) GetMyStream(id uint64) (*[]Database_photo, error) {
	var posts []Database_photo

	// Get user's followings
	following, err := db.GetFollowing(id)
	if err != nil {
		return nil, err
	}

	// Get user's following posts
	for _, user := range *following {
		var userPosts []Database_photo
		userPosts, err = db.GetPosts(user.UserId)
		if err != nil {
			return nil, err
		}
		posts = append(posts, userPosts...)
	}
	sort := func(p []Database_photo) {
		sort.Slice(p, func(i, j int) bool {
			return p[i].creationTime > p[j].creationTime
		})
	}
	sort(posts)
	return &posts, nil

}
