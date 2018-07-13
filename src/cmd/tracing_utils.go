package cmd

import (
	"bytes"
	"encoding/hex"
	"os"

	"go.opencensus.io/trace"
	"golang.org/x/net/context"
)

func startInitialSpan(name string) (context.Context, *trace.Span) {
	ctx := context.Background()
	var span *trace.Span
	var traceID trace.TraceID
	var spanID trace.SpanID
	t, _ := hex.DecodeString(os.Getenv("TRACE_ID"))
	s, _ := hex.DecodeString(os.Getenv("SPAN_ID"))
	copy(traceID[:], t[:])
	copy(spanID[:], s[:])

	if bytes.Equal(traceID[:], bytes.Repeat([]byte{0}, len(traceID))) {
		ctx, span = trace.StartSpan(ctx, name)
		traceID, spanID = span.SpanContext().TraceID, span.SpanContext().SpanID
	} else {
		ctx, span = trace.StartSpanWithRemoteParent(ctx, name, trace.SpanContext{TraceID: trace.TraceID(traceID), SpanID: trace.SpanID(spanID)})
	}
	return ctx, span
}
