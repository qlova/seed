package inbed

import (
	"bytes"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

//Root is the location of the project (defaults to the directory the binary is running from).
var Root = filepath.Dir(os.Args[0])

//PackageName is the name of the generated package.
//If the package name has a .go extension then a file is produced instead of a package.
var PackageName = "inbed"

//SingleFile is the filename of the file to embed into, if blank, a package is used instead.
var SingleFile string

//ImporterName is the name of the go source file that imports the generated inbed package.
var ImporterName = "inbed.go"

//FileSystem implements net/http.FileSystem
type FileSystem struct {
	Prefix string
}

var memory = make(map[string][]byte)

//Bytes allows inbedding of bytes.
func Bytes(path string, data []byte) {
	memory[path] = data
}

var files = make(map[string]file)

//Data is a low-level function (only available in production) for embedding data.
func Data(uri string, modtime int64, mode uint32, data []byte) {
	files[uri] = file{
		name: path.Base(uri),

		modTime: time.Unix(0, modtime),
		mode:    os.FileMode(mode),

		data: data,
	}
}

//ServeFile mimics http.ServeFile
func ServeFile(w http.ResponseWriter, r *http.Request, name string) {
	f, err := Open(name)
	if err != nil {
		w.WriteHeader(404)
		return
	}

	info, err := f.Stat()
	if err != nil {
		w.WriteHeader(500)
		return
	}

	if path.Ext(info.Name()) == ".wasm" {
		w.Header().Set("Content-Type", "application/wasm")
	}

	if b, ok := info.Sys().([]byte); ok {

		_, production := info.(file)

		if production {
			w.Header().Set("Content-Encoding", "gzip")
		}

		http.ServeContent(w, r, info.Name(), info.ModTime(), bytes.NewReader(b))
	} else {
		http.ServeContent(w, r, info.Name(), info.ModTime(), f)
	}
}

//Open implements http.FileSystem with inbed.Open
func (fs FileSystem) Open(name string) (http.File, error) {
	return Open(fs.Prefix + name)
}

func (fs FileSystem) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	upath := r.URL.Path
	if !strings.HasPrefix(upath, "/") {
		upath = "/" + upath
		r.URL.Path = upath
	}

	f, err := fs.Open(upath)
	if err != nil {
		w.WriteHeader(404)
		return
	}

	info, err := f.Stat()
	if err != nil {
		w.WriteHeader(500)
		return
	}

	if path.Ext(info.Name()) == ".wasm" {
		w.Header().Set("Content-Type", "application/wasm")
	}

	if b, ok := info.Sys().([]byte); ok {

		_, production := info.(file)
		if production {
			w.Header().Set("Content-Encoding", "gzip")
		}

		http.ServeContent(w, r, info.Name(), info.ModTime(), bytes.NewReader(b))
	} else {
		http.ServeContent(w, r, info.Name(), info.ModTime(), f)
	}
}
