package seed

//Interface is anything that has a Root() seed method.
type Interface interface {
	Root() Seed
}
