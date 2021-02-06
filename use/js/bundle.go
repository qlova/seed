package js

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"qlova.org/seed/assets/inbed"
)

//Bundle downloads the JS from the given url or local file and bundles it as a go file.
func Bundle(name string, url string) error {
	var r io.Reader
	var err error

	if strings.HasPrefix(url, "http") {
		req, err := http.Get(url)
		if err != nil {
			return err
		}
		r = req.Body
	} else {
		r, err = os.Open(url)
		if err != nil {
			return err
		}
	}

	b, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	inbed.Bytes(name, b)

	return nil
}
