package c

import "github.com/NaraLuwan/import-cycle/a"

type C struct {
	Vc int
}

func New(i int) *C {
	return &C{
		Vc: i,
	}
}

func (c *C) Show() {
	a.Printf(c.Vc)
}
