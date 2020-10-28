package js

import (
	"io/ioutil"
	"net/http"

	"qlova.org/seed/asset/inbed"
)

//Bundle downloads the JS from the given url or local file and bundles it as a go file.
func Bundle(name string, url string) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	inbed.Bytes(name, b)

	return nil
}
