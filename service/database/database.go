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
	DoLogin(username string) (*string, error)

	// AddUser creates a user and adds it to the database//
	AddUser(username string, id string) error

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
	FollowUser(to_add_id int, id int) error

	// UnfollowUser //
	UnfollowUser(id int, to_del_id int) error

	// BanUser //
	BanUser(id int, to_ban_id int) error

	// UnbanUser //
	UnbanUser(id string, to_del_id string) error

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
	CommentPhoto(userid string, photoid string, commentid string, commentauthor string, commenttext string) error

	// UncommentPhoto //
	UncommentPhoto(userid string, photoid string, commentid string, commentauthor string) error

	// Ping checks whether the database is available or not (in that case, an error will be returned)
	Ping() error

	// Funzioni ausiliarie definite in database_utilities
	CheckAuthorization(request *http.Request, username string) error
	CheckUserExistence(username string) error
	GetUsernameFromId(id string) (*string, error)
	IsAllowed(id1 string, id2 string) error
	CheckPhotoExistence(user_id string, photoid string) error
	GetIdFromUsername(user string) (int, error)
}

type appdbimpl struct {
	c *sql.DB
}

// New returns a new instance of AppDatabase based on the SQLite connection `db`.
// `db` is required - an error will be returned if `db` is `nil`.
func New(db *sql.DB) (AppDatabase, error) {
	if db == nil {
		return nil, errors.New("database is required when building a AppDatabase")
	}

	// Check if table exists. If not, the database is empty, and we need to create the structure
	var tableName string
	err := db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='users';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {

		// Creazione della tabella authstrings
		// authstrings memorizza per ogni username registrato l'id univoco che riconosce l'utente nel sistema e nelle richieste
		sqlStmt := `CREATE TABLE authstrings (username TEXT NOT NULL PRIMARY KEY,id TEXT NOT NULL);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("errore nella creazione della tabella authstrings: %w", err)
		}
		// Creazione della tabella users
		// users memorizza  il profilo per ogni username registrato in authstrings
		sqlStmt = `CREATE TABLE users (username TEXT NOT NULL PRIMARY KEY,followers INTEGER NOT NULL,following INTEGER NOT NULL,numberofphotos INTEGER NOT NULL);`
		_, err = db.Exec(sqlStmt)
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
