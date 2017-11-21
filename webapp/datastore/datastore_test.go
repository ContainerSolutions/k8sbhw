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

	t.Run("Modify an existing record", func(*testing.T) {
		record := Record{"a", "v"}
		dataStore.Add(record)
		if len(dataStore.slice) != 2 {
			t.Errorf("%d != %d", len(dataStore.slice), 2)
		}
		if dataStore.slice[0].Key != "a" || dataStore.slice[0].Value != "v" {
			t.Errorf("%s != %s", dataStore.slice[0], record)
		}
	})
}
