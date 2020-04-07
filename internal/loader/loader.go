package loader

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

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

// Load Loader.Load
func (loader *Loader) Load(filePath string) {
	// Open the file
	csvfile, err := os.Open(filePath)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	// Parse the file
	r := csv.NewReader(csvfile)

	// Iterate through the records
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		if len(record) < 2 {
			fmt.Println("Skip line " + strings.Join(record, ","))
			continue
		}

		unitValue, err := strconv.ParseFloat(strings.TrimSpace(record[1]), 64)
		if err != nil {
			fmt.Println("Skip line, " + record[1] + " is not a numeric")
			continue
		}
		unitDate, err := time.Parse(loader.DateFormat, strings.TrimSpace(record[0]))
		if err != nil {
			fmt.Println("Skip line, " + record[0] + " is not a date")
			continue
		}
		fmt.Println(unitDate.String() + "=" + strconv.FormatFloat(unitValue, 'f', 6, 64))
	}
}

func isNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}
