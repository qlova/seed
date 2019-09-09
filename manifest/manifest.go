package manifest

import (
	"encoding/json"
	"fmt"
	"image/color"
)

//Icon is an app-manifest icon.
type Icon struct {
	Source string `json:"src"`
	Sizes  string `json:"sizes"`
	Type   string `json:"type,omitempty"`
}

//Manifest is a webapp manifest.
type Manifest struct {
	Name            string `json:"name"`
	ShortName       string `json:"short_name"`
	StartURL        string `json:"start_url"`
	Display         string `json:"display"`
	BackgroundColor string `json:"background_color"`
	Description     string `json:"description"`
	ThemeColor      string `json:"theme_color"`

	Icons []Icon `json:"icons"`
}

//New returns a new webapp manifest.
func New() Manifest {
	var manifest Manifest
	manifest.StartURL = "."
	manifest.Display = "standalone"
	manifest.BackgroundColor = "#ffffff"
	manifest.ThemeColor = "#ffffff"
	manifest.Icons = []Icon{}
	return manifest
}

//Render returns a manifest as json encoded bytes.
func (manifest Manifest) Render() []byte {
	var result, _ = json.Marshal(manifest)
	return result
}

//SetName sets the name of this application.
func (manifest *Manifest) SetName(name string) {
	manifest.Name = name
	if manifest.ShortName == "" {
		manifest.ShortName = name
	}
}

//SetShortName set the short name of the application.
func (manifest *Manifest) SetShortName(name string) {
	manifest.ShortName = name
}

//SetDescription sets the description of the application.
func (manifest *Manifest) SetDescription(description string) {
	manifest.Description = description
}

//SetIcon sets the icon for the application to be the image at the given path.
func (manifest *Manifest) SetIcon(path string) {
	var icon Icon
	icon.Source = path
	icon.Sizes = "192x192"

	manifest.Icons = append(manifest.Icons, icon)
}

//SetThemeColor sets the theme color of the application.
func (manifest *Manifest) SetThemeColor(c color.Color) {
	rgba := color.RGBAModel.Convert(c).(color.RGBA)
	manifest.ThemeColor = fmt.Sprintf("#%.2x%.2x%.2x", rgba.R, rgba.G, rgba.B)
}
