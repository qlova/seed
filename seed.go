package seed

import (
	"reflect"
	"sync/atomic"
)

// id tracker, must be incremented atomically.
var id int64

type number int64

//Node is an extendable entity type. Other packages can associate properties with a seed node.
type Node struct {
	ptr *data
}

// Children returns a slice copy of the node's children.
func (n Node) Children() []Node {

	slice := Slice[child](n)

	var children = make([]Node, len(slice))
	for i, child := range slice {
		children[i] = Node(child)
	}
	return children
}

type data map[reflect.Type]interface{}

// Property of a Node.
type Property interface {
	addTo(Node)
}

//New returns a new seed with the applied options.
func New(properties ...Property) Node {
	var node Node

	data := make(data)
	node.ptr = &data

	Set(number(atomic.AddInt64(&id, 1))).addTo(node)

	for _, prop := range properties {
		if prop != nil {
			prop.addTo(node)
		}
	}

	return node
}

type child Node
type parent Node

//AddTo implements Option.
func (node Node) addTo(other Node) {
	children := Slice[child](node)

	//Don't re-add children.
	if Node(Get[parent](node)) == other {
		return
	}

	if node == other {
		panic("child added to itself")
	}

	for _, c := range children {
		if Node(c) == node {
			return
		}
	}

	Append(child(node)).addTo(other)
	Set(parent(other)).addTo(node)
}

type property func(Node)

func (p property) addTo(c Node) {
	p(c)
}

// Set sets the T typed property value.
func Set[T any](value T) Property {
	return property(func(node Node) {
		data := *node.ptr
		data[reflect.TypeOf(value)] = value
	})
}

// Append appends the value to the []T typed property slice.
func Append[T any](value T) Property {
	return property(func(node Node) {
		data := *node.ptr
		var rtype []T

		key := reflect.TypeOf(rtype)

		slice, ok := data[key].([]T)
		if !ok {
			slice = make([]T, 0, 1)
		}
		data[key] = append(slice, value)
	})
}

// Get the T typed property from the node.
func Get[T any](node Node) T {
	var elem T
	data := *node.ptr
	if value, ok := data[reflect.TypeOf(elem)]; ok {
		return value.(T)
	}
	return elem
}

// Slice returns the []T typed property from the node.
func Slice[T any](node Node) []T {
	var elem []T
	data := *node.ptr

	if value, ok := data[reflect.TypeOf(elem)]; ok {
		return value.([]T)
	}
	return elem
}
