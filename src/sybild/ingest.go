package sybild

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/davecgh/go-spew/spew"
	"go.opencensus.io/trace"
	context "golang.org/x/net/context"
)

func sybilIngest(ctx context.Context, tableName string, buf io.Reader) error {
	ctx, span := trace.StartSpan(ctx, "sybilIngest")
	defer span.End()
	const sybilBinary = "sybil"
	var sybilFlags = []string{"ingest", "-table", tableName}
	c := exec.Command(sybilBinary, sybilFlags...)
	c.Env = append(os.Environ(),
		fmt.Sprintf("TRACE_ID=%s", span.SpanContext().TraceID),
		fmt.Sprintf("SPAN_ID=%s", span.SpanContext().SpanID),
	)
	c.Stderr = os.Stderr
	si, err := c.StdinPipe()
	if err != nil {
		return err
	}
	so, err := c.StdoutPipe()
	if err != nil {
		return err
	}
	if err := c.Start(); err != nil {
		return err
	}
	io.Copy(si, buf)
	if err := si.Close(); err != nil {
		return err
	}
	if err := c.Wait(); err != nil {
		return err
	}
	results := new(bytes.Buffer)
	io.Copy(results, so)
	spew.Fdump(os.Stderr, results.Bytes())
	return nil
}
