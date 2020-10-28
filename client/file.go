package client

import "qlova.org/seed/web/js"

type File interface {
	Value
	GetFile() js.Value
}
