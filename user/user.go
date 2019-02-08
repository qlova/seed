package user

import (
	"fmt"
	"encoding/json"
	"net/http"
	"encoding/base64"
	"math/big"
)

type User struct {
	user

	indices []int
	marker int
}

type user struct {
	http.ResponseWriter
	*http.Request
}
	
func (user User) WriteString(s string) {
	user.user.ResponseWriter.Write([]byte(s))
}

func (User) FromHandler(w http.ResponseWriter, r *http.Request) User {
	return User{user:user{
		Request: r,
		ResponseWriter: w, 
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

var id int64 = 1;

type Data string

func DataType() Data {
	//global identification is compressed to base64 and prefixed with g_.
	var result = "user_"+base64.RawURLEncoding.EncodeToString(big.NewInt(id).Bytes())

	id++

	return Data(result)
}