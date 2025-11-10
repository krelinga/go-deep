package deeptest

import (
	"github.com/krelinga/go-deep/match"
)

func Match[Type any](t T, got Type, m match.Matcher) Result {
	t.Helper()
	vals := match.NewVals1(got)
	return MatchVals(t, vals, m)
}

func MatchVals(t T, vals match.Vals, m match.Matcher) Result {
	t.Helper()
	r := m.Match(t.DeepEnv(), vals)
	if !r.Matched() {
		t.Errorf("Match failed:\n%s", r)
	}
	return NewResult(r.Matched(), t)
}
