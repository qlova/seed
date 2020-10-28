//Package clientsafe provides an error type that is safe to show to clients.
package clientsafe

import "fmt"

//Error is a clientsafe error, that is an error that is safe to show to the client.
//You don't want to expose sensitive-information inside of an error message.
type Error interface {
	error

	//ClientError returns an error string that is safe to show the client.
	ClientError() string
}

type err struct {
	error

	safe string
}

//Err returns a new clientsafe.Error that wraps the given internal error.
//format and args are passed to fmt.Sprintf
func Err(internal error, format string, args ...interface{}) Error {
	if internal == nil {
		return nil
	}
	return err{
		internal,
		fmt.Sprintf(format, args...),
	}
}

func (e err) ClientError() string {
	return e.safe
}
