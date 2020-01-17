package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"regexp"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

//User is a current user of the app.
type User struct {
	user
	conn net.Conn
	buff *bytes.Buffer
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

//Writer returns a writer for the user.
func (user User) Writer() io.Writer {
	if user.conn != nil {
		return user.buff
	}
	return user.ResponseWriter
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
	user.Writer().Write([]byte(s))
}

//FromHandler returns a user from http.Handler arguments.
func (User) FromHandler(w http.ResponseWriter, r *http.Request) User {
	return User{user: user{
		Request:        r,
		ResponseWriter: w,

		Update: Update{
			Response:     new(string),
			Document:     make(map[string]string),
			LocalStorage: make(map[string]string),
			script:       new(bytes.Buffer),
		},
	}}
}

//Send encodes data as json and sends it to the user.
func (user User) Send(data interface{}) {
	json.NewEncoder(user.Writer()).Encode(data)
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
	if user.ResponseWriter != nil {
		user.ResponseWriter.WriteHeader(401)
	}
}

func (user User) Error(err ...string) {
	if user.ResponseWriter != nil {
		user.ResponseWriter.WriteHeader(500)
	}
	for _, e := range err {
		user.Writer().Write([]byte(e))
	}
}

//Close closes the user.
func (user User) Close() {
	if len(user.Update.Document) > 0 ||
		len(user.Update.LocalStorage) > 0 ||
		len(user.Update.script.Bytes()) > 0 ||
		len(user.Update.Data) > 0 ||
		len(*user.Update.Response) > 0 {

		user.Evaluation = user.Update.script.String()
		err := json.NewEncoder(user.Writer()).Encode(user.Update)
		if err != nil {
			fmt.Println(err.Error())
		}
	} else if user.ResponseWriter != nil {
		user.ResponseWriter.WriteHeader(http.StatusOK)
	}

	if user.conn != nil {
		err := wsutil.WriteServerMessage(user.conn, ws.OpText, user.buff.Bytes())
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}
