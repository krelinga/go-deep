package deeptest

type Result interface {
	Log(...any) Result
	Logf(string, ...any) Result
	Ok() bool
	Must()
}

type resultImpl struct {
	ok  bool
	env Env
}

func (r *resultImpl) Log(args ...any) Result {
	if !r.ok {
		r.env.Log(args...)
	}
	return r
}

func (r *resultImpl) Logf(format string, args ...any) Result {
	if !r.ok {
		r.env.Logf(format, args...)
	}
	return r
}

func (r *resultImpl) Ok() bool {
	return r.ok
}

func (r *resultImpl) Must() {
	if !r.ok {
		r.env.FailNow()
	}
}
