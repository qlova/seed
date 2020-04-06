package script

type File struct {
	Interface
}

func (f File) GetFile() File {
	return f
}

type AnyFile interface {
	AnyValue
	GetFile() File
}
