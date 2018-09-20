package manifest

import "encoding/json"

type Icon struct {
	Source string `json:"src"`
	Sizes string `json:"sizes"`
	Type string `json:"type"`
}

type Manifest struct {
	Name string 			`json:"name"`
	ShortName string 		`json:"short_name"`
	StartUrl string 		`json:"start_url"`
	Display string			`json:"display"`
	BackgroundColor string 	`json:"background_color"`
	Description string 		`json:"description"`
	
	Icons []Icon			`json:"icons"`
}

func New() Manifest {
	var manifest Manifest
	manifest.StartUrl = "."
	manifest.Display = "standalone"
	manifest.BackgroundColor = "#fff"
	return manifest
}

func (manifest Manifest) Render() []byte {
	var result, _ = json.Marshal(manifest)
	return result
}
