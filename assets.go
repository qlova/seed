package seed

// An asset is a static resource that is needed by an application, this may include images, audio, video, documents etc.
type Asset struct {
	path          string
	cache, bundle bool
}

//Create a new cached image at the given path.
func NewAsset(path string) Asset {
	return Asset{
		path:  path,
		cache: true,
	}
}

//Add an asset to a seed.
func (asset Asset) AddTo(seed Interface) {
	seed.Root().assets = append(seed.Root().assets, asset)
}
