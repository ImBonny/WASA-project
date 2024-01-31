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
)

// AppDatabase is the high level interface for the DB
type AppDatabase interface {

	// DoLogin resitituisce l'id relativo all'username passato come argomento. //
	// se l'username non è registrato verrà creato e restituito un nuovo id,altrimenti verrò resituito quello esistente //
	DoLogin(username string, token uint64) (uint64, error)

	// GetUserProfile gets a user profile searched via username //
	GetUserProfile(username string) (*Database_profile, error)

	// GetMyStream //
	GetMyStream(id uint64) (*[]Database_photo, error)

	// Deleteuser deletes the user from the system //
	DeleteUser(id uint64) error

	// SetMyUsername modifies the username of a user //
	SetMyUsername(userid uint64, new_username string) error

	// GetFollowers //
	GetFollowers(id uint64) (*[]Database_user, error)

	// GetFollowing //
	GetFollowing(id uint64) (*[]Database_user, error)

	// GetBanned //
	GetBanned(id uint64) (*[]Database_user, error)

	// FollowUser //
	FollowUser(to_add_id uint64, id uint64) error

	// UnfollowUser //
	UnfollowUser(id uint64, to_del_id uint64) error

	// BanUser //
	BanUser(id uint64, to_ban_id uint64) error

	// UnbanUser //
	UnbanUser(id uint64, to_del_id uint64) error

	// UploadPhoto //
	UploadPhoto(id uint64, image string, caption string) (uint64, error)

	// DeletePhoto //
	DeletePhoto(postId uint64) error

	// GetLikes //
	GetLikes(photoid uint64) (*[]Database_like, error)

	// LikePhoto //
	LikePhoto(userid uint64, photoid uint64) error

	// UnlikePhoto //
	UnlikePhoto(userid uint64, photoid uint64) error

	// GetComments //
	GetComments(photoid uint64) (*[]Database_comment, error)

	// CommentPhoto //
	CommentPhoto(userid uint64, photoid uint64, commenttext string) (uint64, error)

	// UncommentPhoto //
	UncommentPhoto(commentid uint64) error

	SearchUser(username string) (Database_user, error)

	// Ping checks whether the database is available or not (in that case, an error will be returned)
	Ping() error

	// Funzioni ausiliarie definite in database_utilities
	CheckAuthorization(token uint64) (bool, error)
	GetUsernameFromId(id uint64) (string, error)
	GetIdFromUsername(user string) (uint64, error)
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
	err := db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='userDb';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		// Creazione della tabella users
		// userDb memorizza gli username e gli id di ogni utente
		userDb := `CREATE TABLE userDb (username TEXT NOT NULL UNIQUE, UserId INTEGER PRIMARY KEY AUTOINCREMENT);`
		_, err = db.Exec(userDb)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='authorizedDb';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		// Creazione della tabella users
		// userDb memorizza gli username e gli id di ogni utente
		authorizedDb := `CREATE TABLE authorizedDb (UserId INTEGER PRIMARY KEY);`
		_, err = db.Exec(authorizedDb)
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
                         FOREIGN KEY (userFollowingId) REFERENCES userDb(UserId),
                         FOREIGN KEY (userToFollowId) REFERENCES userDb(UserId),
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
                         FOREIGN KEY (userBanningId) REFERENCES userDb(UserId),
                         FOREIGN KEY (userToBanId) REFERENCES userDb(UserId),
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
										 nComments INTEGER NOT NULL,
										 nLikes INTEGER NOT NULL,
										 creationTime DATETIME NOT NULL,
										 FOREIGN KEY (postOwner) REFERENCES userDb(UserId)
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
										 FOREIGN KEY (commentOwner) REFERENCES userDb(UserId),
                    					 FOREIGN KEY (postId) REFERENCES postDb(postId)
                         );`
		_, err = db.Exec(commentDb)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}

	}

	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='likesDb';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		// commentDb memorizza i commenti
		likesDb := `CREATE TABLE likesDb (postId INTEGER NOT NULL,
										 userId INTEGER NOT NULL,
										 creationTime TEXT NOT NULL,
										 FOREIGN KEY (userId) REFERENCES userDb(UserId),
                    					 FOREIGN KEY (postId) REFERENCES postDb(postId),
										 PRIMARY KEY (postId, userId)
                         );`
		_, err = db.Exec(likesDb)
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
