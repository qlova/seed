package user

import (
	"fmt"
	"strconv"
)

//Report reports the provided error to the user, it should not be used for errors containing sensitive information.
func (u Ctx) Report(err error) {
	if u.w != nil {
		u.w.WriteHeader(500)
	}
	u.Execute(fmt.Sprintf(`console.error(%v);`, strconv.Quote(err.Error())))
}

//AreHacking sends an HTTP 400 status code to the user.
func (u Ctx) AreHacking() {
	if u.w != nil {
		u.w.WriteHeader(400)
	}
}

//NeedToLogin sends an HTTP 401 status code to the user.
func (u Ctx) NeedToLogin() {
	if u.w != nil {
		u.w.WriteHeader(401)
	}
}

//NeedToPurchase sends an HTTP 402 status code to the user.
func (u Ctx) NeedToPurchase() {
	if u.w != nil {
		u.w.WriteHeader(402)
	}
}

//AreNotAdmin sends an HTTP 403 status code to the user.
func (u Ctx) AreNotAdmin() {
	if u.w != nil {
		u.w.WriteHeader(403)
	}
}

//AreLost sends an HTTP 404 status code to the user.
func (u Ctx) AreLost() {
	if u.w != nil {
		u.w.WriteHeader(404)
	}
}
