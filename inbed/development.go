//+build !production

package inbed

import (
	"bufio"
	"compress/gzip"
	"encoding/gob"
	"encoding/hex"
	"fmt"

	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strconv"
	"unicode/utf8"

	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/css"
	"github.com/tdewolff/minify/html"
	"github.com/tdewolff/minify/js"
	"github.com/tdewolff/minify/json"
	"github.com/tdewolff/minify/svg"
	"github.com/tdewolff/minify/xml"
)

const production = false

var embeddings []string

//File embeds the named file inside your code, relative to the location of the binary.
//If the file is a directory, the entire directory is recursively embedded.
func File(name string) {
	embeddings = append(embeddings, name)
}

var done bool

func embeddingsEqual(otherEmbeddings []string) bool {
	if len(embeddings) != len(otherEmbeddings) {
		return false
	}
	for i, v := range embeddings {
		if v != otherEmbeddings[i] {
			return false
		}
	}
	return true
}

var mini = minify.New()

func init() {
	mini.AddFunc("text/css", css.Minify)
	mini.AddFunc("text/html", html.Minify)
	mini.AddFunc("image/svg+xml", svg.Minify)
	mini.AddFuncRegexp(regexp.MustCompile("^(application|text)/(x-)?(java|ecma)script$"), js.Minify)
	mini.AddFuncRegexp(regexp.MustCompile("[/+]json$"), json.Minify)
	mini.AddFuncRegexp(regexp.MustCompile("[/+]xml$"), xml.Minify)
	mini.AddFunc("encoding/gzip", func(m *minify.M, w io.Writer, r io.Reader, _ map[string]string) error {
		gw := gzip.NewWriter(w)
		_, err := io.Copy(gw, r)
		if err != nil {
			return fmt.Errorf("could not gzip stream: %w", err)
		}
		gw.Close()
		return nil
	})
	/*mini.AddFunc("image/png", func(m *minify.M, w io.Writer, r io.Reader, _ map[string]string) error {
		img, _, err := image.Decode(r)
		if err != nil {
			return fmt.Errorf("could not decode image: %w", err)
		}

		img = lossypng.Compress(img, lossypng.RGBAConversion, 20)

		return (&png.Encoder{
			CompressionLevel: png.BestCompression,
		}).Encode(w, img)
	})*/
}

func embedFile(name string, w *os.File, r *os.File) error {
	info, err := r.Stat()
	if err != nil {
		return fmt.Errorf("could not stat file: %w", err)
	}

	var reader io.Reader
	var writer io.Writer

	switch path.Ext(info.Name()) {
	case ".html":
		reader = mini.Reader("text/html", r)
	case ".svg":
		reader = mini.Reader("image/svg+xml", r)
	case ".css":
		reader = mini.Reader("text/css", r)
	case ".js":
		reader = mini.Reader("application/javascript", r)
	case ".json":
		reader = mini.Reader("text/json", r)
	case ".xml":
		reader = mini.Reader("text/xml", r)
	/*case ".png":
	reader = mini.Reader("image/png", r)*/
	default:
		reader = bufio.NewReader(r)
	}

	reader = bufio.NewReader(mini.Reader("encoding/gzip", reader))

	writer = bufio.NewWriter(w)

	info, err = w.Stat()
	if err != nil {
		return fmt.Errorf("could not stat file: %w", err)
	}

	if _, err := w.WriteString(fmt.Sprintf(`	inbed.Data(%q, %v, %v, []byte("`,
		name, info.ModTime().UnixNano(), uint32(info.Mode()))); err != nil {

		return fmt.Errorf("could not write assets file: %w", err)
	}

	for {
		reader := reader.(*bufio.Reader)
		writer := writer.(*bufio.Writer)

		peek, err := reader.Peek(4)
		if err != nil && len(peek) == 0 {
			if err == io.EOF {
				break
			}
			return fmt.Errorf("could not read file: %w", err)
		}

		if len(peek) == 0 {
			break
		}

		char, length := utf8.DecodeRune(peek)
		if char == utf8.RuneError {
			var hexed = `\x` + hex.EncodeToString(peek[:1])

			if _, err := writer.WriteString(hexed); err != nil {
				return fmt.Errorf("could not write file: %w", err)
			}

			if _, err := reader.Discard(1); err != nil {
				return fmt.Errorf("could not read file: %w", err)
			}

			if len(peek) == 1 && err == io.EOF {
				break
			}

			continue
		}

		if _, err := reader.Discard(length); err != nil {
			return fmt.Errorf("could not read file: %w", err)
		}

		if char == '\\' {
			if _, err := writer.WriteString(`\x5c`); err != nil {
				return fmt.Errorf("could not write file: %w", err)
			}
		} else if char == '"' {
			if _, err := writer.WriteString(`\x22`); err != nil {
				return fmt.Errorf("could not write file: %w", err)
			}
		} else if char == '\'' {
			if _, err := writer.WriteString(`'`); err != nil {
				return fmt.Errorf("could not write file: %w", err)
			}
		} else if char == 0 {
			if _, err := writer.WriteString(`\x00`); err != nil {
				return fmt.Errorf("could not write file: %w", err)
			}
		} else {
			var q = strconv.QuoteRune(char)

			if _, err := writer.WriteString(q[1 : len(q)-1]); err != nil {
				return fmt.Errorf("could not write file: %w", err)
			}
		}

	}

	if _, err := writer.(*bufio.Writer).WriteString(`"))` + "\n"); err != nil {
		return fmt.Errorf("could not write file: %w", err)
	}

	if err := writer.(*bufio.Writer).Flush(); err != nil {
		return fmt.Errorf("could not write file: %w", err)
	}

	return nil
}

func buildEmbeddings() error {
	assets, err := os.Create(filepath.Join(Root, PackageName, PackageName+".go"))
	if err != nil {
		return fmt.Errorf("could not create inbed package file: %w", err)
	}

	if _, err := assets.WriteString(`//+build production

package ` + PackageName + `

import "github.com/qlova/seed/inbed"

func init() {
`); err != nil {
		return fmt.Errorf("could not write inbed package header: %w", err)
	}

	for _, embedding := range embeddings {

		file, err := os.Open(embedding)
		if err != nil {
			return fmt.Errorf("could not embed file %v: %w", embedding, err)
		}

		if stat, err := file.Stat(); err == nil && stat.IsDir() {

			if err := filepath.Walk(embedding, func(path string, info os.FileInfo, err error) error {
				if info.Name() == PackageName+".go" || info.Name() == "cache.gob" {
					return nil
				}

				if err != nil {
					return nil
				}

				if info.IsDir() {
					if _, err := assets.WriteString(fmt.Sprintf(`	inbed.Data(%q, %v, %v, nil)`+"\n",
						path, info.ModTime().UnixNano(), uint32(info.Mode()))); err != nil {

						return fmt.Errorf("could not write assets file: %w", err)
					}
					return nil
				}

				data, err := os.Open(path)
				if err != nil {
					return fmt.Errorf("could not open file for embedding %v: %w", path, err)
				}

				if err := embedFile(path, assets, data); err != nil {
					return fmt.Errorf("could not embed file %v: %w", path, err)
				}

				if err := data.Close(); err != nil {
					return fmt.Errorf("could not close asset %v: %w", path, err)
				}

				return nil

			}); err != nil {
				return fmt.Errorf("could not embed directory %v: %w", embedding, err)
			}

		} else if err == nil {
			if err := embedFile(embedding, assets, file); err != nil {
				return fmt.Errorf("could not embed file %v: %w", embedding, err)
			}

		} else if err != nil {
			return fmt.Errorf("could not stat embedded file %v: %w", embedding, err)
		}

		if err := file.Close(); err != nil {
			return fmt.Errorf("could not close embedded file %v: %w", embedding, err)
		}
	}

	if _, err := assets.WriteString(`}`); err != nil {
		return fmt.Errorf("could not write asset file footer: %w", err)
	}

	if err := assets.Close(); err != nil {
		return fmt.Errorf("could not close assets file: %w", err)
	}

	//Update cache.
	if cache, err := os.Create(filepath.Join(Root, PackageName, "cache.gob")); err == nil {

		if err := gob.NewEncoder(cache).Encode(embeddings); err != nil {
			return fmt.Errorf("could not create embedding cache: %w", err)
		}
		cache.Close()
	}

	return nil
}

//Done should be called after all calls to File and before any calls to Open.
func Done() error {
	if done {
		return nil
	}
	done = true

	//Create an inbed.go file in the project root
	if _, err := os.Stat(filepath.Join(Root, ImporterName)); os.IsNotExist(err) {
		file, err := os.Create(filepath.Join(Root, ImporterName))
		if err != nil {
			return fmt.Errorf("could not create %v file: %w", ImporterName, err)
		}

		if _, err := file.WriteString(`//+build production

package main

import _ "./` + PackageName + `"
`); err != nil {
			return fmt.Errorf("could not write %v file: %w", ImporterName, err)
		}

		if err := file.Close(); err != nil {
			return fmt.Errorf("could not close %v file: %w", ImporterName, err)
		}
	}

	//Create an inbed package directory.
	if info, err := os.Stat(filepath.Join(Root, PackageName)); os.IsNotExist(err) {
		if err := os.Mkdir(filepath.Join(Root, PackageName), os.ModePerm); err != nil {
			return fmt.Errorf("could not create %v directory: %w", PackageName, err)
		}
	} else if err == nil && !info.IsDir() {
		return fmt.Errorf("%v is not a directory", PackageName)
	}

	inbedInfo, err := os.Stat(filepath.Join(Root, PackageName, PackageName+".go"))
	if err == nil {
		var lastInbedTime = inbedInfo.ModTime()

		for _, embedding := range embeddings {
			info, err := os.Stat(filepath.Join(Root, embedding))
			if err != nil {
				return fmt.Errorf("could not stat embedding %v: %w", embedding, err)
			}
			if info.ModTime().After(lastInbedTime) {
				return buildEmbeddings()
			}
		}

		//Try the cache.
		if cache, err := os.Open(filepath.Join(Root, PackageName, "cache.gob")); err == nil {
			var oldEmbeddings []string
			if err := gob.NewDecoder(cache).Decode(&oldEmbeddings); err == nil {
				if !embeddingsEqual(oldEmbeddings) {
					return buildEmbeddings()
				}
			}
			cache.Close()
		}

		return nil
	}

	return buildEmbeddings()
}

//Open opens a previously embedded file. If Done hasn't been called, it is called.
func Open(name string) (http.File, error) {
	if !done {
		if err := Done(); err != nil {
			return nil, err
		}
	}

	return os.Open(filepath.Join(Root, name))
}
