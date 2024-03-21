package database

type Database_user struct {
	Username string
	UserId   uint64
}

type Database_photo struct {
	PostOwner    uint64
	Image        []byte
	Description  string
	nComments    uint
	nLikes       uint
	creationTime string
	PostId       uint64
	Comments     []Database_comment
}

type Database_like struct {
	postId       uint64
	likeOwner    string
	creationTime string
}

type Database_comment struct {
	PostId       uint64
	CommentOwner string
	CommentText  string
	CreationTime string
	CommentId    uint64
}

type Database_profile struct {
	Username       string
	Posts          []uint64
	NumberOfPhotos int
}
