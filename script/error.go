package script

//Error is a Go call error.
type Error struct {
	Q Ctx
	String
	Code Int
}
