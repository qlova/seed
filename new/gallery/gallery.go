package gallery

import (
	"fmt"
	"image"
	"os"

	"qlova.org/seed"
	"qlova.org/seed/assets/inbed"
	"qlova.org/seed/client"
	seed_image "qlova.org/seed/new/image"
	"qlova.org/seed/use/css"
	"qlova.org/seed/use/js"
)

func getImageDimension(assetPath string) string {
	file, err := inbed.Open(assetPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return fmt.Sprint("w:", "0", ",h:", "0")
	}

	image, _, err := image.DecodeConfig(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", assetPath, err)
	}

	file.Close()

	return fmt.Sprint("w:", image.Width, ",h:", image.Height)
}

type data []string

//Set the images that the gallery should display.
func Set(images ...string) seed.Option {
	return seed.Mutate(func(d *data) {
		*d = images
	})
}

//New returns a new gallery that can view images from the assets folder.
func New(options ...seed.Option) seed.Seed {
	Gallery := seed_image.New(
		js.Require("/assets/js/photoswipe.js", ""),
		js.Require("/assets/js/photoswipe-ui.js", ""),
		css.Require("/assets/css/photoswipe.css", ""),
		css.Require("/assets/photoswipe/default-skin.css", ""),

		seed.Options(options),
	)

	var images data
	Gallery.Load(&images)

	if len(images) > 0 {
		Gallery.With(
			seed_image.Set(images[0]),

			seed.NewOption(func(c seed.Seed) {
				client.OnLoad(js.Script(func(q js.Ctx) {
					fmt.Fprintf(q, client.Element(c)+".items = [")
					for i, img := range images {

						var dimensions = getImageDimension(img)

						fmt.Fprintf(q, `{src:"`+img+`", `+dimensions+` }`)
						if i < len(images)-1 {
							fmt.Fprintf(q, `,`)
						}
					}
					fmt.Fprintf(q, "];")
				})).AddTo(Gallery)

				client.OnClick(js.Script(func(q js.Ctx) {
					fmt.Fprintf(q, `if (document.querySelectorAll(".pswp").length == 0) {
						document.body.insertAdjacentHTML("beforeend", '<div class="pswp" tabindex="-1" role="dialog" aria-hidden="true"><!-- Background of PhotoSwipe. Its a separate element as animating opacity is faster than rgba(). --> <div class="pswp__bg"></div><div class="pswp__scroll-wrap"><!-- Container that holds slides. PhotoSwipe keeps only 3 of them in the DOM to save memory. Dont modify these 3 pswp__item elements, data is added later on. --> <div class="pswp__container"> <div class="pswp__item"></div><div class="pswp__item"></div><div class="pswp__item"></div></div><div class="pswp__ui pswp__ui--hidden"> <div class="pswp__top-bar"> <div class="pswp__counter"></div><button class="pswp__button pswp__button--close" title="Close (Esc)"></button><button class="pswp__button pswp__button--fs" title="Toggle fullscreen"></button> <button class="pswp__button pswp__button--zoom" title="Zoom in/out"></button> <div class="pswp__preloader"> <div class="pswp__preloader__icn"> <div class="pswp__preloader__cut"> <div class="pswp__preloader__donut"></div></div></div></div></div><div class="pswp__caption"> <div class="pswp__caption__center"></div></div></div></div></div>');
					}`)
					fmt.Fprintf(q, `ActivePhotoSwipe = new PhotoSwipe(document.querySelectorAll(".pswp")[0], PhotoSwipeUI_Default, `+client.Element(c)+".items, {history:false}); ActivePhotoSwipe.init(); ActivePhotoSwipe.listen('close', function() { ActivePhotoSwipe = null; });")
				})).AddTo(Gallery)
			}),
		)
	}

	return Gallery
}
