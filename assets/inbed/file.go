package inbed

import (
	"bytes"
	"compress/gzip"
	"errors"
	"io/ioutil"
	"os"
	"time"
)

type file struct {
	name string

	modTime time.Time
	mode    os.FileMode

	data []byte

	*bytes.Reader
}

func (f file) Open() (file, error) {
	gr, err := gzip.NewReader(bytes.NewReader(f.data))
	if err != nil {
		return f, err
	}
	b, err := ioutil.ReadAll(gr)
	if err != nil {
		return f, err
	}
	f.Reader = bytes.NewReader(b)
	return f, err
}

func (f file) Readdir(int) ([]os.FileInfo, error) {
	return nil, errors.New("permission denied")
}

func (f file) Stat() (os.FileInfo, error) {
	return f, nil
}

func (f file) Name() string {
	return f.name
}

func (f file) Size() int64 {
	return int64(len(f.data))
}

func (f file) Mode() os.FileMode {
	return f.mode
}

func (f file) ModTime() time.Time {
	return f.modTime
}

func (f file) IsDir() bool {
	return f.mode.IsDir()
}

func (f file) Sys() interface{} {
	return f.data
}

func (f file) Close() error {
	return nil
}
