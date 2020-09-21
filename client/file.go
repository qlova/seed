package client

import "qlova.org/seed/js"

type File interface {
	Value
	GetFile() js.Value
}
