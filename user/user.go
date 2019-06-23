package user

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"regexp"
	"time"
)

type User struct {
	user

	indices []int
	marker  int
}

var Production bool

type user struct {
	http.ResponseWriter
	*http.Request

	//The pending update for the user.
	Update
}

func (user User) WriteString(s string) {
	user.user.ResponseWriter.Write([]byte(s))
}

func (User) FromHandler(w http.ResponseWriter, r *http.Request) User {
	return User{user: user{
		Request:        r,
		ResponseWriter: w,

		Update: Update{
			Document:     make(map[string]string),
			LocalStorage: make(map[string]string),
		},
	}}
}

func (user *User) SetIndices(i []int) {
	user.indices = i
	user.marker = 0
}

func (user User) Index() int {
	if user.marker < len(user.indices) {
		user.marker++
		return user.indices[user.marker-1]
	} else {
		return -1
	}
}

func (user User) Send(data interface{}) {
	json.NewEncoder(user.ResponseWriter).Encode(data)
}

func (user User) Get(data Data) string {
	result, err := user.Request.Cookie(string(data))
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	return result.Value
}

var Intranet *regexp.Regexp

func init() {
	var err error
	Intranet, err = regexp.Compile(`(^192\.168\.([0-9]|[0-9][0-9]|[0-2][0-5][0-5])\.([0-9]|[0-9][0-9]|[0-2][0-5][0-5]):.*$)`)
	if err != nil {
		panic("invalid regexp!")
	}
}

func (user User) Set(data Data, value string) {
	http.SetCookie(user.ResponseWriter, &http.Cookie{
		Name:    string(data),
		Value:   value,
		Secure:  Production,
		Expires: time.Now().Add(time.Hour * 24 * 365),
	})
}

func (user User) NotAuthorised() {
	user.ResponseWriter.WriteHeader(401)
}

func (user User) Error(err ...string) {
	user.ResponseWriter.WriteHeader(500)
	for _, e := range err {
		user.user.ResponseWriter.Write([]byte(e))
	}
}

func (user User) Close() {
	if len(user.Update.Document) > 0 || len(user.Update.LocalStorage) > 0 {
		json.NewEncoder(user.ResponseWriter).Encode(user.Update)
	} else {
		user.ResponseWriter.WriteHeader(http.StatusOK)
	}
}

var id int64 = 1

type Data string

func DataType() Data {
	//global identification is compressed to base64 and prefixed with g_.
	var result = "user_" + base64.RawURLEncoding.EncodeToString(big.NewInt(id).Bytes())

	id++

	return Data(result)
}
