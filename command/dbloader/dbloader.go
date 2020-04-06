package main

import (
	"log"

	"github.com/wayming/superdata/internal/db"
)

func main() {
	var mydb = db.Connect()

	tx, err := mydb.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare("CREATE TABLE SUNSUPER (number integer)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec()

	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}

	db.Disconnect(mydb)
}
