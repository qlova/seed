package gallery

import "github.com/qlova/seed"
import "github.com/qlova/seed/script"

func init() {
	seed.Embed("/photoswipe.js", []byte(Javascript))
	seed.Embed("/photoswipe.css", []byte(CSS))
	seed.Embed("/photoswipe-ui.js", []byte(UI))

	seed.Tail(`	
	<!-- Root element of PhotoSwipe. Must have class pswp. -->
	<div class="pswp" tabindex="-1" role="dialog" aria-hidden="true">

    <!-- Background of PhotoSwipe. 
         It's a separate element, as animating opacity is faster than rgba(). -->
    <div class="pswp__bg"></div>

    <!-- Slides wrapper with overflow:hidden. -->
    <div class="pswp__scroll-wrap">

        <!-- Container that holds slides. PhotoSwipe keeps only 3 slides in DOM to save memory. -->
        <div class="pswp__container">
            <!-- don't modify these 3 pswp__item elements, data is added later on -->
            <div class="pswp__item"></div>
            <div class="pswp__item"></div>
            <div class="pswp__item"></div>
        </div>

        <!-- Default (PhotoSwipeUI_Default) interface on top of sliding area. Can be changed. -->
        <div class="pswp__ui pswp__ui--hidden">

            <div class="pswp__top-bar">

                <!--  Controls are self-explanatory. Order can be changed. -->

                <div style="display:none;" class="pswp__counter"></div>

                <button style="display:none;" class="pswp__button pswp__button--close" title="Close (Esc)"></button>

                <button style="display:none;" class="pswp__button pswp__button--share" title="Share"></button>

                <button style="display:none;" class="pswp__button pswp__button--fs" title="Toggle fullscreen"></button>

                <button style="display:none;" class="pswp__button pswp__button--zoom" title="Zoom in/out"></button>

                <!-- Preloader demo https://codepen.io/dimsemenov/pen/yyBWoR -->
                <!-- element will get class pswp__preloader--active when preloader is running -->
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

//Returns a full-featured text editor with line numbers and optional syntax highlighting.
func New(images ...string) Widget {
	gallery := seed.New()

	gallery.Require("photoswipe.js")
	gallery.Require("photoswipe.css")
	gallery.Require("photoswipe-ui.js")
	
	gallery.OnReady(func(q seed.Script) {
		q.Javascript(gallery.Script(q).Element()+".items = [")
		for i, img := range images {
			q.Javascript(`{src:"`+img+`", w:900, h:600 }`)
			if i < len(images)-1 {
				q.Javascript(`,`)
			}
		}
		q.Javascript("];")
	})
	
	return Widget{gallery}
}

//Create a new Text widget and add it to the provided parent.
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
	widget.Q.Javascript(`ActivePhotoSwipe = new PhotoSwipe(document.querySelectorAll(".pswp")[0], PhotoSwipeUI_Default, `+widget.Element()+".items, {history:false}); ActivePhotoSwipe.init();")
}