/*
Package database is the middleware between the app database and the code. All data (de)serialization (save/load) from a
persistent database are handled here. Database specific logic should never escape this package.

To use this package you need to apply migrations to the database if needed/wanted, connect to it (using the database
data source name from config), and then initialize an instance of AppDatabase from the DB connection.

For example, this code adds a parameter in `webapi` executable for the database data source name (add it to the
main.WebAPIConfiguration structure):

	DB struct {
		Filename string `conf:""`
	}

This is an example on how to migrate the DB and connect to it:

	// Start Database
	logger.Println("initializing database support")
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		logger.WithError(err).Error("error opening SQLite DB")
		return fmt.Errorf("opening SQLite: %w", err)
	}
	defer func() {
		logger.Debug("database stopping")
		_ = db.Close()
	}()

Then you can initialize the AppDatabase and pass it to the api package.
*/
package database

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
)

type Database_user struct {
	Username string
	UserId   uint64
}

type Database_follower struct {
	Username string
}

type Database_following struct {
	Username string
}

type Database_banned struct {
	Username       string
	Userid         uint64
	BannedUsername string
	BannedUserid   uint64
}

type Database_photo struct {
	postOwner    string
	image        string
	nComments    uint
	nLikes       uint
	creationTime string
	postId       string
}

type Database_like struct {
	postId       string
	likeOwner    string
	creationTime string
	likeId       string
}

type Database_comment struct {
	postId       int
	commentOwner string
	commentText  string
	creationTime string
	commentId    int
}

type Database_profile struct {
	Username       string
	Posts          []string
	NumberOfPhotos uint
}

type Database_photostream_component struct {
	stream []Database_photo
}

// AppDatabase is the high level interface for the DB
type AppDatabase interface {

	// DoLogin resitituisce l'id relativo all'username passato come argomento. //
	// se l'username non è registrato verrà creato e restituito un nuovo id,altrimenti verrò resituito quello esistente //
	DoLogin(username string) (uint64, error)

	// GetUserProfile gets a user profile searched via username //
	GetUserProfile(username string) (*Database_profile, error)

	// GetMyStream //
	GetMyStream(id string) (*[]Database_photostream_component, error)

	// Deleteuser deletes the user from the system //
	DeleteUser(id string, username string) error

	// SetMyUsernamer modifies the username of a user //
	SetMyUsername(old_username string, new_username string) error

	// GetFollowers //
	GetFollowers(id string) (*[]Database_follower, error)

	// GetFollowing //
	GetFollowing(id string) (*[]Database_following, error)

	// GetBanned //
	GetBanned(id string) (*[]Database_banned, error)

	// FollowUser //
	FollowUser(to_add_id uint64, id uint64) error

	// UnfollowUser //
	UnfollowUser(id uint64, to_del_id uint64) error

	// BanUser //
	BanUser(id uint64, to_ban_id uint64) error

	// UnbanUser //
	UnbanUser(id uint64, to_del_id uint64) error

	// UploadPhoto //
	UploadPhoto(photo Database_photo, id string) error

	// DeletePhoto //
	DeletePhoto(userid string, photoid string) error

	// GetLikes //
	GetLikes(photoid string) (*[]Database_like, error)

	// LikePhoto //
	LikePhoto(userid string, photoid string, likeid string) error

	// UnlikePhoto //
	UnlikePhoto(userid string, photoid string, likeid string) error

	// GetComments //
	GetComments(photoid string) (*[]Database_comment, error)

	// CommentPhoto //
	CommentPhoto(userid uint64, photoid uint64, commenttext string) (uint64, error)

	// UncommentPhoto //
	UncommentPhoto(userid uint64, photoid uint64, commentid uint64, commentauthor uint64) error

	// Ping checks whether the database is available or not (in that case, an error will be returned)
	Ping() error

	// Funzioni ausiliarie definite in database_utilities
	CheckAuthorization(request *http.Request, username string) error
	CheckUserExistence(username string) error
	GetUsernameFromId(id uint64) (string, error)
	IsAllowed(id1 uint64, id2 uint64) error
	CheckPhotoExistence(user_id uint64, photoid uint64) error
	GetIdFromUsername(user string) (uint64, error)
}

type appdbimpl struct {
	c *sql.DB
}

func (db *appdbimpl) GetUserProfile(username string) (*Database_profile, error) {
	//TODO implement me
	panic("implement me")
}

func (db *appdbimpl) GetMyStream(id string) (*[]Database_photostream_component, error) {
	//TODO implement me
	panic("implement me")
}

func (db *appdbimpl) DeleteUser(id string, username string) error {
	//TODO implement me
	panic("implement me")
}

func (db *appdbimpl) SetMyUsername(old_username string, new_username string) error {
	//TODO implement me
	panic("implement me")
}

func (db *appdbimpl) GetFollowers(id string) (*[]Database_follower, error) {
	//TODO implement me
	panic("implement me")
}

func (db *appdbimpl) GetFollowing(id string) (*[]Database_following, error) {
	//TODO implement me
	panic("implement me")
}

func (db *appdbimpl) GetBanned(id string) (*[]Database_banned, error) {
	//TODO implement me
	panic("implement me")
}

func (db *appdbimpl) FollowUser(to_add_id uint64, id uint64) error {
	//TODO implement me
	panic("implement me")
}

func (db *appdbimpl) UnfollowUser(id uint64, to_del_id uint64) error {
	//TODO implement me
	panic("implement me")
}

func (db *appdbimpl) UploadPhoto(photo Database_photo, id string) error {
	//TODO implement me
	panic("implement me")
}

func (db *appdbimpl) DeletePhoto(userid string, photoid string) error {
	//TODO implement me
	panic("implement me")
}

func (db *appdbimpl) GetLikes(photoid string) (*[]Database_like, error) {
	//TODO implement me
	panic("implement me")
}

func (db *appdbimpl) LikePhoto(userid string, photoid string, likeid string) error {
	//TODO implement me
	panic("implement me")
}

func (db *appdbimpl) UnlikePhoto(userid string, photoid string, likeid string) error {
	//TODO implement me
	panic("implement me")
}

func (db *appdbimpl) GetComments(photoid string) (*[]Database_comment, error) {
	//TODO implement me
	panic("implement me")
}

func (db *appdbimpl) UncommentPhoto(userid uint64, photoid uint64, commentid uint64, commentauthor uint64) error {
	//TODO implement me
	panic("implement me")
}

func (db *appdbimpl) CheckAuthorization(request *http.Request, username string) error {
	//TODO implement me
	panic("implement me")
}

func (db *appdbimpl) CheckUserExistence(username string) error {
	//TODO implement me
	panic("implement me")
}

func (db *appdbimpl) IsAllowed(id1 uint64, id2 uint64) error {
	//TODO implement me
	panic("implement me")
}

func (db *appdbimpl) CheckPhotoExistence(user_id uint64, photoid uint64) error {
	//TODO implement me
	panic("implement me")
}

// New returns a new instance of AppDatabase based on the SQLite connection `db`.
// `db` is required - an error will be returned if `db` is `nil`.
func New(db *sql.DB) (AppDatabase, error) {
	if db == nil {
		return nil, errors.New("database is required when building a AppDatabase")
	}

	// Check if table exists. If not, the database is empty, and we need to create the structure
	var tableName string
	err := db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='usersDb';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		// Creazione della tabella users
		// usersDb memorizza gli username e gli id di ogni utente
		usersDb := `CREATE TABLE usersDb (username TEXT NOT NULL UNIQUE, UserId INTEGER PRIMARY KEY AUTOINCREMENT);`
		_, err = db.Exec(usersDb)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='followersDb';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		// Creazione della tabella users
		// followersDb memorizza i seguiti e i seguaci di ogni utente
		followersDb := `CREATE TABLE followersDb (userFollowingId INTEGER NOT NULL,
                         userToFollowId INTEGER NOT NULL,
                         FOREIGN KEY (userFollowingId) REFERENCES usersDb(UserId),
                         FOREIGN KEY (userToFollowId) REFERENCES usersDb(UserId),
                         PRIMARY KEY (userFollowingId, userToFollowId)
                         );`
		_, err = db.Exec(followersDb)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}

	}

	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='bannedDb';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		// Creazione della tabella users
		// followersDb memorizza i bannati
		bannedDb := `CREATE TABLE bannedDb (userBanningId INTEGER NOT NULL,
                         userToBanId INTEGER NOT NULL,
                         FOREIGN KEY (userBanningId) REFERENCES usersDb(UserId),
                         FOREIGN KEY (userToBanId) REFERENCES usersDb(UserId),
                         PRIMARY KEY (userBanningId, userToBanId)
                         );`
		_, err = db.Exec(bannedDb)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}

	}

	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='postDb';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		// Creazione della tabella users
		// postDb memorizza i post
		postDb := `CREATE TABLE postDb (postId INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
										 postOwner INTEGER NOT NULL,
										 image BLOB NOT NULL,
										 description TEXT NOT NULL,
										 uploadTime DATETIME NOT NULL,
										 nComments INTEGER NOT NULL,
										 nLikes INTEGER NOT NULL,
										 creationTime TEXT NOT NULL,
										 FOREIGN KEY (postOwner) REFERENCES usersDb(UserId)
                         );`
		_, err = db.Exec(postDb)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}

	}

	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='commentDb';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		// commentDb memorizza i commenti
		commentDb := `CREATE TABLE commentDb (commentId INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
										 commentOwner INTEGER NOT NULL,
										 postId INTEGER NOT NULL,
										 content TEXT NOT NULL,
										 creationTime TEXT NOT NULL,
										 FOREIGN KEY (commentOwner) REFERENCES usersDb(UserId),
                    					 FOREIGN KEY (postId) REFERENCES postDb(postId)
                         );`
		_, err = db.Exec(commentDb)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}

	}

	return &appdbimpl{
		c: db,
	}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
