package style

import "github.com/qlova/seed/style/css"

type composition struct {
	A, B css.Stylable
}

func (c composition) Set(property, value string) {
	c.A.Set(property, value)
	c.B.Set(property, value)
}

func (c composition) Get(property string) (value string) {
	return c.A.Get(property)
}

func (c composition) Bytes() []byte {
	return c.A.Bytes()
}

//Compose returns a new style that writes to both input styles.
func Compose(a, b Style) Style {
	return From(composition{
		A: a.Style.Stylable,
		B: b.Style.Stylable,
	})
}
