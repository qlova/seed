package seed


type Application struct {
	Seed
}

func NewApp() Application {
	return Application {
		Seed: New(),
	}
}

func (app Application) OnReadyGoto(page Seed) {
	app.Seed.OnReady(func(q Script) {
		//Don't bypass persistence features.
		q.Javascript(`if (!window.localStorage.getItem('*CurrentPage'))`)
		q.Goto(page)
	})
}