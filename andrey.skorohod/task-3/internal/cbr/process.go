package cbr

import (
	"cmp"
	"slices"
)

func (val *ValCurs) ProcessData() {
	slices.SortFunc(val.Valutes, func(a, b Valute) int {
		return cmp.Compare(b.Value, a.Value)
	})
}
