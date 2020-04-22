package user

import (
	"fmt"
	"strconv"

	"github.com/qlova/seed/js"
)

//Report reports the provided error to the user, it should not be used for errors containing sensitive information.
func (u Ctx) Report(err error) {
	u.Execute(func(q js.Ctx) {
		q(fmt.Sprintf(`throw %v;`, strconv.Quote(err.Error())))
	})
}
