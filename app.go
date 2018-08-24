package app

import (
	"net/http"
	"fmt"
	"path"
	"bytes"
)

var id = 0;

//Everything is an app.
type App struct {
	Style
	
	id int
	tag, attr string
	children []*App
	
	content []byte
	page bool
	
	onclick []byte
}

//Create a new qlapp, an amazing progressive web app.
func New() *App {
	app := new(App)
	app.id = id
	app.tag = "div"
	id++
	return app
}

func (app *App) ID() string {
	return fmt.Sprint(app.id)
}

//Add a child app to the app. Remember, everything is an app!
func (app *App) Add(child *App) {
	app.children = append(app.children, child)
}

//Add text, html or whatever!
func (app *App) SetContent(data string) {
	app.content = []byte(data)
}

func (app *App) OnClick(f func(*Script)) {
	var script = new(Script)
	f(script)
	
	app.onclick = script.script.Bytes()
}

func (app *App) SetPage(page *App) {
	for _, child := range app.children {
		if child.page {
			if child == page {
				child.SetVisible()
			} else {
				child.SetHidden()
			}
		}
	}
}

func (app *App) Render() ([]byte) {
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
	
	if app.Style.css.Bytes() != nil {
		html.WriteString(" style='")
		html.Write(app.Style.css.Bytes())
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

func (app *App) Host(hostport string) error {
	
	var html = app.Render()
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request)  {
		
		if path.Ext(r.URL.Path) != "" {
			http.ServeFile(w, r, "assets/"+r.URL.Path)
			return
		}
		
		w.Write([]byte(`<html><head>
			<meta name="viewport" content="height=device-height, 
                      width=device-width, initial-scale=1.0, 
                      minimum-scale=1.0, maximum-scale=1.0, 
                      user-scalable=no, target-densitydpi=device-dpi">
			
			
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
			
			<script>
				window.addEventListener("load",function() {
					setTimeout(function(){
						// This hides the address bar:
						window.scrollTo(0, 1);
					}, 0);
				});
			</script>
		</head><body>`))
			w.Write(html)
		w.Write([]byte(`</body></html>`))
	})
	
	return http.ListenAndServe(hostport, nil)
}
