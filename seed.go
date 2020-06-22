package seed

import (
	"os"
	"path/filepath"
	"reflect"
	"sort"
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
		if o != nil {
			o.AddTo(c)
		}
	}
}

func (options Options) And(more ...Option) Options {
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

	With(...Option) Seed
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

func (c seed) With(options ...Option) Seed {
	for _, o := range options {
		o.AddTo(c)
	}
	return c
}

func Add(a, b Seed) {
	var d data
	b.Read(&d)

	//Don't re-add children.
	if parent := a.Parent(); parent != nil {
		if parent.ID() == b.ID() {
			return
		}
	}

	for _, child := range d.children {
		if child.ID() == a.ID() {
			return
		}
	}

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
		if o != nil {
			o.AddTo(c)
		}
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

type Set struct {
	mapping map[int]Seed
}

func NewSet(seeds ...Seed) Set {
	var set Set
	set.mapping = make(map[int]Seed, len(seeds))
	for _, seed := range seeds {
		set.Add(seed)
	}
	return set
}

func (s *Set) Add(seeds ...Seed) {
	if seeds == nil {
		return
	}
	if s.mapping == nil {
		s.mapping = make(map[int]Seed)
	}
	for _, seed := range seeds {
		if seed == nil {
			continue
		}
		s.mapping[seed.ID()] = seed
	}
}

func (s *Set) Remove(c Seed) {
	if c == nil {
		return
	}
	delete(s.mapping, c.ID())
}

//Slice returns an ordered list of seeds by id.
func (s *Set) Slice() []Seed {
	var slice = make([]Seed, 0, len(s.mapping))
	for _, seed := range s.mapping {
		slice = append(slice, seed)
	}

	sort.Slice(slice, func(i, j int) bool {
		return slice[i].ID() < slice[j].ID()
	})

	return slice
}
