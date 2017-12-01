package datastore

//Record contains a Key and a Value
type Record struct {
	Key   string
	Value string
}

//Datastore is a datastore, man
type Datastore interface {
	Init(...interface{}) error
	Get() []Record
	Add(Record)
	Rem(Record)
}
