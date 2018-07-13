package sybil

import (
	"context"

	"go.opencensus.io/trace"
)

func nop() {}

var ctxTODO = context.TODO()

func (t *Table) trace(ctxs ...context.Context) (context.Context, func()) {
	ctx := t.ctx
	if len(ctxs) == 1 {
		ctx = ctxs[0]
	}
	if ctx == nil || ctx == ctxTODO {
		return ctxTODO, nop
	}
	ctx, span := trace.StartSpan(ctx, getCallerName(3))
	return ctx, span.End
}
