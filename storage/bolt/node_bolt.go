package bolt

import bolt "go.etcd.io/bbolt"
import "github.com/qlova/seed/storage"
import "encoding/binary"
import "fmt"

//A storage.Node can be thought of as a database or directory.
type Node struct {
	*bolt.DB
}

func Open(path string) Node {
	node, err := bolt.Open(path, 0600, nil)
	if err != nil {
		//Report errors?
	}
	return Node{node}
}

// itob returns an 8-byte big endian representation of v.
func itob(v uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, v)
	return b
}

func GetBucket(tx *bolt.Tx, view storage.View) *bolt.Bucket {
	var bucket = tx.Bucket([]byte(view.Path[0]))

	if len(view.Path) > 1 {
		for _, name := range view.Path {
			bucket = bucket.Bucket([]byte(name))
			if bucket == nil {
				fmt.Println("Could not find bucket for ", view.Path)
				return nil
			}
		}
	}

	return bucket
}

func (node Node) Create(view storage.View) bool {

	if len(view.Path) == 0 {
		return false
	}

	var success bool

	node.Update(func(tx *bolt.Tx) error {

		var bucket, err = tx.CreateBucketIfNotExists([]byte(view.Path[0]))
		if err != nil {
			return nil
		}

		if len(view.Path) > 1 {
			for _, name := range view.Path {
				bucket, err = bucket.CreateBucketIfNotExists([]byte(name))
				if bucket == nil || err != nil {
					fmt.Println("Could not create bucket for ", view.Path)
					return nil
				}
			}
		}

		return nil
	})

	return success
}

func (node Node) Put(view storage.View, data []byte) []byte {

	if len(view.Path) == 0 {
		return nil
	}

	var id []byte

	node.Update(func(tx *bolt.Tx) error {

		var bucket = GetBucket(tx, view)
		if bucket == nil {
			return nil
		}

		numeric_id, err := bucket.NextSequence()
		if err != nil {
			return nil
		}

		if bucket.Put(itob(numeric_id), data) != nil {
			return nil
		}

		id = itob(numeric_id)

		return nil
	})

	return id
}

func (node Node) Set(view storage.View, key []byte, data []byte) bool {

	if len(view.Path) == 0 {
		return false
	}

	var success bool

	node.Update(func(tx *bolt.Tx) error {

		var bucket = GetBucket(tx, view)
		if bucket == nil {
			return nil
		}

		if bucket.Put(key, data) == nil {
			success = true
		}

		return nil
	})

	return success
}

func (node Node) Get(view storage.View, key []byte) []byte {

	if len(view.Path) == 0 {
		return nil
	}

	var result []byte

	node.View(func(tx *bolt.Tx) error {

		var bucket = GetBucket(tx, view)
		if bucket == nil {
			return nil
		}

		data := bucket.Get(key)
		result = make([]byte, len(data))
		copy(result, data)

		return nil
	})

	return result
}

func (node Node) Read(view storage.View, key []byte, handler func(data []byte)) {

	if len(view.Path) == 0 {
		return
	}

	node.View(func(tx *bolt.Tx) error {

		var bucket = GetBucket(tx, view)
		if bucket == nil {
			return nil
		}

		data := bucket.Get(key)
		if data == nil {
			return nil
		}
		handler(data)

		return nil
	})
}

func (node Node) ForEach(view storage.View, f func(key []byte, data []byte)) {
	if len(view.Path) == 0 {
		return
	}

	node.View(func(tx *bolt.Tx) error {

		var bucket = GetBucket(tx, view)
		if bucket == nil {
			return nil
		}

		bucket.ForEach(func(key []byte, data []byte) error {
			f(key, data)
			return nil
		})

		return nil
	})
}
