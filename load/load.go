package load

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/mholt/archiver"
	"github.com/qlova/seed"
)

//File scans a Pencil file and returns it as an app.
func File(name string) *seed.App {
	var file, err = os.Open(name)
	if err != nil {
		//TODO error handling
		fmt.Println(err)
		os.Exit(1)
	}

	var decompressor = archiver.NewTarGz()
	err = decompressor.Open(file, 0)
	if err != nil {
		//TODO error handling
		fmt.Println(err)
		os.Exit(1)
	}

	var project = make(map[string][]byte)

	for {
		f, err := decompressor.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			//TODO error handling
			fmt.Println(err)
			os.Exit(1)
		}

		project[f.Name()], err = ioutil.ReadAll(f)

		if err != nil {
			//TODO error handling
			fmt.Println(err)
			os.Exit(1)
		}

		err = f.Close()
		if err != nil {
			//TODO error handling
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println(f.Name())
	}

	os.Exit(0)

	return seed.NewApp(name)
}
