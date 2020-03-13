package inbed

import (
	"bytes"
	"net/http"
	"os"
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

	if b, ok := info.Sys().([]byte); ok {
		http.ServeContent(w, r, info.Name(), info.ModTime(), bytes.NewReader(b))
	} else {
		http.ServeContent(w, r, info.Name(), info.ModTime(), f)
	}
}
