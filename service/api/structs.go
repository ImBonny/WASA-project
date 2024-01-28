package api

// User represents a user
type User struct {
	Username    string   `json:"username"`
	Profile     Profile  `json:"profile"`
	BannedUsers []string `json:"bannedUsers"`
	UserId      uint64   `json:"id"`
}

func (u User) isBanned(following string) bool {
	for _, username := range users[following].BannedUsers {
		if username == u.Username {
			return true
		}
	}
	return false
}

// Comment represents a comment
type Comment struct {
	CommentOwner string `json:"commentOwner"`
	CommentText  string `json:"commentText"`
	CreationTime string `json:"creationTime"`
	CommentId    int    `json:"commentId"`
}

// Like represents a like
type Like struct {
	LikeOwner uint64 `json:"likeOwner"`
	PostId    uint64 `json:"postId"`
}

// Profile represents a user's profile
type Profile struct {
	Username       string   `json:"username"`
	Posts          []int    `json:"posts"`
	NumberOfPhotos int      `json:"numberOfPhotos"`
	Followers      []string `json:"followers"`
	Following      []string `json:"following"`
}

// Post represents a post
type Post struct {
	PostOwner    string    `json:"postOwner"`
	Image        string    `json:"image"`
	Comments     []Comment `json:"comments"`
	NComments    uint      `json:"nComments"`
	Likes        []Like    `json:"likes"`
	NLikes       uint      `json:"nLikes"`
	CreationTime string    `json:"creationTime"`
	PostId       int       `json:"postId"`
}
