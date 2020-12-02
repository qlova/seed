//+build bundle

package inbed

import (
	"net/http"
	"os"
	"strings"
)

const production = true

//File embeds the named file inside your code, relative to the location of the binary.
//If the file is a directory, the entire directory is recursively embedded.
func File(name string) {}

//Done should be called after all calls to File and before any calls to Open.
func Done() error {
	return nil
}

//Open opens a previously embedded file. If Done hasn't been called, it is called.
func Open(name string) (http.File, error) {
	if len(name) > 0 && name[0] == '/' {
		name = name[1:]
	}

	asset, ok := files[name]
	if !ok {
		return nil, os.ErrNotExist
	}
	return asset.Open()
}

//List lists the paths of any files starting with the prefix 'folder'
func List(folder string) []string {
	var result []string

	folder = strings.TrimPrefix(folder, "./")

	for f := range files {
		if strings.HasPrefix(f, folder) {
			result = append(result, f)
		}
	}

	return result
}
