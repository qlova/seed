package seed

import (
	"bytes"
	"path"
)

//Returns rendered html of the entire app.
func (app App) HTML() []byte {
	var Style = app.BuildStyleSheet(0).Bytes()
	var Portrait = app.BuildStyleSheetForPortrait(0).Bytes()
	var Landscape = app.BuildStyleSheetForLandscape(0).Bytes()

	var scripts = app.Scripts(app.platform)
	var HTML = app.Seed.Render(app.platform)
	var StateHandlers = app.StateHandlers()
	var OnReady = app.OnReadyHandler()
	var DynamicHandlers = app.DynamicHandlers()

	var buffer bytes.Buffer
	buffer.WriteString(`<!DOCTYPE html>`)
	buffer.WriteString(`<html lang="en">`)
	buffer.WriteString(`<head>`)

	//Ensure the screen size matches the device size, disable zoom.
	buffer.WriteString(`<meta name="viewport" content="viewport-fit=cover, height=device-height, width=device-width, initial-scale=1.0, minimum-scale=1.0, maximum-scale=1.0, user-scalable=no">`)

	//This is a webapp.
	buffer.WriteString(`
		<meta name="mobile-web-app-capable" content="yes">
		<meta name="apple-mobile-web-app-capable" content="yes">
		<meta name="apple-mobile-web-app-status-bar-style" content="black">
		<meta name="apple-mobile-web-app-title" content="` + app.Name + `" />
	`)

	//Important meta tags.
	buffer.WriteString(`
		<meta name="theme-color" content="` + app.ThemeColor + `">

		<title>` + app.Name + `</title>

		<link rel="manifest" href="/app.webmanifest">
	`)

	for i, icon := range app.Icons {
		//The first icon can be the Favicon. TODO better heuristic? allow other file types.
		if i == 0 {
			buffer.WriteString(`<link rel="shortcut icon" type="image/png" href="` + icon.Source + `"/>`)
		}

		buffer.Write([]byte(`<link rel="apple-touch-icon" sizes="` + icon.Sizes + `" href="` + icon.Source + `">`))
	}

	for script := range scripts {
		if path.Ext(script) == ".css" {
			buffer.Write([]byte(`<link rel="stylesheet" href="` + script + `" />`))
		}
	}

	buffer.WriteString(`<script>`)
	{
		if app.production && (app.rest != "") {
			buffer.WriteString(`var host = "https://` + app.rest + `";`)
		} else {
			buffer.WriteString(`var host = "";`)
		}
		buffer.WriteString(`var starting_page = "` + app.startingPage.ID() + `";`)
		buffer.WriteString(`var loading_page = "` + app.loadingPage.ID() + `";`)

		//TODO clean this up.
		buffer.WriteString(`
			window.onorientationchange = function() {
				window.dispatchEvent(new Event('orientationchange'));
			}

			var ActivePhotoSwipe = null;
		`)

		if app.production {
			//Disable back-button.
			buffer.WriteString(`var production = true;
			function setCookie(cname, cvalue, exdays) {
					var d = new Date();
					d.setTime(d.getTime() + (exdays*24*60*60*1000));
					var expires = "expires="+ d.toUTCString();
					document.cookie = cname + "=" + cvalue + ";" + expires + ";secure;path=/";
				}
			history.pushState(null, null, document.URL);
			window.addEventListener('popstate', function () {
				back();
				history.pushState(null, null, document.URL);
			});`)
		}

		//Some developer-friendly features.
		if !app.production {
			buffer.WriteString(`var production = false;
			
				let url = new URL('/socket', window.location.href);
				url.protocol = url.protocol.replace('http', 'ws');
				let Socket = new WebSocket(url.href);
				Socket.onclose = function() {
					close();
				}
				Socket.onerror = function() {
					close();
				}
				Socket.onmessage = function(event) {
					eval(event.data);
				}
				//Disable refresh on chrome because otherwise the app will close.
				document.onkeydown = function() {    
					switch (event.keyCode) { 
						case 116 : //F5 button
							event.returnValue = false;
							event.keyCode = 0;
							return false; 
						case 82 : //R button
							if (event.ctrlKey) { 
								event.returnValue = false; 
								event.keyCode = 0;  
								return false; 
							} 
					}
				}
			`)
		}

		//ServiceWorker OnUpdateFound
		buffer.WriteString(`
			if ('serviceWorker' in navigator) {
				navigator.serviceWorker.register('/index.js').then(function(registration) {

					registration.onupdatefound = function() {
						if (!window.localStorage.getItem("update")) {
							window.localStorage.setItem("update", "false");
							return;
						}
						window.localStorage.setItem("update", "true");
						
						//Clear all unnamed variables because they could have changed!
						//Unamed variables have a 'g_' prefix.
						for (let i in localStorage) {
							let item = window.localStorage[i];
							if (item.substring && item.substring(0, 3) == "g_") {
								window.localStorage.removeItem(i);
							}
						}
		`)
		buffer.Write(app.ToJavascript(app.onupdatefound))

		buffer.WriteString(`
					}
				}, function(err) {
					
				});
			}
		`)

		//Mitigation for IOS viewport-fit: cover bug (https://openradar.appspot.com/radar?id=5018321736957952)
		buffer.WriteString(`
			var canvas = document.createElement("canvas");
			if (canvas) {
				var context = canvas.getContext("webgl") || canvas.getContext("experimental-webgl");
				if (context) {
					var info = context.getExtension("WEBGL_debug_renderer_info");
					if (info) {
						var renderer = context.getParameter(info.UNMASKED_RENDERER_WEBGL);

						switch (renderer) {
							case "PowerVR SGX 543":
							case "Apple A8 GPU":
							case "Apple A9 GPU":
							case "Apple A10 GPU":
								document.querySelector('meta[name=viewport]').
									setAttribute('content', 'width=device-width, initial-scale=1.0, minimum-scale=1.0, maximum-scale=1.0, user-scalable=no');
								
						}
					}
				}
			}
		`)

		//NO BOUNCE BUGS https://github.com/lazd/iNoBounce
		buffer.WriteString(`
			(function(global){var startY=0;var enabled=false;var supportsPassiveOption=false;try{var opts=Object.defineProperty({},"passive",{get:function(){supportsPassiveOption=true}});window.addEventListener("test",null,opts)}catch(e){}var handleTouchmove=function(evt){var el=evt.target;var zoom=window.innerWidth/window.document.documentElement.clientWidth;if(evt.touches.length>1||zoom!==1){return}while(el!==document.body&&el!==document){var style=window.getComputedStyle(el);if(!style){break}if(el.nodeName==="INPUT"&&el.getAttribute("type")==="range"){return}var scrolling=style.getPropertyValue("-webkit-overflow-scrolling");var overflowY=style.getPropertyValue("overflow-y");var height=parseInt(style.getPropertyValue("height"),10);var isScrollable=scrolling==="touch"&&(overflowY==="auto"||overflowY==="scroll");var canScroll=el.scrollHeight>el.offsetHeight;if(isScrollable&&canScroll){var curY=evt.touches?evt.touches[0].screenY:evt.screenY;var isAtTop=startY<=curY&&el.scrollTop===0;var isAtBottom=startY>=curY&&el.scrollHeight-el.scrollTop===height;if(isAtTop||isAtBottom){evt.preventDefault()}return}el=el.parentNode}evt.preventDefault()};var handleTouchstart=function(evt){startY=evt.touches?evt.touches[0].screenY:evt.screenY};var enable=function(){window.addEventListener("touchstart",handleTouchstart,supportsPassiveOption?{passive:false}:false);window.addEventListener("touchmove",handleTouchmove,supportsPassiveOption?{passive:false}:false);enabled=true};var disable=function(){window.removeEventListener("touchstart",handleTouchstart,false);window.removeEventListener("touchmove",handleTouchmove,false);enabled=false};var isEnabled=function(){return enabled};var testDiv=document.createElement("div");document.documentElement.appendChild(testDiv);testDiv.style.WebkitOverflowScrolling="touch";var scrollSupport="getComputedStyle"in window&&window.getComputedStyle(testDiv)["-webkit-overflow-scrolling"]==="touch";document.documentElement.removeChild(testDiv);if(scrollSupport){enable()}var iNoBounce={enable:enable,disable:disable,isEnabled:isEnabled};if(typeof module!=="undefined"&&module.exports){module.exports=iNoBounce}if(typeof global.define==="function"){(function(define){define("iNoBounce",[],function(){return iNoBounce})})(global.define)}else{global.iNoBounce=iNoBounce}})(this);
		`)

		//User-defined js functions. TODO functions should not be global.
		for name, function := range functions {
			buffer.WriteString("function ")
			buffer.WriteString(name)
			buffer.WriteString("() {")
			buffer.Write(app.ToJavascript(function))
			buffer.WriteString("}")
		}

		//Dependencies
		for script := range app.Context.Dependencies {
			buffer.WriteString(script)
		}
	}
	buffer.WriteString(`</script>`)

	buffer.WriteString(`<style>`)
	{
		//Default css from css.go
		buffer.WriteString(CSS)

		//Dependencies
		for animation, id := range app.Context.Animations {
			buffer.WriteString(`@keyframes ` + id + " {")
			buffer.Write(animation.Bytes())
			buffer.WriteString(`}`)
		}

		buffer.Write(app.Fonts())
		buffer.Write(app.Animations())
		buffer.Write(Style)

		buffer.WriteString(`@media screen and (orientation: landscape) {`)
		buffer.Write(Landscape)
		buffer.WriteString(`}`)

		buffer.WriteString(`@media screen and (orientation: portrait) {`)
		buffer.Write(Portrait)
		buffer.WriteString(`}`)

		buffer.Write(app.MediaQueries())
	}
	buffer.WriteString(`</style>`)

	//User modified head can go here.
	buffer.Write(app.Head.Bytes())
	buffer.WriteString(`</head>`)

	buffer.WriteString(`<body>`)
	buffer.Write(app.Neck.Bytes())
	buffer.Write(HTML)
	buffer.Write(app.Tail.Bytes())

	buffer.WriteString(`<script>`)
	buffer.Write(StateHandlers)
	buffer.Write(DynamicHandlers)
	buffer.Write(OnReady)

	buffer.WriteString(`</script>`)

	for script := range scripts {
		if path.Ext(script) == ".js" {
			buffer.Write([]byte(`<script src="` + script + `"></script>`))
		}
	}

	buffer.WriteString(`</body>`)
	buffer.WriteString(`</html>`)

	return buffer.Bytes()
}
