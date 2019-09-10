package user

//Argument is a user provided argument.
type Argument struct {
	string
}

//Args returns the argument with the given name.
func (user User) Args(name string) Argument {
	return Argument{user.Request.FormValue(name)}
}

func (arg Argument) String() string {
	return arg.string
}
