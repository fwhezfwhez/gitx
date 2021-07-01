package core

import (
	"math"
	"os"
	"strings"
)

const ABORT = math.MaxInt32 - 10000

type Context struct {
	Args       []string
	Command    string
	SubCommand string

	OsArgs []string

	offset   int
	handlers []func(*Context)
}

func NewContext(args []string) *Context {
	return &Context{
		offset:     -1,
		handlers:   make([]func(*Context), 0, 10),
		Command:    args[0],
		SubCommand: args[1],
		Args:       args,

		OsArgs: os.Args,
	}
}
func (ctx *Context) Next() {
	ctx.offset ++
	s := len(ctx.handlers)
	for ; ctx.offset < s; ctx.offset++ {
		if !ctx.isAbort() {
			ctx.handlers[ctx.offset](ctx)
		} else {
			return
		}
	}
}
func (ctx *Context) Reset() {
	//ctx.PerRequestContext = &sync.Map{}
	ctx.offset = -1
	ctx.handlers = ctx.handlers[:0]
}

// stop middleware chain
func (ctx *Context) Abort() {
	ctx.offset = math.MaxInt32 - 10000
}

func (ctx *Context) isAbort() bool {
	if ctx.offset >= ABORT {
		return true
	}
	return false
}

func (ctx *Context) addHandler(f func(ctx *Context)) {
	ctx.handlers = append(ctx.handlers, f)
}

// 0 - gitx
// 1 - branch/merge/add/push
// 2 - dev/origin
func (ctx *Context) Index(i int) string {
	return GetValueOfIndex(ctx.Args, i)
}

func (ctx *Context) RawCommand() string {
	return strings.Join(ctx.Args, " ")
}
