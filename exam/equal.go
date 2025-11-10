package exam

import (
	"github.com/krelinga/go-deep/match"
)

func Equal[Type any](t E, a, b Type) Result {
	t.Helper()
	m := match.Equal(b)
	return Match(t, a, m)
}

func NotEqual[Type any](t E, a, b Type) Result {
	t.Helper()
	m := match.NotEqual(b)
	return Match(t, a, m)
}
