package db
import "gopkg.in/mgo.v2"

type DataStore struct {
	collection string
}

func NewDataStore(collection string) *DataStore {
	return &DataStore{collection}
}

func (ds *DataStore) Insert(model interface{}) (err error) {
	query := func(c *mgo.Collection) error {
		fn := c.Insert(model)
		return fn
	}

	create := func() error {
		return WithCollection(ds.collection, query)
	}
	err = create()
	return
}

func (ds *DataStore) Find(q interface{}, skip int, limit int, result interface{}) (err error) {
	query := func(c *mgo.Collection) error {
		fn := c.Find(q).Skip(skip).Limit(limit).All(result)
		if limit < 0 {
			fn = c.Find(q).Skip(skip).All(result)
		}
		return fn
	}
	find := func() error {
		return WithCollection(ds.collection, query)
	}
	err = find()

	return
}

var Villager = NewDataStore("villager")
