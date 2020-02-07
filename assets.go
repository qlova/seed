package seed

import "github.com/qlova/seed/inbed"

func init() {
	inbed.PackageName = "assets"
	inbed.ImporterName = "production.go"
}

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
