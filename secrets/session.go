package secrets

import (
	"encoding/base64"
	"encoding/binary"
	"math/big"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/qlova/seed/user"
)

var intranet, _ = regexp.Compile(`(^192\.168\.([0-9]|[0-9][0-9]|[0-2][0-5][0-5])\.([0-9]|[0-9][0-9]|[0-2][0-5][0-5]):.*$)`)

func isLocal(r *http.Request) (local bool) {
	local = strings.Contains(r.RemoteAddr, "[::1]") || strings.Contains(r.RemoteAddr, "127.0.0.1")
	if local {
		return
	}
	if intranet.Match([]byte(r.RemoteAddr)) {
		local = true
	}

	split := strings.Split(r.Host, ":")
	if len(split) == 0 {
		local = false
	} else {
		if split[0] != "localhost" {
			local = false
		}
	}

	return
}

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

//Secret is a server-side session string that is securely stored and encrypted within a session cookie.
type Secret Value

//NewSecret returns a new Secret.
func New(name ...string) Secret {
	return Secret(newValue(name...))
}

//For gets the session String value for the specified user.
func (s Secret) For(u user.Ctx) string {

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
func (s Secret) SetFor(u user.Ctx, value string) {
	var buffer [8]byte
	binary.PutVarint(buffer[:], time.Now().Unix())

	var data = append(buffer[:], value...)
	var cookie string
	if value != "" {
		cookie = Encrypt(data)
	}
	http.SetCookie(u.ResponseWriter(), &http.Cookie{
		Name:     s.string,
		Value:    cookie,
		Secure:   !isLocal(u.Request()),
		HttpOnly: true,
		Expires:  time.Now().Add(time.Hour * 24 * 30),
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
	})
}
