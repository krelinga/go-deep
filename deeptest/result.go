package deeptest

type Result interface {
	Log(...any) Result
	Logf(string, ...any) Result
	Ok() bool
	Must()
}

func NewResult(ok bool, t T) Result {
	return &resultImpl{
		ok:  ok,
		t: t,
	}
}

type resultImpl struct {
	ok  bool
	t T
}

func (r *resultImpl) Log(args ...any) Result {
	if !r.ok {
		r.t.Log(args...)
	}
	return r
}

func (r *resultImpl) Logf(format string, args ...any) Result {
	if !r.ok {
		r.t.Logf(format, args...)
	}
	return r
}

func (r *resultImpl) Ok() bool {
	return r.ok
}

func (r *resultImpl) Must() {
	if !r.ok {
		r.t.FailNow()
	}
}
