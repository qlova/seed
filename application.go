package seed


type Application struct {
	Seed

	//These booleans enable the application to serve different content depending on the user agent passed to the server.
	Desktop, Mobile, Tablet, Watch bool
}

func NewApp() Application {
	return Application {
		Seed: New(),
	}
}
