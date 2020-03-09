package state

type storage string

//Session configures the state to be backed by session storage.
func Session() Option {
	return func(v *Value) {
		v.storage = "sessionStorage"
	}
}
