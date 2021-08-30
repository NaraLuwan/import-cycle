package b

import (
	"github.com/NaraLuwan/import-cycle/a"
)

type B struct {
	Pa *a.A
}

func New(a *a.A) *B {
	return &B{
		Pa: a,
	}
}

func (b *B) DisplayC() {
	b.Pa.Pc.Show()
}
