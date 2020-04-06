package popup

import (
	"github.com/qlova/seed"
	"github.com/qlova/seed/script"
)

func init() {
	script.RegisterRenderer(func(c seed.Seed) []byte {
		return []byte(`
seed.show = async function(id) {
	//Collect arguments.
	let args = [];
	if (arguments.length > 1) {
		for (let i = 1; i < arguments.length; i++) {
			args.push(arguments[i]);
		}
	}

	let popup = seed.get(id);
	if (!popup) {
		console.error("seed.show: invalid popup ", id);
		return;
	}
	popup.template = popup.parent;

	popup.parent.parentElement.appendChild(popup);

	if (popup.onshow) await popup.onshow();
};

seed.hide = async function(id) {
	let popup = seed.get(id);
	if (!popup) {
		console.error("seed.show: invalid popup ", id);
		return;
	}

	if (popup.onhide) await popup.onhide();

	popup.template.content.appendChild(popup);
};`)
	})
}
