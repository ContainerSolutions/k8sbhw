package datastore

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // gorm
)

type gormRecord struct {
	gorm.Model
	Key   string `gorm:"not_null;unique"`
	Value string
}

// SQLDataStore is a datastore that uses an SQL server
type SQLDataStore struct {
	driver     string
	connection string
	db         *gorm.DB
}

// NewSQLDatastore returns a new SQLDataStore
func NewSQLDatastore() *SQLDataStore {
	return &SQLDataStore{}
}

// Init the SQL datastore, connect to the database
func (d *SQLDataStore) Init(parameters ...interface{}) error {
	var err error
	d.driver = parameters[0].(string)
	d.connection = parameters[1].(string)
	d.db, err = gorm.Open(d.driver, d.connection)
	if err == nil {
		d.db.AutoMigrate(&gormRecord{})
	}
	return err
}

// Get all records from the datastore
func (d *SQLDataStore) Get() []Record {
	var gormRecords []gormRecord
	records := make([]Record, 0)
	d.db.Find(&gormRecords)
	for _, gormRecord := range gormRecords {
		records = append(records, Record{gormRecord.Key, gormRecord.Value})
	}
	return records
}

// Add the Record record to the datastore
func (d *SQLDataStore) Add(record Record) {
	if len(record.Key) == 0 {
		return
	}
	d.db.Where(gormRecord{Key: record.Key}).
		Assign(gormRecord{Key: record.Key, Value: record.Value}).
		FirstOrCreate(&gormRecord{Key: record.Key, Value: record.Value})
}

// Rem the Record record from the datastore
func (d *SQLDataStore) Rem(record Record) {
	if len(record.Key) == 0 {
		return
	}
	d.db.Unscoped().Where("Key = ?", record.Key).Delete(&gormRecord{})
}
