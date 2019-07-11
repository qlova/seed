package seed

import "fmt"
import "os"
import "github.com/qlova/seed/manifest"
import "github.com/qlova/seed/service"
import "github.com/qlova/seed/style/css"

type App struct {
	Seed
	manifest.Manifest
	service.Worker
	*harvester

	host, rest, pkg, tracking string

	hashes []string

	onupdatefound func(Script)

	loadingPage  Page
	startingPage Page
}

//Create a new application, accepts title and content arguments.
func NewApp(args ...string) *App {
	var app = App{
		Seed:      New(),
		Manifest:  manifest.New(),
		Worker:    service.NewWorker(),
		harvester: newHarvester(),
	}

	app.Seed.app = &app

	app.SetSize(100, 100)
	app.SetDisplay(css.Flex)
	app.Align(0)
	app.SetFlexDirection(css.Column)

	if len(args) > 0 {
		app.SetName(args[0])
	}

	if len(args) > 1 {
		app.SetContent(args[1])
	}

	return &app
}

func (app *App) NewPage() Page {

	if app.loadingPage.Null() {
		app.loadingPage = AddPageTo(app)
		app.loadingPage.SetVisible()
	}

	return AddPageTo(app)
}

func (app *App) SetPage(page Page) {
	app.startingPage = page
	app.loadingPage.OnReady(func(q Script) {
		q.Javascript(`if (window.localStorage.getItem("update")) {`)
		q.Javascript(`window.localStorage.removeItem("update");`)
		q.Javascript(`window.localStorage.removeItem("*CurrentPage");`)
		q.Javascript(`}`)
		q.Javascript(`if (!window.localStorage.getItem("*CurrentPage")) goto("` + page.id + `");`)
	})
}

//Return the loadingpage (like a splashscreen) for this app that displays while the app is loading.
func (app *App) LoadingPage() Page {
	return app.loadingPage
}

//Set the hostname of this app, this is where the app is expected to be hosted from.
func (app *App) SetHost(name string) {
	app.host = name
}

//Set the REST hostname of this app, this is where the app will serve and request API calls.
func (app *App) SetRest(name string) {
	app.rest = name
}

//Set the package name of this application on android.
func (app *App) SetPackage(name string) {
	app.pkg = name
}

//Add a hash of the certificate that you will sign your android app with.
func (app *App) AddHash(name string) {
	app.hashes = append(app.hashes, name)
}

//Set the google analytics tracking code.
func (app *App) SetTrackingCode(code string) {
	app.tracking = code
}

//TODO random port, can be set with enviromental variables.
func (app *App) Launch(listen ...string) error {

	if len(os.Args) == 2 && os.Args[1] == "-deploy" {
		err := app.Deploy()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		os.Exit(0)
	}

	launcher{App: *app}.Launch(listen...)
	return nil
}

func (app *App) OnUpdateFound(f func(Script)) {
	app.onupdatefound = f
}
