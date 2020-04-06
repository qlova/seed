package js

//Promise is a js promise.
type Promise struct {
	Value
}

//AnyPromise is anything that can return a promise.
type AnyPromise interface {
	AnyValue
	GetPromise() Promise
}
