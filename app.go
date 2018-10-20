package app

import "github.com/qlova/app/worker"
import "github.com/qlova/app/manifest"
import "github.com/qlova/app/style"
import "github.com/qlova/app/script"
import "github.com/qlova/app/interfaces"

import (
	"net/http"
	"fmt"
	"path"
	"path/filepath"
	"bytes"
	"os"
	"log"
	"html"
	"strings"
)


var ServiceWorker worker.Service

func RegisterAsset(path string) {
	ServiceWorker.Assets = append(ServiceWorker.Assets, path)
}

var id = 0;

func App() *Web {
	return New()
}

type Web struct {
	style.Style
	
	id string
	tag, attr string
	children []interfaces.App
	
	fonts bytes.Buffer
	
	content []byte
	page bool
	
	onclick func(*script.Script)
	
	parent interfaces.App
	
	manifest manifest.Manifest
	handlers []func(w http.ResponseWriter, r *http.Request)
}

//Create a new qlapp, an amazing progressive web app.
func New() *Web {
	app := new(Web)
	app.Style = style.New()
	app.id = fmt.Sprint(id)
	app.tag = "div"
	
	app.manifest = manifest.New()
	
	
	id++
	return app
}

func (app *Web) ID() string {
	return app.id
}


func (app *Web) SetName(name string) {
	app.manifest.Name = name
}
func (app *Web) SetDescription(description string) {
	app.manifest.Description = description
}
func (app *Web) SetIcon(path string) {
	
}

func (app *Web) AddFont(name, file, weight string) {
	
	switch weight {
		case "black":
			weight = "900"
		case "semi-bold":
			weight = "600"
		case "regular":
			weight = "400"
		case "light":
			weight = "300"
		case "extra-light":
			weight = "200"
	}
	
	RegisterAsset(file)
	
	app.fonts.Write([]byte(`@font-face {
	font-family: '`+name+`';
	src: url('`+file+`');
	font-weight: `+weight+`;
}
`))
}

func (app *Web) GetStyle() *style.Style {
	return &app.Style
}

func (app *Web) Page() bool {
	return app.page
}

//Add a child app to the app. Remember, everything is an app!
func (app *Web) Add(child interfaces.App) {
	app.children = append(app.children, child)
	child.SetParent(app)
}

//Add a child app to the app. Remember, everything is an app!
func (app *Web) AddHandler(handler func(w http.ResponseWriter, r *http.Request)) {
	app.handlers = append(app.handlers, handler)
}


func (app *Web) GetParent() interfaces.App {
	return app.parent
}


func (app *Web) SetParent(parent interfaces.App) {
	app.parent = parent
}

func (app *Web) GetChildren() []interfaces.App {
	return app.children
}

//Add text, html or whatever!
func (app *Web) SetContent(data string) {
	app.content = []byte(data)
}

//Add text, html or whatever!
func (app *Web) SetText(data string) {
	data = html.EscapeString(data)
	data = strings.Replace(data, "\n", "<br>", -1)
	app.content = []byte(data)
}


func (app *Web) OnClick(f func(*script.Script)) {
	app.onclick = f
}

func SetPage(page interfaces.App) {
	for _, child := range page.GetParent().GetChildren() {
		if child.Page() {
			if child.ID() == page.ID() {
				child.GetStyle().SetVisible()
			} else {
				child.GetStyle().SetHidden()
			}
		}
	}
}

func (app *Web) Render() ([]byte) {
	var html bytes.Buffer
	
	html.WriteByte('<')
	html.WriteString(app.tag)
	html.WriteByte(' ')
	if app.attr != "" {
		html.WriteString(app.attr)
		html.WriteByte(' ')
	}
	html.WriteString("id='")
	html.WriteString(fmt.Sprint(app.id))
	html.WriteByte('\'')
	
	if app.Style.Css.(*style.StaticCss).Data.Bytes() != nil {
		html.WriteString(" style='")
		html.Write(app.Style.Css.(*style.StaticCss).Data.Bytes())
		html.WriteByte('\'')
	}
	
	if app.onclick != nil {
		var script = new(script.Script)
		app.onclick(script)
		
		html.WriteString(" onclick='")
		html.Write(script.Bytes())
		html.WriteByte('\'')
	}
	html.WriteByte('>')
	
	if app.content != nil {
		html.Write(app.content)
	}
	
	for _, child := range app.children {
		html.Write(child.Render())
	}
	
	html.WriteString("</")
	html.WriteString(app.tag)
	html.WriteByte('>')
	
	return html.Bytes()
}

func (app *Web) Host(hostport string) error {
	
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
    if err != nil {
            log.Fatal(err)
    }
	
	var html = app.Render()
	var worker = ServiceWorker.Render()
	var manifest = app.manifest.Render()
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request)  {
		
		fmt.Println(r.URL.Path)
		
		if r.URL.Path == "/index.js" {
			w.Header().Set("content-type", "text/javascript")
			w.Write(worker)
			return
		}
		
		if r.URL.Path == "/app.webmanifest" {
			w.Header().Set("content-type", "application/json")
			w.Write(manifest)
			return
		}
		
		if path.Ext(r.URL.Path) != "" {
			http.ServeFile(w, r, dir+"/assets"+r.URL.Path)
			return
		}
		
		if r.URL.Path != "/" {
			for _, handler := range app.handlers {
				handler(w, r)
			}
			return
		}
		
		w.Write([]byte(`<html><head>
			<meta name="viewport" content="height=device-height, 
                      width=device-width, initial-scale=1.0, 
                      minimum-scale=1.0, maximum-scale=1.0, 
                      user-scalable=no, target-densitydpi=device-dpi">

			<link rel="manifest" href="/app.webmanifest">
			
			<script>
				if ('serviceWorker' in navigator) {
					window.addEventListener('load', function() {
						navigator.serviceWorker.register('/index.js').then(function(registration) {
							console.log('ServiceWorker registration successful with scope: ', registration.scope);
						}, function(err) {
							console.log('ServiceWorker registration failed: ', err);
						});
					});
				}
			</script>
			
			
			<style>
				`))
		
		w.Write(app.fonts.Bytes())
		
		w.Write([]byte(`
			</style>
			
		<style>
			
			 .circle {
				width: 50px;
				height: 50px;
				-webkit-border-radius: 25px;
				-moz-border-radius: 25px;
				border-radius: 25px;
				background: red;
			}
			
			 html, body {
				overscroll-behavior: none; 
				cursor: pointer; 
				margin: 0; 
				height: 100%;
				-webkit-touch-callout: none;
				-webkit-user-select: none;
				-khtml-user-select: none;
				-moz-user-select: none;
				-ms-user-select: none;
				user-select: none;
				-webkit-tap-highlight-color: transparent;
			}
		</style>
		
		</head><body>`))
			w.Write(html)
		w.Write([]byte(`</body></html>`))
	})
	
	return http.ListenAndServe(hostport, nil)
}
