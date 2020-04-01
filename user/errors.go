package user

import (
	"fmt"
	"strconv"
)

//Report reports the provided error to the user, it should not be used for errors containing sensitive information.
func (u Ctx) Report(err error) {
	u.Execute(fmt.Sprintf(`throw %v;`, strconv.Quote(err.Error())))
}
