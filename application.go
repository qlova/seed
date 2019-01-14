package seed

type Application struct {
	Seed

	
}

func NewApp() Application {
	return Application {
		Seed: New(),
	}
}
