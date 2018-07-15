package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/logv/sybil/src/sybil"
	"github.com/pkg/errors"
)

func RunDigestCmdLine() {
	flag.Parse()

	if err := runDigestCmdLine(&sybil.FLAGS); err != nil {
		fmt.Fprintln(os.Stderr, errors.Wrap(err, "digest"))
		os.Exit(1)
	}
}

func runDigestCmdLine(flags *sybil.FlagDefs) error {
	if flags.TABLE == "" {
		flag.PrintDefaults()
		return sybil.ErrMissingTable
	}

	if flags.PROFILE {
		profile := sybil.RUN_PROFILER()
		defer profile.Start().Stop()
	}
	ctx, span := startInitialSpan("digest")
	defer span.End()

	t := sybil.GetTable(flags.TABLE).WithContext(ctx)
	if err := t.LoadTableInfo(); err != nil {
		return err
	}
	return t.DigestRecords()
}
