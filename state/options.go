package state

type storage string

//Session configures the state to be backed by session storage.
func Session() Option {
	return func(v *Value) {
		v.storage = "sessionStorage"
	}
}

//Global configures the state to be backed by a global variable.
func Global() Option {
	return func(v *Value) {
		v.storage = "seed.storage"
	}
}

//Local configures the state to be backed by a local variable.
func Local() Option {
	return func(v *Value) {
		v.storage = ""
	}
}

//ReadOnly configures the state to be readonly.
func ReadOnly() Option {
	return func(v *Value) {
		v.ro = true
	}
}

//SetKey sets the state's key.
func SetKey(key string) Option {
	return func(v *Value) {
		v.key = key
	}
}
