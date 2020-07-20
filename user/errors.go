package user

import (
	"errors"
	"fmt"
	"strconv"

	"qlova.org/seed/js"
)

//Cancel returns an invisible error to the client that cancels the current script.
func (u Ctx) Cancel() error {
	return errors.New("")
}

//Report reports the provided error to the user, it should not be used for errors containing sensitive information.
func (u Ctx) Report(err error) {
	u.Execute(js.Script(func(q js.Ctx) {
		q(fmt.Sprintf(`throw %v;`, strconv.Quote(err.Error())))
	}))
}
