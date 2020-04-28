package seed

import (
	"os"
	"path/filepath"
	"reflect"
)

type Link struct {
	Seed
}

func (link *Link) Link() Option {
	return NewOption(func(c Seed) {
		link.Seed = c
	})
}

func NewLink() *Link {
	return new(Link)
}

type Creator interface {
	New(...Option)
}

//Dir is the working directory of the seed.
var Dir = filepath.Dir(os.Args[0])

const Production = true

//Data is any data associated with a seed.
//You must provide a way for your data to be deleted.
type Data interface {
	data()
}

var id int

type data struct {
	Data

	id int

	used bool

	parent Seed

	children []Seed
}

//Option can be used to modify a seed.
type Option interface {
	AddTo(Seed)
}

type Options []Option

func (options Options) AddTo(c Seed) {
	for _, o := range options {
		o.AddTo(c)
	}
}

func (options Options) And(more ...Option) Option {
	return append(options, more...)
}

//OptionFunc can be used to create an Option.
type OptionFunc func(c Seed)

//NewOption can be used to create options.
type NewOption = OptionFunc

//AddTo implements Option.AddTo
func (o OptionFunc) AddTo(c Seed) {
	o(c)
}

//And implements Option.And
func (o OptionFunc) And(more ...Option) Option {
	return And(o, more...)
}

//And implements Option.And
func And(o Option, more ...Option) Option {
	return NewOption(func(c Seed) {
		o.AddTo(c)
		for _, o = range more {
			o.AddTo(c)
		}
	})
}

//Seed is a generic reference component, 'everything is a seed'.
//Like an enitity in a ECS, other packages can associate Data with this reference.
type Seed interface {
	Option

	ID() int
	Use()
	Used() bool

	Read(Data)
	Write(Data)

	Parent() Seed
	Children() []Seed

	Add(...Option) Option
}

type seed map[reflect.Type]reflect.Value

func (c seed) seed() seed {
	return c
}

func (c seed) Read(d Data) {
	t := reflect.TypeOf(d).Elem()
	if v, ok := c[t]; ok {
		reflect.ValueOf(d).Elem().Set(v)
		return
	}
	reflect.ValueOf(d).Elem().Set(reflect.Zero(t))
}

func (c seed) Write(d Data) {
	t := reflect.TypeOf(d)
	if t.Kind() == reflect.Ptr {
		panic("do not pass pointer to seed.Seed.Write")
	}
	c[t] = reflect.ValueOf(d)
}

func (c seed) ID() int {
	var d data
	c.Read(&d)
	return d.id
}

func (c seed) Parent() Seed {
	var d data
	c.Read(&d)
	return d.parent
}

func (c seed) Use() {
	var d data
	c.Read(&d)
	d.used = true
	c.Write(d)
}

func (c seed) Used() bool {
	var d data
	c.Read(&d)
	return d.used
}

func (c seed) Children() []Seed {
	var d data
	c.Read(&d)
	return d.children
}

func (c seed) Add(options ...Option) Option {
	for _, o := range options {
		o.AddTo(c)
	}
	return NewOption(func(Seed) {})
}

func Add(a, b Seed) {
	var d data
	b.Read(&d)
	d.children = append(d.children, a)
	b.Write(d)

	a.Read(&d)
	d.parent = b
	a.Write(d)
}

func (c seed) AddTo(other Seed) {
	Add(c, other)
}

func (c seed) And(more ...Option) Option {
	return And(c, more...)
}

//New returns a new seed with the applied options.
func New(options ...Option) Seed {
	c := make(seed)

	id++
	var d data
	c.Read(&d)
	d.id = id
	c.Write(d)

	for _, o := range options {
		o.AddTo(c)
	}

	return c
}

//If applies the options if the condition is true.
func If(condition bool, options ...Option) Option {
	return NewOption(func(c Seed) {
		if condition {
			for _, o := range options {
				o.AddTo(c)
			}
		}
	})
}

//Do runs a function with the seed scoped as the first argument.
func Do(f func(c Seed)) Option {
	return NewOption(func(c Seed) {
		f(c)
	})
}

//Apply runs a function with the seed scoped as the first argument.
func Scope(f func(c Seed) Option) Option {
	return NewOption(func(c Seed) {
		f(c).AddTo(c)
	})
}
