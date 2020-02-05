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
func (q Ctx) Update() {
	q.Require(Update)
	q.Javascript(`setTimeout(update, 100);`)
}

//CheckForUpdates checks for an update.
func (q Ctx) CheckForUpdates() {
	q.Javascript(`ServiceWorker_Registration.update();`)
}

//Restart peforms a restart of the app.
func (q Ctx) Restart() {
	q.Javascript("window.location.reload();")
}

//HardReset restarts the app and clears all user data.
func (q Ctx) HardReset() {
	q.Javascript(`window.localStorage.clear(); window.location = "/";`)
}

//InstallJS installs the app to the user's device.
const InstallJS = `
	function install() {
		if (AddToHomeScreenEvent) {
			AddToHomeScreenEvent.prompt();
			return;
		}

		//Provide instructions.
		let instructions = document.createElement("div");
		instructions.id = "install_instructions"

		let text = document.createElement("span");
		text.style.padding = "2em";

		//IOS Safari.
		if (navigator.vendor && navigator.vendor.indexOf('Apple') > -1 &&
			navigator.userAgent &&
			navigator.userAgent.match(/iPhone|iPad|iPod/i) &&
			navigator.userAgent.indexOf('CriOS') == -1 &&
			navigator.userAgent.indexOf('FxiOS') == -1) {

			text.innerText = "Click on the share button\n swipe the bottom row to the left\n tap on 'Add to Home Screen'";

			let iosShareButton = document.createElement("div");
			iosShareButton.innerHTML = '<svg xmlns="http://www.w3.org/2000/svg" width="51.518" height="75.01" viewBox="0 0 51.518 75.01"><g id="noun_Share_1504278" transform="translate(-25 -12.2)"><path id="Path_2055" data-name="Path 2055" d="M52.307,14.055,50.452,12.2,48.6,14.055h-.206v.206L35.1,27.552l2.885,2.885L48.392,20.031V65.573h4.121V20.031L62.92,30.437,65.8,27.552,52.513,14.261v-.206Z" transform="translate(0.307)" fill="rgba(2,0,0,0.8)"/><path id="Path_2056" data-name="Path 2056" d="M25,86.427H76.518V38H62.093v4.121H72.4V82.306H29.121V42.121h10.3V38H25Z" transform="translate(0 0.783)" fill="rgba(2,0,0,0.8)"/></g></svg>'

			instructions.appendChild(iosShareButton);
			instructions.appendChild(text);
		} else if (navigator.userAgent.match(/iPhone|iPad|iPod/i)) {
			text.innerText = "Open this app in Safari to install it";
			instructions.appendChild(text);
		} else {       
			alert("Sorry, it looks like we cannot install this app to your device, you can continue using it in the browser.");
			return;
		}

		document.body.appendChild(instructions);
	}
`

//Install installs the app to the user's device.
func (q Ctx) Install() {
	q.Require(InstallJS)
	q.Javascript(`install();`)
}
