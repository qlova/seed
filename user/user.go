package user

import (
	"bytes"
	"encoding/json"
	"net/http"
	"regexp"
)

//User is a current user of the app.
type User struct {
	user
}

//Production specifies if we are running in production.
var Production bool

type user struct {
	http.ResponseWriter
	*http.Request

	//The pending update for the user.
	Update

	indices []int
	marker  int
}

//SetIndices sets the indicies of a feed request.
func (user *User) SetIndices(i []int) {
	user.indices = i
	user.marker = 0
}

//Index returns the current index.
func (user User) Index() int {
	if user.marker < len(user.indices) {
		user.marker++
		return user.indices[user.marker-1]
	} else {
		return -1
	}
}

//WriteString writes a string to the user.
func (user User) WriteString(s string) {
	user.user.ResponseWriter.Write([]byte(s))
}

//FromHandler returns a user from http.Handler arguments.
func (User) FromHandler(w http.ResponseWriter, r *http.Request) User {
	return User{user: user{
		Request:        r,
		ResponseWriter: w,

		Update: Update{
			Document:     make(map[string]string),
			LocalStorage: make(map[string]string),
			script:       bytes.NewBuffer(nil),
		},
	}}
}

//Send encodes data as json and sends it to the user.
func (user User) Send(data interface{}) {
	json.NewEncoder(user.ResponseWriter).Encode(data)
}

var intranet *regexp.Regexp

func init() {
	var err error
	intranet, err = regexp.Compile(`(^192\.168\.([0-9]|[0-9][0-9]|[0-2][0-5][0-5])\.([0-9]|[0-9][0-9]|[0-2][0-5][0-5]):.*$)`)
	if err != nil {
		panic("invalid regexp!")
	}
}

//NotAuthorised returns a 401 error to the user.
func (user User) NotAuthorised() {
	user.ResponseWriter.WriteHeader(401)
}

func (user User) Error(err ...string) {
	user.ResponseWriter.WriteHeader(500)
	for _, e := range err {
		user.user.ResponseWriter.Write([]byte(e))
	}
}

//Close closes the user.
func (user User) Close() {
	if len(user.Update.Document) > 0 ||
		len(user.Update.LocalStorage) > 0 ||
		len(user.Update.script.Bytes()) > 0 ||
		len(user.Update.Data) > 0 {
		user.Evaluations = user.Update.script.String()
		json.NewEncoder(user.ResponseWriter).Encode(user.Update)
	} else {
		user.ResponseWriter.WriteHeader(http.StatusOK)
	}
}
