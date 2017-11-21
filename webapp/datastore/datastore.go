package datastore

//Record contains a Key and a Value
type Record struct {
	Key   string
	Value string
}

//Datastore is a datastore, man
type Datastore interface {
	Get() []Record
	Add(Record)
}

//SliceDataStore is a Datastore that uses a slice as data store
type SliceDataStore struct {
	slice []Record
}

//NewSliceDataStore initializes a new slice based data store
func NewSliceDataStore() *SliceDataStore {
	return &SliceDataStore{make([]Record, 0)}
}

//Get all the elements in the SliceDataStore d
func (d *SliceDataStore) Get() []Record {
	return d.slice
}

//Add an element to the SliceDataStore d
func (d *SliceDataStore) Add(record Record) {
	for i, r := range d.slice {
		if r.Key == record.Key {
			d.slice[i].Value = record.Value
			return
		}
	}

	d.slice = append(d.slice, record)
}
