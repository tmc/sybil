package cmd

import (
	"flag"
	"fmt"
	"os"

	"github.com/logv/sybil/src/sybil"
)

func askConfirmation() bool {

	var response string
	_, err := fmt.Scanln(&response)
	if err != nil {
		sybil.Error(err)
	}

	if response == "Y" {
		return true
	}

	if response == "N" {
		return false
	}

	fmt.Println("Y or N only")
	return askConfirmation()

}

func RunTrimCmdLine() {
	MB_LIMIT := flag.Int("mb", 0, "max table size in MB")
	DELETE_BEFORE := flag.Int("before", 0, "delete blocks with data older than TIMESTAMP")
	DELETE := flag.Bool("delete", false, "delete blocks? be careful! will actually delete your data!")
	REALLY := flag.Bool("really", false, "don't prompt before deletion")

	flag.StringVar(&sybil.FLAGS.TIME_COL, "time-col", "", "which column to treat as a timestamp [REQUIRED]")
	flag.Parse()

	if sybil.FLAGS.TABLE == "" || sybil.FLAGS.TIME_COL == "" {
		flag.PrintDefaults()
		return
	}

	if sybil.FLAGS.PROFILE {
		profile := sybil.RUN_PROFILER()
		defer profile.Start().Stop()
	}

	t := sybil.GetTable(sybil.FLAGS.TABLE)
	if !t.LoadTableInfo() {
		sybil.Warn("Couldn't read table info, exiting early")
		return
	}

	loadSpec := t.NewLoadSpec()
	loadSpec.SkipDeleteBlocksAfterQuery = true
	loadSpec.Int(sybil.FLAGS.TIME_COL)

	trimSpec := sybil.TrimSpec{}
	trimSpec.DeleteBefore = int64(*DELETE_BEFORE)
	trimSpec.MBLimit = int64(*MB_LIMIT)

	toTrim := t.TrimTable(&trimSpec)

	sybil.Debug("FOUND", len(toTrim), "CANDIDATE BLOCKS FOR TRIMMING")
	if len(toTrim) > 0 {
		for _, b := range toTrim {
			fmt.Println(b.Name)
		}
	}

	if *DELETE {
		if !*REALLY {
			// TODO: prompt for deletion
			fmt.Println("DELETE THE ABOVE BLOCKS? (Y/N)")
			if !askConfirmation() {
				sybil.Debug("ABORTING")
				return
			}

		}

		sybil.Debug("DELETING CANDIDATE BLOCKS")
		for _, b := range toTrim {
			sybil.Debug("DELETING", b.Name)
			if len(b.Name) > 5 {
				os.RemoveAll(b.Name)
			} else {
				sybil.Debug("REFUSING TO DELETE", b.Name)
			}
		}

	}
}
