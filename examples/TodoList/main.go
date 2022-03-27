package main

import (
	"qlova.tech/use/html/button"
	"qlova.tech/use/html/input"
	"qlova.tech/use/html/linebreak"
	"qlova.tech/use/html/listitem"
	"qlova.tech/use/html/span"
	"qlova.tech/use/html/unorderedlist"
	"qlova.tech/web/data"
	"qlova.tech/web/node"
	"qlova.tech/web/site"
	"qlova.tech/web/tree"
)

type TodoList struct {
	Input string

	Tasks []string
}

func (list *TodoList) RenderTree(seed tree.Seed) tree.Node {
	return tree.New(
		input.New(data.Sync(&list.Input)),
		button.New("Add", data.Push(&list.Tasks, &list.Input)),
		data.When(&list.Tasks,
			linebreak.New(),
			span.New("(click on an item to remove it)"),
		),

		data.Feed(&list.Tasks, unorderedlist.Tag,
			listitem.New(data.View(data.Value), node.OnClick(data.Pull(&list.Tasks, data.Index))),
		),
	)
}

func main() {
	site.Open(new(TodoList))
}
