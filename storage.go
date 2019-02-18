package seed

import "strings"

import "github.com/qlova/seed/storage"
import "github.com/qlova/seed/storage/bolt"

var Database = bolt.Open(Dir+"/seed.db")

func Store(path string) storage.JSON {

	var view = storage.View{
		Node: Database,
		Path: strings.Split(path, "/"),
	}

	view.Create()

	return view.JSON()
}