package js

//Set is a javascript Set.
type Set struct {
	Value
}

//AnySet is anything that can retrieve a string.
type AnySet interface {
	AnyValue
	GetSet() Set
}
