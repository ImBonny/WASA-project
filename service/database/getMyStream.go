package database

import "sort"

func (db *appdbimpl) GetMyStream(id uint64) (*[]uint64, error) {
	var posts []uint64

	// Get user's followings
	following, err := db.GetFollowing(id)
	if err != nil {
		return nil, err
	}

	// Get user's following posts
	for _, user := range *following {
		var userPosts []uint64
		userPosts, err = db.GetPosts(user.UserId)
		if err != nil {
			return nil, err
		}
		posts = append(posts, userPosts...)
	}

	sort := func(p []uint64) {
		sort.SliceStable(p, func(i, j int) bool {
			return p[i] > p[j]
		})
	}
	sort(posts)
	return &posts, nil

}
