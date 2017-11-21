package datastore

import "testing"

func TestSliceDatastore(t *testing.T) {
	dataStore := SliceDataStore{}

	t.Run("Add a record", func(*testing.T) {
		record := Record{"a", "b"}
		dataStore.Add(record)
		if dataStore.slice[0].Key != "a" || dataStore.slice[0].Value != "b" {
			t.Errorf("%s != %s", dataStore.slice[0], record)
		}
	})

	t.Run("Add another record", func(*testing.T) {
		record := Record{"c", "d"}
		dataStore.Add(record)
		if len(dataStore.slice) != 2 {
			t.Errorf("%d != %d", len(dataStore.slice), 2)
		}
		if dataStore.slice[1].Key != "c" || dataStore.slice[1].Value != "d" {
			t.Errorf("%s != %s", dataStore.slice[1], record)
		}
	})

	t.Run("Add over existing record", func(*testing.T) {
		record := Record{"a", "v"}
		dataStore.Add(record)
		if len(dataStore.slice) != 2 {
			t.Errorf("%d != %d", len(dataStore.slice), 2)
		}
		if dataStore.slice[0].Key != "a" || dataStore.slice[0].Value != "v" {
			t.Errorf("%s != %s", dataStore.slice[0], record)
		}
	})

	t.Run("Remove a record", func(*testing.T) {
		dataStore.Rem(Record{"a", ""})
		if len(dataStore.slice) != 1 {
			t.Errorf("%d != %d", len(dataStore.slice), 1)
		}
	})

	t.Run("Remove a non existing record", func(*testing.T) {
		dataStore.Rem(Record{"x", ""})
		if len(dataStore.slice) != 1 {
			t.Errorf("%d != %d", len(dataStore.slice), 1)
		}
	})

	t.Run("Add after Remove", func(*testing.T) {
		record := Record{"f", "g"}
		dataStore.Add(record)
		if len(dataStore.slice) != 2 {
			t.Errorf("%d != %d", len(dataStore.slice), 2)
		}
		if dataStore.slice[1].Key != "f" || dataStore.slice[1].Value != "g" {
			t.Errorf("%s != %s", dataStore.slice[1], record)
		}
	})

	t.Run("Remove remaining records", func(*testing.T) {
		dataStore.Rem(Record{"c", ""})
		dataStore.Rem(Record{"f", ""})
		if len(dataStore.slice) != 0 {
			t.Errorf("%d != %d", len(dataStore.slice), 0)
		}
	})

	t.Run("Remove on empty", func(*testing.T) {
		record := Record{"c", "v"}
		dataStore.Rem(record)
		if len(dataStore.slice) != 0 {
			t.Errorf("%d != %d", len(dataStore.slice), 0)
		}
	})
}
