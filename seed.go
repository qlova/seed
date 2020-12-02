package seed

import (
	"os"
	"path/filepath"
	"reflect"
	"sort"
)

//data associated with a seed by this package.
type data struct {
	used bool

	parent   Seed
	children []Seed
}

//Seed is an extendable entity type. Other packages can associate data with a seed.
//Seeds can have options applied to them and are also options themselves (seeds can be added to seeds).
type Seed struct {
	id   int
	data map[reflect.Type]reflect.Value
}

//New returns a new seed with the applied options.
func New(options ...Option) Seed {
	var c Seed

	id++
	c.id = id
	c.data = make(map[reflect.Type]reflect.Value)

	for _, o := range options {
		if o != nil {
			o.AddTo(c)
		}
	}

	return c
}

//Load loads the associated data of data's type into data and then returns true.
//If no data is found, data is set to the empty value and false is returned instead.
func (c Seed) Load(data interface{}) bool {
	t := reflect.TypeOf(data).Elem()
	if v, ok := c.data[t]; ok {
		reflect.ValueOf(data).Elem().Set(v)
		return true
	}
	reflect.ValueOf(data).Elem().Set(reflect.Zero(t))
	return false
}

//Save saves data of the given type to the seed.
func (c Seed) Save(data interface{}) {
	t := reflect.TypeOf(data)
	v := reflect.ValueOf(data)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}
	c.data[t] = v
}

//ID returns the Seed's ID.
func (c Seed) ID() int {
	return c.id
}

//Parent returns the seeds parent.
func (c Seed) Parent() Seed {
	var d data
	c.Load(&d)
	return d.parent
}

//Use marks the seed as used.
func (c Seed) Use() {
	var d data
	c.Load(&d)
	d.used = true
	c.Save(d)
}

//Used reports whether or not the seed has been used.
func (c Seed) Used() bool {
	var d data
	c.Load(&d)
	return d.used
}

//Children returns the seeds children.
func (c Seed) Children() []Seed {
	var d data
	c.Load(&d)
	return d.children
}

//With is depreciated.
func (c Seed) With(options ...Option) Seed {
	for _, o := range options {
		o.AddTo(c)
	}
	return c
}

//Builder is a function that takes one or more options and returns a seed.
type Builder func(...Option) Seed

//Dir is the working directory of the seed.
var Dir = filepath.Dir(os.Args[0])

var id int

//AddTo implements Option.
func (c Seed) AddTo(other Seed) {
	var d data
	other.Load(&d)

	//Don't re-add children.
	if c.Parent().ID() == other.ID() {
		return
	}

	if c.ID() == other.ID() {
		panic("child added to itself")
	}

	for _, child := range d.children {
		if child.ID() == c.ID() {
			return
		}
	}

	d.children = append(d.children, c)
	other.Save(d)

	c.Load(&d)
	d.parent = other
	c.Save(d)
}

//Set returns an unordered set of seeds.
type Set struct {
	mapping map[int]Seed
}

//NewSet creates a new unordered set of seeds out of the given seed arguments.
func NewSet(seeds ...Seed) Set {
	var set Set
	set.mapping = make(map[int]Seed, len(seeds))
	for _, seed := range seeds {
		set.Add(seed)
	}
	return set
}

//Add adds seeds to the set.
func (s *Set) Add(seeds ...Seed) {
	if seeds == nil {
		return
	}
	if s.mapping == nil {
		s.mapping = make(map[int]Seed)
	}
	for _, seed := range seeds {
		if seed.id == 0 {
			continue
		}
		s.mapping[seed.ID()] = seed
	}
}

//Remove removes a seed from the set.
func (s *Set) Remove(c Seed) {
	if c.id == 0 {
		return
	}
	delete(s.mapping, c.ID())
}

//Slice returns a slice of all the seeds in this set, ordered by id.
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
