package database

import (
	"database/sql"
	"log"

	// Using the blank identifier in order to solely
	// provide the side-effects of the package.
	// Eseentially the side effect is calling the `init()`
	// method of `lib/pq`:
	//	func init () {  sql.Register("postgres", &Driver{} }
	// which you can see at `github.com/lib/pq/conn.go`
	_ "github.com/lib/pq"
)

var db *sql.DB
var err error

// Init is ..
func Init() {
	connStr := "postgres://itsme:enter@127.0.0.1:5432/gorest_db?sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
}

// CreateConn is ..
func CreateConn() *sql.DB {
	return db
}
