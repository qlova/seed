package inbed

import (
	"net/http"
	"os"
	"path/filepath"
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

//Open implements http.FileSyetem with inbed.Open
func (fs FileSystem) Open(name string) (http.File, error) {
	return Open(fs.Prefix + name)
}

//FileServer provides a http FileServer for serving embedded files.
var FileServer = http.FileServer(FileSystem{})
