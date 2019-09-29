package user

//Update describes a JSON update that gets sent to the user, it defines a change in state that will be applied to the user.
type Update struct {

	//a mapping of document paths to values,
	// eg. #gg.style.display => hidden;
	Document map[string]string

	//a mapping of local-storage updates.
	LocalStorage map[string]string

	//A list of js statements that will be evaluated.
	Evaluations map[string][]string
}
