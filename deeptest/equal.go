package deeptest

import "github.com/krelinga/go-deep/match"

func Equal[T any](env Env, a, b T) Result {
	m := match.Equal(b)
	return Match(env, a, m)
}

func NotEqual[T any](env Env, a, b T) Result {
	m := match.NotEqual(b)
	return Match(env, a, m)
}