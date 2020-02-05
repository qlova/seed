package seed

import (
	"bufio"
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

//Asset is a static resource that is needed by an application, this may include images, audio, video, documents etc.
type Asset struct {
	path          string
	cache, bundle bool
}

//NewAsset creates a new cached image at the given path.
//This asset will be bundled with the app and cached.
func NewAsset(path string) Asset {
	return Asset{
		path:  path,
		cache: true,
	}
}

//AddTo adds an asset to a seed.
func (asset Asset) AddTo(seed Interface) {
	seed.Root().assets = append(seed.Root().assets, asset)
}

type embeddedAsset struct {
	name string

	size    int64
	modTime time.Time
	mode    os.FileMode

	data []byte

	*bytes.Reader
}

func (asset embeddedAsset) Open() embeddedAsset {
	asset.Reader = bytes.NewReader(asset.data)
	return asset
}

func (asset embeddedAsset) Readdir(int) ([]os.FileInfo, error) {
	return nil, errors.New("permission denied")
}

func (asset embeddedAsset) Stat() (os.FileInfo, error) {
	return asset, nil
}

func (asset embeddedAsset) Name() string {
	return asset.name
}

func (asset embeddedAsset) Size() int64 {
	return asset.size
}

func (asset embeddedAsset) Mode() os.FileMode {
	return asset.mode
}

func (asset embeddedAsset) ModTime() time.Time {
	return asset.modTime
}

func (asset embeddedAsset) IsDir() bool {
	return asset.mode.IsDir()
}

func (asset embeddedAsset) Sys() interface{} {
	return asset.data
}

func (asset embeddedAsset) Close() error {
	return nil
}

var embeddedAssets = make(map[string]embeddedAsset)

//EmbedAsset embeds an asset.
func EmbedAsset(pathToAsset string, size int64, modtime int64, mode uint32, data []byte) {
	embeddedAssets[pathToAsset] = embeddedAsset{
		name: path.Base(pathToAsset),
		size: size,

		modTime: time.Unix(0, modtime),
		mode:    os.FileMode(mode),

		data: data,
	}
}

type embeddedFileSystem struct{}

func (fs embeddedFileSystem) Open(name string) (http.File, error) {
	asset, ok := embeddedAssets[name]
	if !ok {
		return nil, os.ErrNotExist
	}
	return asset.Open(), nil
}

var embeddedFileServer = http.FileServer(embeddedFileSystem{})

func (app *App) embedAssets() error {
	if Live {
		return nil
	}

	if _, err := os.Stat(Dir + "/production.go"); os.IsNotExist(err) {
		file, err := os.Create(Dir + "/production.go")
		if err != nil {
			return fmt.Errorf("could not create production file: %w", err)
		}

		if _, err := file.WriteString(`//+build production

package main

import _ "./assets"
`); err != nil {
			return fmt.Errorf("could not write production file: %w", err)
		}

		if err := file.Close(); err != nil {
			return fmt.Errorf("could not close production file: %w", err)
		}
	}

	assets, err := os.Create(Dir + "/assets/assets.go")
	if err != nil {
		return fmt.Errorf("could not create asset file: %w", err)
	}

	if _, err := assets.WriteString(`//+build production

package assets

import "github.com/qlova/seed"

func init() {
`); err != nil {
		return fmt.Errorf("could not write asset file header: %w", err)
	}

	if err := filepath.Walk(Dir+"/assets", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}

		if path == "./assets" {
			return nil
		}

		if info.Name() == "assets.go" {
			return nil
		}

		if info.IsDir() {
			if _, err := assets.WriteString(fmt.Sprintf(`	seed.EmbedAsset(%q, %v, %v, %v, nil)`+"\n",
				strings.TrimPrefix(path, "assets"), info.Size(), info.ModTime().UnixNano(), uint32(info.Mode()))); err != nil {

				return fmt.Errorf("could not write assets file: %w", err)
			}
			return nil
		}

		if _, err := assets.WriteString(fmt.Sprintf(`	seed.EmbedAsset(%q, %v, %v, %v, []byte("`,
			strings.TrimPrefix(path, "assets"), info.Size(), info.ModTime().UnixNano(), uint32(info.Mode()))); err != nil {

			return fmt.Errorf("could not write assets file: %w", err)
		}

		data, err := os.Open(path)
		if err != nil {
			return fmt.Errorf("could not open asset %v: %w", path, err)
		}

		reader, writer := bufio.NewReader(data), bufio.NewWriter(assets)

		for {
			peek, err := reader.Peek(4)
			if err != nil && len(peek) == 0 {
				if err == io.EOF {
					break
				}
				return fmt.Errorf("could not read asset %v: %w", path, err)
			}

			if len(peek) == 0 {
				break
			}

			char, length := utf8.DecodeRune(peek)
			if char == utf8.RuneError {
				var hexed = `\x` + hex.EncodeToString(peek[:1])

				if _, err := writer.WriteString(hexed); err != nil {
					return fmt.Errorf("could not write assets file: %w", err)
				}

				if _, err := reader.Discard(1); err != nil {
					return fmt.Errorf("could not read asset %v: %w", path, err)
				}

				if len(peek) == 1 && err == io.EOF {
					break
				}

				continue
			}

			if _, err := reader.Discard(length); err != nil {
				return fmt.Errorf("could not read asset %v: %w", path, err)
			}

			if char == '\\' {
				if _, err := writer.WriteString(`\x5c`); err != nil {
					return fmt.Errorf("could not write assets file: %w", err)
				}
			} else if char == '"' {
				if _, err := writer.WriteString(`\x22`); err != nil {
					return fmt.Errorf("could not write assets file: %w", err)
				}
			} else if char == '\'' {
				if _, err := writer.WriteString(`'`); err != nil {
					return fmt.Errorf("could not write assets file: %w", err)
				}
			} else if char == 0 {
				if _, err := writer.WriteString(`\x00`); err != nil {
					return fmt.Errorf("could not write assets file: %w", err)
				}
			} else {
				var q = strconv.QuoteRune(char)

				if _, err := writer.WriteString(q[1 : len(q)-1]); err != nil {
					return fmt.Errorf("could not write assets file: %w", err)
				}
			}

		}

		if err := writer.Flush(); err != nil {
			return fmt.Errorf("could not write assets file: %w", err)
		}

		if _, err := assets.WriteString(`"))` + "\n"); err != nil {
			return fmt.Errorf("could not write assets file: %w", err)
		}

		if err := data.Close(); err != nil {
			return fmt.Errorf("could not close asset %v: %w", path, err)
		}

		return nil
	}); err != nil {
		return fmt.Errorf("could not generate assets file: %w", err)
	}

	if _, err := assets.WriteString(`}`); err != nil {
		return fmt.Errorf("could not write asset file footer: %w", err)
	}

	if err := assets.Close(); err != nil {
		return fmt.Errorf("could not close assets file: %w", err)
	}

	return nil
}
