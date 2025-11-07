package deeptest

import (
	"context"
	"testing"
	"time"

	"github.com/krelinga/go-deep"
)

type Env interface {
	deep.Env

	// Methods from testing.T
	// TODO: add Attr()
	Chdir(dir string)
	Cleanup(f func())
	Context() context.Context
	Deadline() (deadline time.Time, ok bool)
	Error(args ...any)
	Errorf(format string, args ...any)
	Fail()
	FailNow()
	Failed() bool
	Fatal(args ...any)
	Fatalf(format string, args ...any)
	Helper()
	Log(args ...any)
	Logf(format string, args ...any)
	Name() string
	// TODO: add Output()
	Parallel()
	Setenv(key, value string)
	Skip(args ...any)
	SkipNow()
	Skipf(format string, args ...any)
	Skipped() bool
	TempDir() string
}

type envImpl struct {
	deep.Env
	*testing.T
}

func NewEnv(t *testing.T) Env {
	return &envImpl{
		Env: deep.NewEnv(),
		T:   t,
	}
}