package seed

//Runtime is a runtime that can serve an app.
type Runtime struct {
	app App

	bootstrapWasm bool

	//Hostname and port where you want the server to be listening on.
	Listen string
}
