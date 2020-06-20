package transition

import (
	"qlova.org/seed"
	"qlova.org/seed/script"
)

func init() {
	script.RegisterRenderer(func(c seed.Seed) []byte {
		return []byte(`
		seed.in = function(element, duration) {
			if (element.classList.contains("page")) {
				if (!seed.LastPage) return;
			}
		
			element.style.animationFillMode = "forwards";
			element.style.animationDuration = duration+"s";
			element.style.animationIterationCount = 1;
			element.style.zIndex = 50;
			element.style.position = "absolute";
		
			seed.goto.in = new Promise(resolve => {
				setTimeout(function() {
					element.style.animation = "";
					element.style.zIndex = "";
					element.style.position = "";
					resolve()
				}, duration*1000);
			});
		}
		seed.out = function(element, duration) {
			if (element.classList.contains("page")) {
				if (!seed.LastPage) return;
			}
		
			element.style.animationFillMode = "forwards";
			element.style.animationDuration = duration+"s";
			element.style.animationIterationCount = 1;
			element.style.zIndex = 50;
		
			seed.goto.out = new Promise(resolve => {
				setTimeout(function() {
					element.style.animation = "";
					element.style.zIndex = "";
					resolve()
				}, duration*1000);
			});
		}
		`)
	})
}
