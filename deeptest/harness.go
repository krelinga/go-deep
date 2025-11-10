package deeptest

import (
	"context"
	"time"

	"github.com/krelinga/go-deep"
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

	Children map[string]*Log
}

func (l *Log) Find(names ...string) (*Log, bool) {
	if l.Children == nil {
		return nil, false
	}
	cur := l
	for _, name := range names {
		child, ok := cur.Children[name]
		if !ok {
			return nil, false
		}
		cur = child
	}
	return cur, true
}

type Harness struct {
	DeepEnv  deep.Env
	Ctx      context.Context
	Name     string
	TempDir  string
	Deadline time.Time
}

func (h *Harness) Run(f func(T)) Log {
	panic("Harness Run: not implemented") // TODO: implement
}
