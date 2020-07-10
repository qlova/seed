package client

import (
	"io"
	"mime/multipart"
)

//Stream is an incomming data steam from the client.
type Stream struct {
	head *multipart.FileHeader
	file multipart.File
}

//Name returns the name of the steam, if a file is being sent, this could be the name of the file.
func (s Stream) Name() string {
	if s.head == nil {
		return ""
	}

	return s.head.Filename
}

//Read implements io.Reader
func (s Stream) Read(b []byte) (int, error) {
	if s.head == nil {
		return 0, io.EOF
	}
	return s.file.Read(b)
}

//Close implements  io.Closer
func (s Stream) Close() error {
	if s.head == nil {
		return nil
	}
	return s.file.Close()
}

//Size returns the expected size of the stream.
func (s Stream) Size() int64 {
	if s.head == nil {
		return 0
	}

	return s.head.Size
}
