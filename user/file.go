package user

import (
	"io"
	"mime/multipart"
)

//File is an attached file.
type File struct {
	head *multipart.FileHeader
	file multipart.File
}

//Name returns the name of the file.
func (f File) Name() string {
	if f.head == nil {
		return ""
	}

	return f.head.Filename
}

func (f File) Read(b []byte) (int, error) {
	if f.head == nil {
		return 0, io.EOF
	}
	return f.file.Read(b)
}

func (f File) Close() error {
	return nil
}

//Size returns the size of the file.
func (f File) Size() int64 {
	if f.head == nil {
		return 0
	}

	return f.head.Size
}

//File decodes the file of the given argument.
func (arg Arg) File() File {
	file, header, err := arg.u.r.FormFile(arg.i)
	if err != nil {
		return File{}
	}

	return File{
		head: header,
		file: file,
	}
}
