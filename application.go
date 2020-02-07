package seed

import (
	"bytes"
	"log"
	"net/http"
	"os"

	"github.com/qlova/seed/inbed"
	"github.com/qlova/seed/manifest"
	"github.com/qlova/seed/script"
	"github.com/qlova/seed/service"
	"github.com/qlova/seed/style/css"
	"github.com/qlova/seed/user"
)

//App is an entire app, with multiple pages and/or seeds.
type App struct {
	Seed
	manifest.Manifest
	service.Worker
	*harvester

	description string

	host, rest, pkg, tracking string

	hashes []string

	onupdatefound func(script.Ctx)

	loadingPage  Page
	startingPage Page

	Head, Neck, Body, Tail bytes.Buffer

	production bool
	platform   Platform

	built bool

	Handlers map[string]http.Handler
}

//NewApp creates a new application. The first string argument is the name, the second is a description of the app.
func NewApp(args ...string) *App {
	var app = App{
		Seed:      New(),
		Manifest:  manifest.New(),
		Worker:    service.NewWorker(),
		harvester: newHarvester(),
		Handlers:  make(map[string]http.Handler),
	}

	app.Seed.app = &app

	app.SetSize(100, 100)
	app.CSS().SetDisplay(css.Flex)
	app.TextAlign().Center()
	app.CSS().SetFlexDirection(css.Column)

	app.Icons = append(app.Icons, manifest.Icon{
		Source: "/Qlovaseed.png",
		Sizes:  "512x512",
	})

	if len(args) > 0 {
		app.SetName(args[0])
		app.description = args[0]
	}

	if len(args) > 1 {
		app.description = args[1]
	}

	return &app
}

//ToJavascript converts a Ctx function into javascript bytes.
func (app *App) ToJavascript(f func(script.Ctx)) []byte {
	return script.ToJavascript(f, app.Context)
}

//NewPage creates and returns a new page for the app.
//If this is the first page created this way, it will also create a loading "splash" page.
func (app *App) NewPage() Page {

	if app.loadingPage.Null() {
		app.loadingPage = AddPageTo(app)
		app.loadingPage.splash = true
	}

	return AddPageTo(app)
}

//SetPage sets the default starting page for this App.
//This page will be presented after the app has finished loading.
func (app *App) SetPage(page Page) {

	app.Context.AddPage(page.id, page)

	app.LoadingPage()

	app.startingPage = page

	app.OnReady(func(q script.Ctx) {
		q.Require(script.Goto)
	})
}

//SetDescription sets the description of your app.
func (app *App) SetDescription(description string) {
	app.description = description
}

//LoadingPage returns the loadingpage (like a splashscreen) for this app that displays while the app is loading.
func (app *App) LoadingPage() Page {

	if app.loadingPage.Null() {
		app.loadingPage = AddPageTo(app)
		app.loadingPage.splash = true
	}

	return app.loadingPage
}

//SetHost sets the hostname of this app, this is where the app is expected to be hosted from.
//ie. "app.example.com"
func (app *App) SetHost(name string) {
	app.host = name
}

//SetRest sets the REST hostname of this app, this is where the app will serve and request API calls.
//By default this will be relative to the current location of the app, however you may want to change this.
//ie. "api.example.com"
func (app *App) SetRest(name string) {
	app.rest = name
}

//SetPackage sets the package name of this application on android.
//By default it will be set to the reverse of the hostname.
//ie. "com.example.app"
func (app *App) SetPackage(name string) {
	app.pkg = name
}

//AddHash allows you to add a hash of the certificate that you will sign your android app with.
//This is useful for Deep linking and assetlinks.json will be automatically served.
func (app *App) AddHash(name string) {
	app.hashes = append(app.hashes, name)
}

//SetTrackingCode adds a google analytics tracking code to the app.
func (app *App) SetTrackingCode(code string) {
	app.Head.WriteString(`
		<!-- Global site tag (gtag.js) - Google Analytics -->
		<script async src="https://www.googletagmanager.com/gtag/js?id=` + code + `"></script>
		<script>
			window.dataLayer = window.dataLayer || [];
			function gtag(){dataLayer.push(arguments);}
			gtag('js', new Date());
		
			gtag('config', 'UA-134084549-1');
		</script>
	`)
}

//Launch launches the app and attempts to open it with a local browser if possible.
//This method is suitable for launching the app in development and in production.
//However, to actually launch the app in production, a -production flag needs to passed to the app.
func (app *App) Launch(listen ...string) error {

	inbed.File("assets")

	if err := inbed.Done(); err != nil {
		log.Println(err)
	}

	app.build()

	if exporting {
		app.Export(Website)
		return nil
	}

	var wasm bool
	for _, arg := range os.Args {
		if arg == "-wasm" {
			wasm = true
		}
	}

	Runtime{
		app:           *app,
		bootstrapWasm: wasm,
	}.Launch(listen...)
	return nil
}

//Handler returns a handler for this app.
func (app *App) Handler() http.Handler {
	app.build()
	return Runtime{app: *app}.Handler()
}

//OnUpdateFound will be called when an update is found for the app.
func (app *App) OnUpdateFound(f func(script.Ctx)) {
	app.onupdatefound = f
}

//While allows a persistent user-connection to be opened when a state is active.
func (app *App) While(state State, handler user.Handler) {
	app.When(state, func(q script.Ctx) {
		q.Open(handler)
	})
}
