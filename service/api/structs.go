package api

type User struct {
	Username    string   `json:"username"`
	Profile     Profile  `json:"profile"`
	BannedUsers []string `json:"bannedUsers"`
}

type Comment struct {
	CommentOwner string `json:"commentOwner"`
	CommentText  string `json:"commentText"`
	CreationTime string `json:"creationTime"`
	CommentId    int    `json:"commentId"`
}

type Like struct {
	LikeOwner    string `json:"likeOwner"`
	CreationTime string `json:"creationTime"`
	LikeId       int    `json:"likeId"`
}

type Profile struct {
	Username       string   `json:"username"`
	Posts          []int    `json:"posts"`
	NumberOfPhotos int      `json:"numberOfPhotos"`
	Followers      []string `json:"followers"`
	Following      []string `json:"following"`
}
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
