package database

type Database_user struct {
	Username string
	UserId   uint64
}

type Database_photo struct {
	PostOwner    uint64
	Image        []byte
	Description  string
	NComments    uint
	NLikes       uint
	CreationTime string
	PostId       uint64
}

type Database_like struct {
	PostId       uint64
	LikeOwner    uint64
	CreationTime string
}

type Database_comment struct {
	PostId       uint64
	CommentOwner uint64
	CommentText  string
	CreationTime string
	CommentId    uint64
}

type Database_profile struct {
	Username       string
	Posts          []Database_photo
	NumberOfPhotos int
}
