package deeptest

import (
	"context"
	"time"
)

type Log struct {
	Error    []string
	Fatal    []string
	Log      []string
	Skip     []string
	Parallel bool
	Fail     bool
	FailNow  bool
	Helper   bool
	SkipNow  bool
}

type Harness struct {
	Ctx      context.Context
	Name     string
	TempDir  string
	Deadline time.Time
}

func (h Harness) Run(f func(T)) Log {
	panic("Harness Run: not implemented") // TODO: implement
}
