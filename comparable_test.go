package xx

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestEq(t *testing.T) {
	require.True(t, Eq(1)(1))
	require.False(t, Eq(1)(2))
}

func TestNeq(t *testing.T) {
	require.False(t, Neq(1)(1))
	require.True(t, Neq(1)(2))
}
