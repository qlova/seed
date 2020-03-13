package seed

import (
	"encoding/base64"
	"fmt"
	"math/big"
	"os"
	"path/filepath"

	"github.com/qlova/script"
	"github.com/qlova/script/language"
)

//Dir is the working directory of the seed.
var Dir = filepath.Dir(os.Args[0])

const Production = true

//Data is any data associated with a seed.
//You must provide a way for your data to be deleted.
type Data interface {
	Delete(Seed)
}

type data struct {
	used bool

	parent Seed

	children []Any
}

var seeds = make(map[Seed]data)

type options []Option

func (o options) With(more ...Option) options {
	return append(o, more...)
}

func Options(options ...Option) options {
	return options
}

//Option can be used to modify a seed.
type Option interface {
	AddTo(Any)
	Apply(Ctx)
	Reset(Ctx)

	And(...Option) Option
}

type option struct {
	addto func(Any)

	apply, reset func(Ctx)
}

func NewOption(addto func(Any), apply, reset func(Ctx)) Option {
	return option{
		addto: addto,
		apply: apply,
		reset: reset,
	}
}

//AddTo implements Option.AddTo
func (o option) AddTo(seed Any) {
	o.addto(seed)
}

//AddTo implements Option.AddTo
func (o option) Apply(seed Ctx) {
	o.apply(seed)
}

//Reset implements Option.Reset
func (o option) Reset(seed Ctx) {
	o.reset(seed)
}

func And(o Option, more ...Option) Option {
	return option{
		addto: func(any Any) {
			o.AddTo(any)
			for _, o := range more {
				o.AddTo(any)
			}
		},
		apply: func(ctx Ctx) {
			o.Apply(ctx)
			for _, o := range more {
				o.Apply(ctx)
			}
		},
		reset: func(ctx Ctx) {
			o.Reset(ctx)
			for _, o := range more {
				o.Reset(ctx)
			}
		},
	}
}

//And implements Option.And
func (o option) And(more ...Option) Option {
	return And(o, more...)
}

//Any is a component of your app.
type Any interface {
	Root() Seed
	Add(...Option)
	Option
}

type Seed int

func (s Seed) Use() {
	var data = seeds[s]
	data.used = true
	seeds[s] = data
}

func (s Seed) Used() bool {
	data := seeds[s]
	return data.used
}

func (s Seed) Root() Seed {
	return s
}

func (s Seed) Parent() Seed {
	data := seeds[s]
	return data.parent
}

func (s Seed) Children() []Any {
	data := seeds[s]
	return data.children
}

func (s Seed) Add(options ...Option) {
	for _, o := range options {
		o.AddTo(s)
	}
}

func (s Seed) AddTo(other Any) {
	if s == 0 {
		panic("seed must not be 0")
	}
	if s == other {
		panic("child must not contain itself")
	}
	data := seeds[other.Root()]
	data.children = append(data.children, s)
	seeds[other.Root()] = data

	data = seeds[s.Root()]
	data.parent = other.Root()
	seeds[s.Root()] = data
}

var Apply = func(s Seed, other Ctx) {
	panic("cannot apply a seed on another seed")
}

func (s Seed) Apply(other Ctx) {
	Apply(s, other)
}

var Reset = func(s Seed, other Ctx) {
	panic("cannot apply a seed on another seed")
}

func (s Seed) Reset(other Ctx) {
	Reset(s, other)
}

func (s Seed) And(more ...Option) Option {
	return option{
		addto: func(any Any) {
			s.AddTo(any)
			for _, o := range more {
				o.AddTo(any)
			}
		},
		apply: func(ctx Ctx) {
			panic("cannot apply a seed on another seed")
			for _, o := range more {
				o.Apply(ctx)
			}
		},
		reset: func(ctx Ctx) {
			panic("cannot reset a seed on another seed")
			for _, o := range more {
				o.Reset(ctx)
			}
		},
	}
}

type Ctx struct {
	root Seed
	script.Native
}

func (s Seed) Ctx(q script.AnyCtx) Ctx {
	var id = base64.RawURLEncoding.EncodeToString(big.NewInt(int64(s)).Bytes())
	return Ctx{s, script.Native{Type: language.Expression(q, fmt.Sprintf(`seed.get("%v")`, id))}}
}

func (s Ctx) Element() string {
	s.root.Use()
	return s.Ctx.Raw(s.Native)
}

func (s Ctx) Root() Seed {
	return s.root
}

var id Seed

//New returns a new seed with the applied options.
func New(options ...Option) Seed {
	id++
	var seed = id
	for _, o := range options {
		o.AddTo(seed)
	}
	return seed
}

//If applies the options if the condition is true.
func If(condition bool, options ...Option) Option {
	return NewOption(func(any Any) {
		if condition {
			for _, o := range options {
				o.AddTo(any)
			}
		}
	}, func(s Ctx) {
		if condition {
			for _, o := range options {
				o.Apply(s)
			}
		}
	}, func(s Ctx) {
		if condition {
			for _, o := range options {
				o.Reset(s)
			}
		}
	})
}

//Do runs a function with the seed scoped as the first argument.
func Do(f func(c Seed)) Option {
	return NewOption(func(c Any) {
		f(c.Root())
	}, func(c Ctx) {
		f(c.Root())
	}, func(c Ctx) {
		f(c.Root())
	})
}
