package loader

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/wayming/superdata/internal/db"
	"github.com/wayming/superdata/internal/record"
)

// Loader db loader
type Loader struct {
	TableName  string
	DateFormat string
	conn       *sql.DB
	data       record.UnitRecords
}

func (loader *Loader) fillGap() {
	sort.Sort(loader.data)

	var nullRecords record.UnitRecords
	for idx, curr := range loader.data {
		if idx == loader.data.Len()-1 {
			// last one
			continue
		}
		nextRecord := loader.data[idx+1]
		nextDay := curr.UnitDate.AddDate(0, 0, 1)
		for nextRecord.UnitDate.After(nextDay) {
			nullRecords = append(nullRecords, record.UnitRecord{nextDay, curr.UnitValue})
			nextDay = nextDay.AddDate(0, 0, 1)
		}
	}

	loader.data = append(loader.data, nullRecords...)
	sort.Sort(loader.data)
}

func (loader *Loader) importRecords() {
	for _, curr := range loader.data {
		sqlStatement :=
			"INSERT INTO " + loader.TableName + " VALUES ($1, $2)" +
				" ON CONFLICT (vdate) DO NOTHING"
		_, err := loader.conn.Exec(sqlStatement, curr.UnitDate, curr.UnitValue)
		if err != nil {
			fmt.Println(
				"Failed to insert line, " +
					curr.UnitDate.String() + "=" + strconv.FormatFloat(curr.UnitValue, 'f', 6, 64) +
					". Error: " + err.Error())
			continue
		}
	}
}

func isNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
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
			" (vdate date PRIMARY KEY, value real)")
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
		// Read each row from csv
		fields, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		if len(fields) < 2 {
			fmt.Println("Skip line " + strings.Join(fields, ","))
			continue
		}

		unitDate, err := time.Parse(loader.DateFormat, strings.TrimSpace(fields[0]))
		if err != nil {
			fmt.Println("Skip line, \"" + fields[0] + "\" is not a date")
			continue
		}

		unitValue, err := strconv.ParseFloat(strings.TrimSpace(fields[1]), 64)
		if err != nil {
			fmt.Println("Skip line, \"" + fields[1] + "\" is not a numeric")
			continue
		}

		loader.data = append(loader.data, record.UnitRecord{unitDate, unitValue})
	}

	loader.fillGap()

	loader.importRecords()
}
