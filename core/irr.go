package core

type WrapF struct {
	f func()

	ctx *Context
}

func WrapFuncWithContext(f func(), c *Context) *WrapF {
	return &WrapF{
		f:   f,
		ctx: c,
	}
}

func (wf *WrapF) Use(f func(c *Context)) {
	wf.ctx.addHandler(f)
}

func (wf *WrapF) Handle() {
	wf.ctx.handlers = append(wf.ctx.handlers, func(c *Context) {
		wf.f()
	})

	if len(wf.ctx.handlers) > 0 {
		wf.ctx.Next()
	}
	wf.ctx.Reset()
}
