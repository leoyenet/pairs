package pairs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEqual(t *testing.T) {
	first := Pair[int, int]{1, 2}
	second := Pair[int, int]{1, 2}
	assert.True(t, first.Equal(&second) && second.Equal(&first), "it should be equal")
}

func TestString(t *testing.T) {
	p := Pair[int, int]{1, 2}
	assert.Equal(t, p.String(), "(1, 2)", "result should be (1, 2)")
}

func TestSort(t *testing.T) {
	p := []Pair[int, int]{{1, 10}, {2, 9}, {4, 5}}

	t.Run("low to high", func(t *testing.T) {
		assert.Equal(t, []Pair[int, int]{{4, 5}, {2, 9}, {1, 10}}, Sort(p, func(i, j int) bool {
			return p[i].Second < p[j].Second
		}))
	})

	t.Run("high to low", func(t *testing.T) {
		assert.Equal(t, []Pair[int, int]{{1, 10}, {2, 9}, {4, 5}}, Sort(p, func(i, j int) bool {
			return p[i].Second > p[j].Second
		}))
	})
}

func TestNew(t *testing.T) {
	first := New(1, 2)
	second := Pair[int, int]{1, 2}
	assert.Equal(t, first, &second, "they should be equal")
}
