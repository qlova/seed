package script

type File struct {
	Value
}

func (f File) GetFile() File {
	return f
}

type AnyFile interface {
	AnyValue
	GetFile() File
}
