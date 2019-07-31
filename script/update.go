package script

//Update is Javascript code for updating.
const Update = `
let iframe = null;

function pop_update() {
	iframe.style.visibility = "visible";
	window.requestAnimationFrame(function() {
		for (let i = 0; i < 4; i++) {
			let element = document.body.children[0];
			if (element != iframe) element.remove();
		}
	});
}

function update() {
	if ( window.location !== window.parent.location ) {
		window.parent.update();
		return;
	}

	window.localStorage.setItem("updating", "true");

	iframe = document.createElement("iframe");
	iframe.setAttribute("src", "./");
	iframe.setAttribute("width", "100%");
	iframe.setAttribute("height", "100%");
	iframe.style.visibility = "hidden";
	iframe.style.position = "absolute";
	iframe.style.border = "none";
	iframe.onload = function() {
		setTimeout(function() {
			pop_update();
		}, 500)
	}
	document.body.appendChild(iframe);
}
`

//Update seamlessly restarts and updates the app from the server.
func (q Script) Update() {
	q.Require(Update)
	q.Javascript(`setTimeout(update, 100);`)
}

//Restart peforms a hard reboot of the app.
func (q Script) Restart() {
	q.Javascript("window.location.reload();")
}
