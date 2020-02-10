//+build production

package inbed

import (
	"bytes"
	"compress/gzip"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"time"
)

const production = true

//File embeds the named file inside your code, relative to the location of the binary.
//If the file is a directory, the entire directory is recursively embedded.
func File(name string) {}

//Done should be called after all calls to File and before any calls to Open.
func Done() error {
	return nil
}

//Data is a low-level function (only available in production) for embedding data.
func Data(uri string, modtime int64, mode uint32, data []byte) {
	files[uri] = file{
		name: path.Base(uri),

		modTime: time.Unix(0, modtime),
		mode:    os.FileMode(mode),

		data: data,
	}
}

//Open opens a previously embedded file. If Done hasn't been called, it is called.
func Open(name string) (http.File, error) {
	asset, ok := files[name]
	if !ok {
		return nil, os.ErrNotExist
	}
	return asset.Open()
}

var files = make(map[string]file)

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
