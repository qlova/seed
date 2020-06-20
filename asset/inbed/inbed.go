package inbed

import (
	"bytes"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
)

//Root is the location of the project (defaults to the directory the binary is running from).
var Root = filepath.Dir(os.Args[0])

//PackageName is the name of the generated package.
var PackageName = "inbed"

//ImporterName is the name of the go source file that imports the generated inbed package.
var ImporterName = "inbed.go"

//FileSystem implements net/http.FileSystem
type FileSystem struct {
	Prefix string
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

	if production {
		w.Header().Set("Content-Encoding", "gzip")
	}

	if path.Ext(info.Name()) == ".wasm" {
		w.Header().Set("Content-Type", "application/wasm")
	}

	if b, ok := info.Sys().([]byte); ok {
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

	if production {
		w.Header().Set("Content-Encoding", "gzip")
	}

	if path.Ext(info.Name()) == ".wasm" {
		w.Header().Set("Content-Type", "application/wasm")
	}

	if b, ok := info.Sys().([]byte); ok {
		http.ServeContent(w, r, info.Name(), info.ModTime(), bytes.NewReader(b))
	} else {
		http.ServeContent(w, r, info.Name(), info.ModTime(), f)
	}
}
