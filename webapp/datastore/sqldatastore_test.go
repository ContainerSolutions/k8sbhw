package datastore

import (
	"fmt"
	"os"
	"testing"
	"time"

	_ "github.com/jinzhu/gorm/dialects/sqlite" // gorm
)

func TestSQLDatastore(t *testing.T) {
	tempFile := fmt.Sprintf("/tmp/sqldatastoretest-%d.db", time.Now().UnixNano())
	defer os.Remove(tempFile)
	dataStore := NewSQLDatastore()
	dataStore.Init("sqlite3", tempFile)

	t.Run("Add a record", func(*testing.T) {
		record := Record{"c", "d"}
		dataStore.Add(record)

		records := []gormRecord{}
		dataStore.db.Find(&records)
		if len(records) != 1 {
			t.Errorf("%d != %d", len(records), 1)
		}
		if records[0].Key != "c" || records[0].Value != "d" {
			t.Errorf("%v != %v", records[0], record)
		}
	})

	t.Run("Add another record", func(*testing.T) {
		record := Record{"a", "v"}
		dataStore.Add(record)

		records := []gormRecord{}
		dataStore.db.Find(&records)
		if len(records) != 2 {
			t.Errorf("%d != %d", len(records), 1)
		}
		if records[1].Key != "a" || records[1].Value != "v" {
			t.Errorf("%v != %v", records[1], record)
		}
	})

	t.Run("Add over existing record", func(*testing.T) {
		record := Record{"a", "g"}
		dataStore.Add(record)

		records := []gormRecord{}
		dataStore.db.Find(&records)
		if len(records) != 2 {
			t.Errorf("%d != %d", len(records), 1)
		}
		if records[1].Key != "a" || records[1].Value != "g" {
			t.Errorf("%v != %v", records[1], record)
		}
	})

	t.Run("Remove a record", func(*testing.T) {
		dataStore.Rem(Record{"a", ""})

		records := []gormRecord{}
		dataStore.db.Find(&records)
		if len(records) != 1 {
			t.Errorf("%d != %d", len(records), 1)
		}
	})

	t.Run("Remove non existing record", func(*testing.T) {
		dataStore.Rem(Record{"x", ""})

		records := []gormRecord{}
		dataStore.db.Find(&records)
		if len(records) != 1 {
			t.Errorf("%d != %d", len(records), 1)
		}
	})

	t.Run("Add after Remove", func(*testing.T) {
		record := Record{"a", "g"}
		dataStore.Add(record)

		records := []gormRecord{}
		dataStore.db.Find(&records)
		if len(records) != 2 {
			t.Errorf("%d != %d", len(records), 1)
		}
		if records[1].Key != "a" || records[1].Value != "g" {
			t.Errorf("%v != %v", records[1], record)
		}
	})

	t.Run("Remove remaining records", func(*testing.T) {
		dataStore.Rem(Record{"c", ""})
		dataStore.Rem(Record{"a", ""})

		records := []gormRecord{}
		dataStore.db.Find(&records)
		if len(records) != 0 {
			t.Errorf("%d != %d", len(records), 0)
		}
	})

	t.Run("Remove on empty", func(*testing.T) {
		dataStore.Rem(Record{"c", ""})

		records := []gormRecord{}
		dataStore.db.Find(&records)
		if len(records) != 0 {
			t.Errorf("%d != %d", len(records), 0)
		}
	})
}
