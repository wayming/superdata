package loader

import (
	"database/sql"
	"log"

	"github.com/wayming/superdata/internal/db"
)

// Loader db loader
type Loader struct {
	TableName  string
	DateFormat string
	conn       *sql.DB
}

// Connect Loader.Connect
func (loader *Loader) Connect() {
	loader.conn = db.Connect()
}

// Disconnect Loader.Disconnect
func (loader *Loader) Disconnect() {
	db.Disconnect(loader.conn)
}

// Create Loader.Create
func (loader *Loader) Create() {
	tx, err := loader.conn.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()

	dropStmt, err := tx.Prepare(
		"DROP TABLE if exists " + loader.TableName)
	if err != nil {
		log.Fatal(err)
	}
	defer dropStmt.Close()

	_, err = dropStmt.Exec()
	if err != nil {
		log.Fatal(err)
	}

	createStmt, err := tx.Prepare(
		"CREATE TABLE " + loader.TableName +
			" (vdate date, value integer)")
	if err != nil {
		log.Fatal(err)
	}
	defer createStmt.Close()

	_, err = createStmt.Exec()
	if err != nil {
		log.Fatal(err)
	}

	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}
}
