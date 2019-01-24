package seed

import (
	"net/http"
	"path"
)

type embedding struct {
	ContentType string
	Data []byte
}

var embeddings = make(map[string]embedding)

func Embed(name string, data []byte) {
	var ContentType string

	if path.Ext(name) == ".js" {
		ContentType = "application/javascript"
	}
	if path.Ext(name) == ".css" {
		ContentType = "text/css"
	}		
	
	embeddings[name] = embedding{
		ContentType: ContentType,
		Data: data,
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
