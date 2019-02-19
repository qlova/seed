package seed

import "strings"

import "github.com/qlova/seed/storage"

var Database storage.Node = storage.NewMap()

func Store(path string) storage.JSON {

	var view = storage.View{
		Node: Database,
		Path: strings.Split(path, "/"),
	}

	view.Create()

	return view.JSON()
}