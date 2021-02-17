package clientrpc

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"reflect"
	"strings"
	"time"

	"github.com/spf13/cast"
)

//Scan attempts to convert the given value into the given type.
//Returns the value or an error the operation failed.
func Scan(into reflect.Type, value interface{}) (reflect.Value, error) {
	if into.Implements(reflect.TypeOf([0]Scanner{}).Elem()) && into.Kind() == reflect.Ptr {
		rvalue := reflect.New(into.Elem())

		if err := rvalue.Interface().(Scanner).Scan(value); err != nil {
			return reflect.Value{}, err
		}

		return rvalue, nil
	}

	switch into.Kind() {
	case reflect.String:
		v, err := cast.ToStringE(value)
		if into != reflect.TypeOf("") {
			var result = reflect.New(into).Elem()
			result.SetString(v)
			return result, err
		}
		return reflect.ValueOf(v), err

	case reflect.Interface:
		if reflect.TypeOf(value).Implements(into) {
			return reflect.ValueOf(value), nil
		}
		return reflect.ValueOf(value), fmt.Errorf("clientrpc.Scan: does not implement interface")
	}

	switch into {
	case reflect.TypeOf(time.Time{}):
		if i, err := cast.ToInt64E(value); err == nil {
			val := time.Millisecond * time.Duration(i)
			return reflect.ValueOf(time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC).Add(val)), nil
		}

		v, err := cast.ToTimeE(value)
		if err != nil {
			return reflect.Value{}, err
		}
		return reflect.ValueOf(v), nil

	case reflect.TypeOf(url.URL{}):
		v, err := cast.ToStringE(value)
		if err != nil {
			return reflect.Value{}, err
		}
		location, err := url.Parse(v)
		if err != nil {
			return reflect.Value{}, err
		}
		return reflect.ValueOf(*location), nil
	}

	var shell = reflect.New(into).Interface()
	val, err := cast.ToStringE(value)
	if err != nil {
		return reflect.Value{}, err
	}
	if err := json.NewDecoder(strings.NewReader(val)).Decode(shell); err == nil {
		return reflect.ValueOf(shell).Elem(), nil
	}

	return reflect.Value{}, errors.New("impossible scan")
}
