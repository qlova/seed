package client

import "qlova.org/seed/use/js"

type File interface {
	Value
	GetFile() js.Value
}
