package session

import (
	"encoding/base64"
	"encoding/binary"
	"math/big"
	"net/http"
	"time"

	"github.com/qlova/seed"
	"github.com/qlova/seed/user"
)

//Value is a session value.
type Value struct {
	string
}

var id int64

//New returns a new globl variable reference.
func newValue(name ...string) Value {
	if len(name) > 0 {
		return Value{"session_" + name[0]}
	}

	//session identification is compressed to base64 and prefixed with s_.
	var result = "s_" + base64.RawURLEncoding.EncodeToString(big.NewInt(id).Bytes())

	id++

	return Value{result}
}

//String is a string that is securely stored within a session cookie.
type String Value

//NewString returns a new session string.
func NewString(name ...string) String {
	return String(newValue(name...))
}

//For gets the session String value for the specified user.
func (s String) For(u user.Ctx) string {

	for _, cookie := range u.Request().Cookies() {
		if cookie.Name == s.string {

			var data = Decrypt(cookie.Value)
			if len(data) < 8 {
				return ""
			}
			var then, _ = binary.Varint(data[0:8])
			if time.Since(time.Unix(then, 0)) > time.Hour*24*30 {
				return ""
			}

			return string(data[8:])
		}
	}
	return ""
}

//SetFor sets the session String value for the specified user.
func (s String) SetFor(u user.Ctx, value string) {
	var buffer [8]byte
	binary.PutVarint(buffer[:], time.Now().Unix())

	var data = append(buffer[:], value...)
	var cookie = Encrypt(data)
	http.SetCookie(u.ResponseWriter(), &http.Cookie{
		Name:     s.string,
		Value:    cookie,
		Secure:   seed.Production,
		HttpOnly: true,
		Expires:  time.Now().Add(time.Hour * 24 * 30),
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
	})
}
