package datastore

//Record contains a Key and a Value
type Record struct {
	Key   string
	Value string
}

//Datastore is a datastore, man
type Datastore interface {
	Init(interface{}) error
	Get() []Record
	Add(Record)
	Rem(Record)
}

//SliceDataStore is a Datastore that uses a slice as data store
type SliceDataStore struct {
	slice []Record
}

//NewSliceDataStore initializes a new slice based data store
func NewSliceDataStore() *SliceDataStore {
	return &SliceDataStore{}
}

//Init initializes the SliceDataStore with initial capacity initialCapacity
func (d *SliceDataStore) Init(initialCapacity interface{}) error {
	d.slice = make([]Record, initialCapacity.(int))
	return nil
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

//Rem ove an element from the SliceDataStore d
func (d *SliceDataStore) Rem(record Record) {
	for i, r := range d.slice {
		if r.Key == record.Key {
			d.slice[i] = d.slice[len(d.slice)-1]
			d.slice = d.slice[:len(d.slice)-1]
		}
	}
}

//Size of the SliceDatastore d
func (d *SliceDataStore) Size() int {
	return len(d.slice)
}
