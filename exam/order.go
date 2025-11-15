package exam

import (
	"github.com/krelinga/go-libs/deep"
	"github.com/krelinga/go-libs/match"
)

func LessThan[Type any](e E, env deep.Env, a, b Type) Result {
	e.Helper()
	m := match.LessThan(b)
	return Match(e, env, a, m)
}

func GreaterThan[Type any](e E, env deep.Env, a, b Type) Result {
	e.Helper()
	m := match.GreaterThan(b)
	return Match(e, env, a, m)
}

func LessThanOrEqual[Type any](e E, env deep.Env, a, b Type) Result {
	e.Helper()
	m := match.LessThanOrEqual(b)
	return Match(e, env, a, m)
}

func GreaterThanOrEqual[Type any](e E, env deep.Env, a, b Type) Result {
	e.Helper()
	m := match.GreaterThanOrEqual(b)
	return Match(e, env, a, m)
}
