package app

import (
	"net/http"
	"fmt"
	"path"
	"path/filepath"
	"bytes"
	"os"
	"log"
)

var id = 0;

//Everything is an app.
type App interface {
	ID() string
	
	GetStyle() *Style
	
	Add(App)
	GetParent() App 
	SetParent(App)
	GetChildren() []App
	
	Page() bool
	
	SetContent(content string)
	
	OnClick(func(*Script))

	Render() []byte
}

type Web struct {
	Style
	
	id string
	tag, attr string
	children []App
	
	content []byte
	page bool
	
	onclick []byte
	
	parent App
}

//Create a new qlapp, an amazing progressive web app.
func New() *Web {
	app := new(Web)
	app.Style.css = new(StaticCss)
	app.id = fmt.Sprint(id)
	app.tag = "div"
	id++
	return app
}

func (app *Web) ID() string {
	return fmt.Sprint(app.id)
}

func (app *Web) GetStyle() *Style {
	return &app.Style
}

func (app *Web) Page() bool {
	return app.page
}

//Add a child app to the app. Remember, everything is an app!
func (app *Web) Add(child App) {
	app.children = append(app.children, child)
	child.SetParent(app)
}

func (app *Web) GetParent() App {
	return app.parent
}


func (app *Web) SetParent(parent App) {
	app.parent = parent
}

func (app *Web) GetChildren() []App {
	return app.children
}

//Add text, html or whatever!
func (app *Web) SetContent(data string) {
	app.content = []byte(data)
}

func (app *Web) OnClick(f func(*Script)) {
	var script = new(Script)
	f(script)
	
	app.onclick = script.Bytes()
}

func SetPage(page App) {
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
	
	if app.Style.css.(*StaticCss).data.Bytes() != nil {
		html.WriteString(" style='")
		html.Write(app.Style.css.(*StaticCss).data.Bytes())
		html.WriteByte('\'')
	}
	
	if app.onclick != nil {
		html.WriteString(" onclick='")
		html.Write(app.onclick)
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
	var worker = DefaultWorker.Render()
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request)  {
		
		if r.URL.Path == "/index.js" {
			w.Header().Set("content-type", "text/javascript")
			w.Write(worker)
			return
		}
		
		if path.Ext(r.URL.Path) != "" {
			http.ServeFile(w, r, dir+"/assets"+r.URL.Path)
			return
		}
		
		w.Write([]byte(`<html><head>
			<meta name="viewport" content="height=device-height, 
                      width=device-width, initial-scale=1.0, 
                      minimum-scale=1.0, maximum-scale=1.0, 
                      user-scalable=no, target-densitydpi=device-dpi">
			
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
			video {
				position: absolute;
				top: 100px;
				left: 0;
				width: 100vw;
				height: calc(100vh - 100px);
				z-index: -100;
			}
			
			input {
				position: absolute;
				top:0;
				left:0;
				width: 100vw;
				height: 100px;
			}
			
			 .circle {
				width: 50px;
				height: 50px;
				-webkit-border-radius: 25px;
				-moz-border-radius: 25px;
				border-radius: 25px;
				background: red;
			}
			
			 html, body {margin: 0; height: 100%}
		</style>
		</head><body>`))
			w.Write(html)
		w.Write([]byte(`</body></html>`))
	})
	
	return http.ListenAndServe(hostport, nil)
}
