package seed

import (
	"bytes"
	"fmt"
	"path"
)

import "github.com/qlova/seed/script"
import "github.com/qlova/seed/style"

type Platform int

const (
	Default Platform = iota

	Desktop
	Mobile
	Tablet
	Watch
	Tv
	Playstation
	Xbox
)

func (seed Seed) ShortCircuit(platform Platform) Seed {
	if platform == Desktop && seed.desktop.seed != nil {
		return seed.desktop
	}
	return Seed{}
}

func (seed Seed) buildStyleSheet(platform Platform, sheet *style.Sheet) {
	if short := seed.ShortCircuit(platform); short.seed != nil {
		short.buildStyleSheet(platform, sheet)
		return
	}

	//seed.postProduction()
	if data := seed.Style.Bytes(); data != nil {
		seed.styled = true
		if seed.template {
			sheet.Add("."+seed.id, seed.Style)
		} else {
			sheet.Add("#"+seed.id, seed.Style)
		}
	}
	for _, child := range seed.children {
		child.Root().buildStyleSheet(platform, sheet)
	}
}

func (seed Seed) BuildStyleSheet(platform Platform) style.Sheet {
	var stylesheet = make(style.Sheet)
	seed.buildStyleSheet(platform, &stylesheet)
	return stylesheet
}

func (seed Seed) buildStyleSheetForLandscape(platform Platform, sheet *style.Sheet) {
	if short := seed.ShortCircuit(platform); short.seed != nil {
		short.buildStyleSheetForLandscape(platform, sheet)
		return
	}

	//seed.postProduction()
	if data := seed.Landscape.Bytes(); data != nil {
		seed.styled = true
		if seed.template {
			sheet.Add("."+seed.id, seed.Landscape)
		} else {
			sheet.Add("#"+seed.id, seed.Landscape)
		}
	}
	for _, child := range seed.children {
		child.Root().buildStyleSheetForLandscape(platform, sheet)
	}
}

func (seed Seed) BuildStyleSheetForLandscape(platform Platform) style.Sheet {
	var stylesheet = make(style.Sheet)
	seed.buildStyleSheetForLandscape(platform, &stylesheet)
	return stylesheet
}

func (seed Seed) BuildStyleSheetForPortrait(platform Platform) style.Sheet {
	var stylesheet = make(style.Sheet)
	seed.buildStyleSheetForPortrait(platform, &stylesheet)
	return stylesheet
}

func (seed Seed) buildStyleSheetForPortrait(platform Platform, sheet *style.Sheet) {
	if short := seed.ShortCircuit(platform); short.seed != nil {
		short.buildStyleSheetForPortrait(platform, sheet)
		return
	}

	//seed.postProduction()
	if data := seed.Portrait.Bytes(); data != nil {
		seed.styled = true
		if seed.template {
			sheet.Add("."+seed.id, seed.Portrait)
		} else {
			sheet.Add("#"+seed.id, seed.Portrait)
		}
	}
	for _, child := range seed.children {
		child.Root().buildStyleSheetForPortrait(platform, sheet)
	}
}

//Replace this seed with its desktop version.

func (seed Seed) HTML(platform Platform) []byte {
	if short := seed.ShortCircuit(platform); short.seed != nil {
		return short.HTML(platform)
	}
	if seed.template {

		for _, child := range seed.children {
			child.Root().Render(platform)
		}

		return nil
	}

	//seed.postProduction()

	var html bytes.Buffer

	html.WriteByte('<')
	html.WriteString(seed.tag)
	html.WriteByte(' ')
	if seed.attr != "" {
		html.WriteString(seed.attr)
		html.WriteByte(' ')
	}
	html.WriteString("id='")
	html.WriteString(fmt.Sprint(seed.id))
	html.WriteByte('\'')

	if seed.class != "" {
		html.WriteString("class='")
		html.WriteString(seed.class)
		html.WriteByte('\'')
	}

	if data := seed.Style.Bytes(); !seed.styled && data != nil {
		html.WriteString(" style='")
		html.Write(data)
		html.WriteByte('\'')
	}

	if seed.onclick != nil && seed.parent == nil {
		html.WriteString(" onclick='")
		html.WriteString(script.ToJavascript(seed.onclick))
		html.WriteByte('\'')
	}

	if seed.onchange != nil {
		html.WriteString(" onchange='")
		html.WriteString(script.ToJavascript(seed.onchange))
		html.WriteByte('\'')
	}

	html.WriteByte('>')

	if seed.content != nil {
		html.Write(seed.content)
	}

	for _, child := range seed.children {
		html.Write(child.Root().Render(platform))
	}

	switch seed.tag {
	case "input", "img", "br", "hr", "meta", "area", "base", "col", "embed", "link", "param", "source", "track", "wbr":

	default:
		html.WriteString("</")
		html.WriteString(seed.tag)
		html.WriteByte('>')
	}

	return html.Bytes()
}

func (seed Seed) Render(platform Platform) []byte {
	return seed.HTML(platform)
}

func (seed Seed) getScripts(platform Platform) []string {
	if short := seed.ShortCircuit(platform); short.seed != nil {
		return short.getScripts(platform)
	}

	if seed.template {
		return nil
	}

	var scripts = seed.scripts

	for _, child := range seed.children {
		scripts = append(scripts, child.Root().getScripts(platform)...)
	}

	return scripts
}

func (seed Seed) Scripts(platform Platform) map[string]struct{} {
	var scripts = seed.getScripts(platform)
	var uniques = make(map[string]struct{})

	for _, script := range scripts {
		uniques[script] = struct{}{}
	}

	return uniques
}

func (seed Seed) buildOnReady(platform Platform, buffer *bytes.Buffer) {

	if short := seed.ShortCircuit(platform); short.seed != nil {
		short.buildOnReady(platform, buffer)
		return
	}

	if seed.template {
		return
	}

	for _, child := range seed.children {
		child.Root().buildOnReady(platform, buffer)
	}

	if seed.onready != nil {
		seed.ready = true
		buffer.WriteByte('{')
		buffer.WriteString(script.ToJavascript(seed.onready))
		buffer.WriteByte('}')
	}
}

func (seed Seed) BuildOnReady(platform Platform) []byte {
	var buffer bytes.Buffer
	buffer.WriteString(`document.addEventListener('DOMContentLoaded', function() {`)

	seed.buildOnReady(platform, &buffer)

	buffer.WriteString(`}, false);`)
	return buffer.Bytes()
}

//Return a fully fully rendered application in HTML for the seed.
func (application App) render(production bool, platform Platform) []byte {
	var seed = application.Seed

	seed.OnReady(func(q Script) {
		q.Javascript(`window.addEventListener('load', function() {
			//Load persistent state.
			if (window.localStorage) {
				let current_page = window.localStorage.getItem('*CurrentPage');
				if (current_page) {
					goto(current_page);
				}
			}
		});		
		`)
	})

	var style = seed.BuildStyleSheet(platform).Bytes()
	var styleForLandscape = seed.BuildStyleSheetForLandscape(platform).Bytes()
	var styleForPortrait = seed.BuildStyleSheetForPortrait(platform).Bytes()
	var html = seed.HTML(platform)
	var scripts = seed.Scripts(platform)
	var onready = seed.BuildOnReady(platform)

	var buffer bytes.Buffer
	buffer.Write([]byte(`<!DOCTYPE html><html><head>
		<meta name="viewport" content="viewport-fit=cover, height=device-height, width=device-width, initial-scale=1.0, minimum-scale=1.0, maximum-scale=1.0, user-scalable=no">
	`))

	if application.tracking != "" {
		buffer.WriteString(`
		<!-- Global site tag (gtag.js) - Google Analytics -->
		<script async src="https://www.googletagmanager.com/gtag/js?id=` + application.tracking + `"></script>
		<script>
		  window.dataLayer = window.dataLayer || [];
		  function gtag(){dataLayer.push(arguments);}
		  gtag('js', new Date());
		
		  gtag('config', 'UA-134084549-1');
		</script>
		`)
	}

	if platform != Desktop {
		buffer.WriteString(`
			<meta name="apple-mobile-web-app-capable" content="yes">
			<meta name="apple-mobile-web-app-status-bar-style" content="black">
			<meta name="mobile-web-app-capable" content="yes">
			<meta name="apple-mobile-web-app-title" content="` + application.Manifest.Name + `" />
		`)
	}

	buffer.Write([]byte(`
		<meta name="theme-color" content="` + application.Manifest.ThemeColor + `">

		<title>` + application.Manifest.Name + `</title>

		<link rel="manifest" href="/app.webmanifest">`))

	for i, icon := range application.Manifest.Icons {

		//The first icon can be the Favicon.
		if i == 0 {
			buffer.WriteString(`<link rel="shortcut icon" type="image/png" href="` + icon.Source + `"/>`)
		}

		buffer.Write([]byte(`<link rel="apple-touch-icon" sizes="` + icon.Sizes + `" href="` + icon.Source + `">`))
	}

	for script := range scripts {
		if path.Ext(script) == ".js" {
			buffer.Write([]byte(`<script src="` + script + `"></script>`))
		} else if path.Ext(script) == ".css" {
			buffer.Write([]byte(`<link rel="stylesheet" href="` + script + `" />`))
		}
	}

	buffer.Write([]byte(`<script>
			if ('serviceWorker' in navigator) {
				navigator.serviceWorker.register('/index.js').then(function(registration) {

					registration.onupdatefound = function() {
						window.localStorage.setItem("update", "true");

						` + script.ToJavascript(application.onupdatefound) + `
					}
				
					console.log('ServiceWorker registration successful with scope: ', registration.scope);
				}, function(err) {
					console.log('ServiceWorker registration failed: ', err);
				});
			}

			//Some SERIOUS HACKS TO WORKAROUND APPLE'S BUGS.
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
		          				console.log("GET YOUR ACT TOGETHER WEBKIT: the viewport-fit: cover bug is seriously annoying.")
		          				document.querySelector('meta[name=viewport]').
		          					setAttribute('content', 'width=device-width, initial-scale=1.0, minimum-scale=1.0, maximum-scale=1.0, user-scalable=no');
		          				
		          		}
		            }
		        }
		    }
			
		</script>
		
		
		<style>
	`))

	buffer.Write(application.Fonts())
	buffer.Write(application.Animations())
	buffer.Write(style)

	buffer.WriteString(`@media screen and (orientation: landscape) {`)
	buffer.Write(styleForLandscape)
	buffer.WriteString(`}`)

	buffer.WriteString(`@media screen and (orientation: portrait) {`)
	buffer.Write(styleForPortrait)
	buffer.WriteString(`}`)

	buffer.Write([]byte(`
		</style>
			
		<style>			
			`))

	/*if platform == Desktop {
		buffer.WriteString(`
			::-webkit-scrollbar {
				display: none;
			}
		`)
	}*/

	buffer.Write([]byte(`
			* {
				-webkit-tap-highlight-color: rgba(255, 255, 255, 0) !important; 
				-webkit-focus-ring-color: rgba(255, 255, 255, 0) !important; 
				outline: none !important;
				flex-shrink: 0;
			}

			a {
				text-decoration: none;
			}
			
			p {
				margin-block-start: 0;
				margin-block-end: 0;
			}
			
			img {
				object-fit: contain;
			}
			
			html {
				height: 100vh;
				box-sizing: border-box;
			}
			*, *:before, *:after {
				box-sizing: inherit;
			}
			pre, hr {
				margin: 0;
			}

			div {
				position: relative;
			}

			<!-- TODO These animations should be added dynamically by seed/script/animation.go -->
			@keyframes slideInFromLeft {
				from { transform: translateX(-100%); }
				to { transform: translateX(0); }
			}

			@keyframes slideInFromRight {
				from { transform: translateX(100%); }
				to { transform: translateX(0); }
			}

			@keyframes slideInFromTop {
				from { transform: translateY(-100%); }
				to { transform: translateY(0); }
			}

			@keyframes slideInFromBottom {
				from { transform: translateY(100%); }
				to { transform: translateY(0); }
			}

			@keyframes slideOutToLeft {
				from { transform: translateX(0); }
				to { transform: translateX(-100%); }
			}

			@keyframes slideOutToRight {
				from { transform: translateX(); }
				to { transform: translateX(100%); }
			}

			@keyframes slideOutToTop {
				from { transform: translateY(0); }
				to { transform: translateY(-100%); }
			}

			@keyframes slideOutToBottom {
				from { transform: translateY(0); }
				to { transform: translateY(100%); }
			}
			
			@keyframes fadeOut {
				from { opacity: 1; }
				to { opacity: 0; }
			}
			
			@keyframes fadeIn {
				from { opacity: 0; }
				to { opacity: 1; }
			}
			
			body {
				top: 0;
				left: 0
				right: 0;
				bottom: 0;
				position: fixed;
				overscroll-behavior: none; 
				-webkit-overscroll-behavior: none; 
				-webkit-overflow-scrolling: none; 
				cursor: pointer; 
				margin: 0; 
				padding: 0;
				height: 100%;
				width: 100vw;
				`))

	//We dont want people to select things on mobile.
	if platform != Desktop {
		buffer.WriteString(`
						-webkit-touch-callout: none;
						-webkit-user-select: none;
						-khtml-user-select: none;
						-moz-user-select: none;
						-ms-user-select: none;
						user-select: none;
						-webkit-tap-highlight-color: transparent;
					`)
	}

	buffer.Write([]byte(`
				/* Some nice defaults for centering content. */
				display: inline-flex;
				align-items: center;
				justify-content: center;
				flex-direction: row;
				overflow: hidden;
			}
		</style>
		
		<script>`))

	if platform != Desktop {
		buffer.WriteString(`
				//Some SERIOUS HACKS TO WORKAROUND APPLE'S BUGS.
				//NO BOUNCE BUGS GODDAMMIT https://github.com/lazd/iNoBounce
							(function(global){var startY=0;var enabled=false;var supportsPassiveOption=false;try{var opts=Object.defineProperty({},"passive",{get:function(){supportsPassiveOption=true}});window.addEventListener("test",null,opts)}catch(e){}var handleTouchmove=function(evt){var el=evt.target;var zoom=window.innerWidth/window.document.documentElement.clientWidth;if(evt.touches.length>1||zoom!==1){return}while(el!==document.body&&el!==document){var style=window.getComputedStyle(el);if(!style){break}if(el.nodeName==="INPUT"&&el.getAttribute("type")==="range"){return}var scrolling=style.getPropertyValue("-webkit-overflow-scrolling");var overflowY=style.getPropertyValue("overflow-y");var height=parseInt(style.getPropertyValue("height"),10);var isScrollable=scrolling==="touch"&&(overflowY==="auto"||overflowY==="scroll");var canScroll=el.scrollHeight>el.offsetHeight;if(isScrollable&&canScroll){var curY=evt.touches?evt.touches[0].screenY:evt.screenY;var isAtTop=startY<=curY&&el.scrollTop===0;var isAtBottom=startY>=curY&&el.scrollHeight-el.scrollTop===height;if(isAtTop||isAtBottom){evt.preventDefault()}return}el=el.parentNode}evt.preventDefault()};var handleTouchstart=function(evt){startY=evt.touches?evt.touches[0].screenY:evt.screenY};var enable=function(){window.addEventListener("touchstart",handleTouchstart,supportsPassiveOption?{passive:false}:false);window.addEventListener("touchmove",handleTouchmove,supportsPassiveOption?{passive:false}:false);enabled=true};var disable=function(){window.removeEventListener("touchstart",handleTouchstart,false);window.removeEventListener("touchmove",handleTouchmove,false);enabled=false};var isEnabled=function(){return enabled};var testDiv=document.createElement("div");document.documentElement.appendChild(testDiv);testDiv.style.WebkitOverflowScrolling="touch";var scrollSupport="getComputedStyle"in window&&window.getComputedStyle(testDiv)["-webkit-overflow-scrolling"]==="touch";document.documentElement.removeChild(testDiv);if(scrollSupport){enable()}var iNoBounce={enable:enable,disable:disable,isEnabled:isEnabled};if(typeof module!=="undefined"&&module.exports){module.exports=iNoBounce}if(typeof global.define==="function"){(function(define){define("iNoBounce",[],function(){return iNoBounce})})(global.define)}else{global.iNoBounce=iNoBounce}})(this);
			`)
	}

	//Need to actually detect if we are running inside a dev environment or not!
	//probably should check the request hostname in launcher to decide if we are in production or not.
	if production && (application.rest != "") {
		buffer.WriteString(`var host = "https://` + application.rest + `";`)
	} else {
		buffer.WriteString(`var host = "";`)
	}

	buffer.Write([]byte(`
			window.onorientationchange = function() {
				window.dispatchEvent(new Event('orientationchange'));
			}
	
			var geoLocation = null;
			var requestGeoLocation = function (options) {
				return new Promise(function (resolve, reject) {
					navigator.geolocation.getCurrentPosition(function(position) {
						geoLocation = position;
						resolve(position);
					}, reject, options);
				});
			}
	
			var animating = false;
			
			var get = function(id) {
				return document.getElementById(id)
			};
			
			var goto_queue = [];
			var animation_complete = function() {
				animating = false;
				
				//Process goto queue.
				let next = goto_queue.shift();
				if (next != null) {
					goto(next);
				}
			}
			
			var goto_history = [];

			var last_page = null;
			var current_page = null;
			var next_page = null;
			var goto = function(next_page_id) {
			
				if (get(next_page_id) == null || get(next_page_id).className != "page" || next_page_id == "` + application.loadingPage.ID() + `") {
					next_page_id = "` + application.startingPage.ID() + `"
				}
			
				if (animating) {
					goto_queue.push(next_page_id)
					return;
				}
				if (current_page == next_page_id) return;
				if (next_page == next_page_id) return;
				next_page = next_page_id;

				for (let element of get(next_page_id).parentElement.childNodes) {
					if (element.classList.contains("page")) {
						if (getComputedStyle(element).display != "none") {
							set(element, 'display', 'none');						
							if (element.exitpage) element.exitpage();
							last_page = element.id;
						}
					}
				}
				
				if (last_page != null) {
					goto_history.push(last_page);
				}
				
				let next_element = get(next_page_id);
				if (next_element) {
					set(next_element, 'display', 'inline-flex');
					if (next_element.enterpage) next_element.enterpage();
					current_page = next_page_id;

					//Persistence.
					window.localStorage.setItem('*CurrentPage', next_page_id);
					
				}
				next_page = null;
				
			
			};
			
			function ` + OnPress + `(id, func, propagate) {
				let element = get(id);
				
				let handler = function(event) {
					func(event);
				};
				
				let moved = false;
				let point = [0, 0];
				
				element.ontouchstart = function(e) {
					var changedTouch = event.changedTouches[0];
						point[0]  = changedTouch.clientX;
						point[1]  = changedTouch.clientY;
				};
				
				element.ontouchmove = function(event) {
					var changedTouch = event.changedTouches[0];
					var elem = document.elementFromPoint(changedTouch.clientX, changedTouch.clientY);
								
					if (elem != event.target) moved = true;
								
					let a = changedTouch.clientX - point[0];
					let b = changedTouch.clientY - point[1];
					if ((a*a + b*b) > 50*50) moved = true;
				};
				
				element.ontouchend = function(ev) {
					if (ev.stopPropagation && !propagate) ev.stopPropagation(); 
					ev.preventDefault(); 
					if (moved) {
						moved = false; 
						return; 
					}
					ev = ev.changedTouches[0];
					handler(ev);
				};

				element.onclick = handler;
			}

			var ActivePhotoSwipe = null;
			
			var back = function() {

				if (ActivePhotoSwipe) {
					ActivePhotoSwipe.close();
					return;
				}
				if (goto_history.length == 0) return;
				
				
				let last_page = goto_history.pop();
				if (last_page == null) return;
						
				let old_length = goto_history.length;
						
				goto(last_page);
				if (goto_history.length  > old_length)
				goto_history.pop();
			};

			function setCookie(cname, cvalue, exdays) {
			  var d = new Date();
			  d.setTime(d.getTime() + (exdays*24*60*60*1000));
			  var expires = "expires="+ d.toUTCString();
			  document.cookie = cname + "=" + cvalue + ";" + expires + ";secure;path=/";
			}

			function getCookie(cname) {
			  var name = cname + "=";
			  var decodedCookie = decodeURIComponent(document.cookie);
			  var ca = decodedCookie.split(';');
			  for(var i = 0; i <ca.length; i++) {
			    var c = ca[i];
			    while (c.charAt(0) == ' ') {
			      c = c.substring(1);
			    }
			    if (c.indexOf(name) == 0) {
			      return c.substring(name.length, c.length);
			    }
			  }
			  return "";
			}

			function request (method, formdata, url, manual) {
				if (url.charAt(0) == "/") url = host+url;
			
				if (manual) {
					 var xhr = new XMLHttpRequest();
					 xhr.open(method, url);
					return xhr;
				}
			
			  return new Promise(function (resolve, reject) {
			    var xhr = new XMLHttpRequest();
			    xhr.open(method, url, true);
			    xhr.onload = function () {
			      if (this.status >= 200 && this.status < 300) {
			        resolve(xhr.response);
			      } else {
			        reject({
			          status: this.status,
			          statusText: xhr.statusText,
					 response: xhr.response
			        });
			      }
			    };
			    xhr.onerror = function () {
			      reject({
			        status: this.status,
			        statusText: xhr.statusText,
			        response: xhr.response
			      });
			    };
			    xhr.send(formdata);
			  });
			}

			!function(t,e){"use strict";function n(){this.dispatchEvent(new CustomEvent("long-press",{bubbles:!0,cancelable:!0})),clearTimeout(o),console&&console.log&&console.log("long-press fired on "+this.outerHTML)}var o=null,u="ontouchstart"in t||navigator.MaxTouchPoints>0||navigator.msMaxTouchPoints>0,s=u?"touchstart":"mousedown",i=u?"touchcancel":"mouseout",a=u?"touchend":"mouseup",c=u?"touchmove":"mousemove";"initCustomEvent"in e.createEvent("CustomEvent")&&(t.CustomEvent=function(t,n){n=n||{bubbles:!1,cancelable:!1,detail:void 0};var o=e.createEvent("CustomEvent");return o.initCustomEvent(t,n.bubbles,n.cancelable,n.detail),o},t.CustomEvent.prototype=t.Event.prototype),e.addEventListener(s,function(t){var e=t.target,u=parseInt(e.getAttribute("data-long-press-delay")||"500",10);o=setTimeout(n.bind(e),u)}),e.addEventListener(a,function(t){clearTimeout(o)}),e.addEventListener(i,function(t){clearTimeout(o)}),e.addEventListener(c,function(t){clearTimeout(o)})}(this,document);
			
		`))

	if !production {
		buffer.Write([]byte(`
		
			var set = function(element, property, value) {
				if (!(element.id in InternalStyleState)) {
					InternalStyleState[element.id] = {};
				}
				element.style[property] = value;
				InternalStyleState[element.id][property] = element.style[property].trim();
			};
			
			var InternalStyleState = {};

			
			function setCookie(cname, cvalue, exdays) {
			  var d = new Date();
			  d.setTime(d.getTime() + (exdays*24*60*60*1000));
			  var expires = "expires="+ d.toUTCString();
			  document.cookie = cname + "=" + cvalue + ";" + expires + ";path=/";
			}
			
			if (window.location.hostname.includes("localhost")) {
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

				function parseCss(attribute) {
					let css = {};
					if (!attribute) return css;
					
					attribute = attribute.replace(/(\/\*([\s\S]*?)\*\/)|(\/\/(.*)$)/gm, '');
											
					//Gonna have to parse the css.
					let styles = attribute.split(';');
					for (let style of styles) {
						if (style == "") continue;
						let splits = style.split(':');
						let property = splits[0];
						let value = splits[1];
						if (value == undefined) continue;
						
						css[property] = value;
					}

					return css;
				}
				
				var edits = {};
				window.addEventListener('load', function() {
					var observer = new MutationObserver(function(mutations) {
						mutations.forEach(function(mutation) {
							if (mutation.target.id == "") return;
							
							let style = parseCss(mutation.target.getAttribute("style"));
								
							for (let property in style) {
								let value = style[property];
							
								if (mutation.target.id in InternalStyleState && InternalStyleState[mutation.target.id][property] == value.trim()) {
									continue;
								}
								
								if (!(mutation.target.id in edits)) {
									edits[mutation.target.id] = {};
								}
								edits[mutation.target.id][property] = true;
							}
							
							//InternalStyleState[mutation.target][]
						});    
					});
	
					const observerConfig = {
					
						attributes: true, // attribute changes will be observed | on add/remove/change attributes
						attributeOldValue: true, // will show oldValue of attribute | on add/remove/change attributes | default: null
						
						characterData: true, // data changes will be observed | on add/remove/change characterData
						characterDataOldValue: true, // will show OldValue of characterData | on add/remove/change characterData | default: null
						
						childList: true, // target childs will be observed | on add/remove
						subtree: true, // target childs will be observed | on attributes/characterData changes if they observed on target
						
						attributeFilter: ['style'] // filter for attributes | array of attributes that should be observed, in this case only style
					
					};
	
					observer.observe(document, observerConfig);
				});
				window.addEventListener("click", function(event) {
					var an = window.getSelection().anchorNode;
				 	// this is the innermost *element*
				 	var element = an;
				 	if (element == null) return;
				 	while (!( element instanceof Element )) {
				    	element = element.parentElement;
				    	if (element == null) return;
				    }

					if (!(element.id in edits)) {
						edits[element.id] = {};
					}
					edits[element.id].text = true;
				});
				
				window.addEventListener("keypress", function(event) {
					//Edit mode.
					if (event.key == "e" && event.ctrlKey) {

						if (document.designMode == "on") {
							document.designMode = "off";
						} else {
							document.designMode = "on";
						}

						event.preventDefault();
						return true;
					}
					//Save Edits.
					if (event.key == "s" && event.ctrlKey) {
	
						for (let edit in edits) {
							
						
							let style = parseCss(get(edit).getAttribute("style"));
							let change = false;

							let message = "#"+edit+" {";

							if (edits[edit].text) {
								message += "text: ` + "`" + `"+get(edit).innerHTML+"` + "`" + `;";
								change = true;
							}
							
							for (let property in style) {
								let value = style[property];
								
								if (edit in InternalStyleState && InternalStyleState[edit][property] == value.trim()) {
									continue;
								}

								message += property.trim()+":"+value.trim()+";";
								change = true;
							}
							message += "}";
							if (change) {
								Socket.send(message)								
							}
						}

						let body = document.querySelector("body");
						body.contentEditable = "false";
						event.preventDefault();
						return true;
					}
				})
			} else {
				history.pushState(null, null, document.URL);
				window.addEventListener('popstate', function () {
					back();
					history.pushState(null, null, document.URL);
				});
			}
			`))
	}

	if production {
		buffer.Write([]byte(`
			var set = function(element, property, value) {
				element.style[property] = value;
			};

			history.pushState(null, null, document.URL);
							window.addEventListener('popstate', function () {
								back();
								history.pushState(null, null, document.URL);
							});`))
	}

	if application.DynamicHandler() != nil {
		buffer.WriteString(`
			var dynamic = new XMLHttpRequest();
	
			dynamic.onreadystatechange = function() {
				if (this.readyState == 4 && this.status == 200) {
					var updates = JSON.parse(this.responseText);
					for (let id in updates) {
						document.getElementById(id).textContent = updates[id];
					}
				}
			};
	
			dynamic.open("GET", "/dynamic", true);
			dynamic.send();`)
	}

	for name, function := range functions {
		buffer.WriteString("function ")
		buffer.WriteString(name)
		buffer.WriteString("() {")
		buffer.WriteString(script.ToJavascript(function))
		buffer.WriteString("}")
	}

	buffer.Write([]byte(`	
				</script>
				
				</head><body>
			`))
	buffer.Write(html)
	buffer.WriteString(tail)

	buffer.Write([]byte(`<script>`))
	buffer.Write(onready)
	buffer.Write([]byte(`</script>`))

	buffer.Write([]byte(`

	</body></html>`))

	return buffer.Bytes()
}

var tail string

func Tail(t string) {
	tail += t
}
