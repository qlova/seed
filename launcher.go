package seed

/* 
	A Launcher is used to launch your seed into an application.
	
*/
type Launcher struct {
	Seed

	//Hostname and port where you want the application to be listening on.
	Listen string

	//Here you can pass different seeds to be used for different devices.
	Mobile Seed
	Tablet Seed
}

func (launcher Launcher) Launch() {
	if launcher.Seed.seed != nil {
		if launcher.Listen == "" {
			launcher.Listen = ":1234"
		}
		launcher.Seed.Host(launcher.Listen)
		return
	}
	panic("No seeds!")
}