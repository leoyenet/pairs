package pairs

import (
	"fmt"
	"sort"
)

type Pair[K, V comparable] struct {
	First  K
	Second V
}

func New[K, V comparable](p1 K, p2 V) Pair[K, V] {
	return Pair[K, V]{p1, p2}
}

func (p *Pair[K, V]) String() string {
	return fmt.Sprintf("(%v, %v)", p.First, p.Second)
}

func (p *Pair[K, V]) Equal(p2 *Pair[K, V]) bool {
	return p.First == p2.First && p.Second == p2.Second
}

func Sort[K, V comparable](slice []Pair[K, V], less func(i, j int) bool) []Pair[K, V] {
	sort.Slice(slice, less)
	return slice
}
