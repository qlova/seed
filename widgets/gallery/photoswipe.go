package gallery

import "github.com/qlova/seed"
import "github.com/qlova/seed/script"

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

func init() {
	seed.Embed("/photoswipe.js", []byte(Javascript))
	seed.Embed("/photoswipe.css", []byte(CSS))
	seed.Embed("/photoswipe-ui.js", []byte(UI))

	seed.Tail(`	
	<div class="pswp" tabindex="-1" role="dialog" aria-hidden="true">
    <div class="pswp__bg"></div>

    <div class="pswp__scroll-wrap">


        <div class="pswp__container">
            <div class="pswp__item"></div>
            <div class="pswp__item"></div>
            <div class="pswp__item"></div>
        </div>


        <div class="pswp__ui pswp__ui--hidden">

            <div class="pswp__top-bar">

                <!--  Controls are hidden. -->
                <div style="display:none;" class="pswp__counter"></div>
                <button style="display:none;" class="pswp__button pswp__button--close" title="Close (Esc)"></button>
                <button style="display:none;" class="pswp__button pswp__button--share" title="Share"></button>
                <button style="display:none;" class="pswp__button pswp__button--fs" title="Toggle fullscreen"></button>
                <button style="display:none;" class="pswp__button pswp__button--zoom" title="Zoom in/out"></button>

                <div class="pswp__preloader">
                    <div class="pswp__preloader__icn">
                      <div class="pswp__preloader__cut">
                        <div class="pswp__preloader__donut"></div>
                      </div>
                    </div>
                </div>
            </div>

            <div class="pswp__share-modal pswp__share-modal--hidden pswp__single-tap">
                <div class="pswp__share-tooltip"></div> 
            </div>

            <button style="display:none;" class="pswp__button pswp__button--arrow--left" title="Previous (arrow left)">
            </button>

            <button style="display:none;" class="pswp__button pswp__button--arrow--right" title="Next (arrow right)">
            </button>

            <div class="pswp__caption">
                <div class="pswp__caption__center"></div>
            </div>

          </div>

        </div>
    </div>
	`)
}

type Widget struct {
	seed.Seed
}

func getImageDimension(imagePath string) string {
	file, err := os.Open(imagePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}

	image, _, err := image.DecodeConfig(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", imagePath, err)
	}

	file.Close()

	return fmt.Sprint("w:", image.Width, ",h:", image.Height)
}

//Returns gallery that displays 'local' images (in the assets directory).
func New(images ...string) Widget {
	gallery := seed.New()

	gallery.Require("photoswipe.js")
	gallery.Require("photoswipe.css")
	gallery.Require("photoswipe-ui.js")

	gallery.OnReady(func(q seed.Script) {
		q.Javascript(gallery.Script(q).Element() + ".items = [")
		for i, img := range images {

			var dimensions = getImageDimension(seed.Dir + "/assets/" + img)

			q.Javascript(`{src:"` + img + `", ` + dimensions + ` }`)
			if i < len(images)-1 {
				q.Javascript(`,`)
			}
		}
		q.Javascript("];")
	})

	return Widget{gallery}
}

func AddTo(parent seed.Interface, images ...string) Widget {
	var Gallery = New(images...)
	parent.Root().Add(Gallery)
	return Gallery
}

type Script struct {
	script.Seed
}

func (w Widget) Script(q script.Script) Script {
	return Script{w.Seed.Script(q)}
}

func (widget Script) Open() {
	widget.Q.Javascript(`ActivePhotoSwipe = new PhotoSwipe(document.querySelectorAll(".pswp")[0], PhotoSwipeUI_Default, ` + widget.Element() + ".items, {history:false}); ActivePhotoSwipe.init();")
}
