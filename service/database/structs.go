package database

type Database_user struct {
	Username string
	UserId   uint64
}

type Database_photo struct {
	postOwner    uint64
	image        string
	description  string
	nComments    uint
	nLikes       uint
	creationTime string
	postId       uint64
}

type Database_like struct {
	postId       uint64
	likeOwner    string
	creationTime string
}

type Database_comment struct {
	postId       uint64
	commentOwner string
	commentText  string
	creationTime string
	commentId    uint64
}

type Database_profile struct {
	Username       string
	Posts          []Database_photo
	NumberOfPhotos int
}
