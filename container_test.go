package xx

import (
	"github.com/stretchr/testify/require"
	"sort"
	"testing"
)

func TestMap(t *testing.T) {
	s := []string{"aaa", "bb", "c"}
	require.Equal(t, []int{3, 2, 1}, Map(s, func(t string) int {
		return len(t)
	}))
}

func TestSet(t *testing.T) {
	require.Equal(t, map[string]struct{}{
		"hello": {},
		"world": {},
	}, Set([]string{"world", "hello"}))
}

func TestKeys(t *testing.T) {
	keys := Keys(map[string]int{
		"world": 3,
		"hello": 2,
	})
	sort.Strings(keys)
	require.Equal(t, []string{"hello", "world"}, keys)
}

func TestValues(t *testing.T) {
	keys := Values(map[string]int{
		"world": 3,
		"hello": 2,
	})
	sort.Ints(keys)
	require.Equal(t, []int{2, 3}, keys)
}

func TestFilter(t *testing.T) {
	require.Equal(t, []string{"hell", "dull"}, Filter([]string{"hell", "hello", "dull"}, func(s string) bool {
		return len(s)%2 == 0
	}))

	require.Equal(t, []string{"hell", "dull"}, Filter([]string{"hell", "hello", "dull"}, Not(Eq("hello"))))
}

func TestRepeat(t *testing.T) {
	require.Equal(t, []string{"a", "a"}, Repeat(2, "a"))
	require.Equal(t, []int{5, 5}, Repeat(2, 5))
}
