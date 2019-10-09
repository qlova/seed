package seed

import (
	"net/http"
	"path"
)

type embedding struct {
	ContentType string
	Data        []byte
}

//GetEmbeddings is an internal function.
func GetEmbeddings() map[string]embedding {
	return embeddings
}

var embeddings = make(map[string]embedding)

//Embed embeds an asset with the specified name and data.
func Embed(name string, data []byte) {
	var ContentType string

	if path.Ext(name) == ".js" {
		ContentType = "application/javascript"
	}
	if path.Ext(name) == ".css" {
		ContentType = "text/css"
	}
	if path.Ext(name) == ".wasm" {
		ContentType = "application/wasm"
	}

	embeddings[name] = embedding{
		ContentType: ContentType,
		Data:        data,
	}
}

func embedded(w http.ResponseWriter, r *http.Request) bool {
	if embedding, ok := embeddings[r.URL.Path]; ok {
		w.Header().Set("Content-Type", embedding.ContentType)
		w.Write(embedding.Data)
		return true
	}

	return false
}
