package record

import "time"

// UnitRecord single db record
type UnitRecord struct {
	UnitDate  time.Time
	UnitValue float64
}

// UnitRecords db records
type UnitRecords []UnitRecord

func (records UnitRecords) Len() int      { return len(records) }
func (records UnitRecords) Swap(i, j int) { records[i], records[j] = records[j], records[i] }
func (records UnitRecords) Less(i, j int) bool {
	return records[i].UnitDate.Before(records[j].UnitDate)
}
