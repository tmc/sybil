package cmd

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/logv/sybil/src/sybil"
	"github.com/logv/sybil/src/sybild"
	pb "github.com/logv/sybil/src/sybilpb"
	"github.com/pkg/errors"
	"go.opencensus.io/examples/exporter"
	"go.opencensus.io/plugin/ocgrpc"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/zpages"
	"google.golang.org/grpc"
)

const defaultServeListenAddr = "localhost:7000"
const defaultDebugListenAddr = "localhost:7001"

func RunServeCmdLine() {
	flag.Parse()
	if err := runServeCmdLine(&sybil.FLAGS); err != nil {
		fmt.Fprintln(os.Stderr, errors.Wrap(err, "serve"))
		os.Exit(1)
	}
}

func runServeCmdLine(flags *sybil.FlagDefs) error {
	//ctx := context.Background()
	// TODO: handle signals, shutdown
	// TODO: auth, tls
	go func() {
		mux := http.NewServeMux()
		zpages.Handle(mux, "/debug")
		log.Fatal(http.ListenAndServe(defaultDebugListenAddr, mux))
	}()

	// Register stats and trace exporters to export
	// the collected data.
	view.RegisterExporter(&exporter.PrintExporter{})

	_, span := startInitialSpan("serve")
	span.End() // immediate stop span as a long running one won't be very interesting
	// Register the views to collect server request count.
	if err := view.Register(ocgrpc.DefaultServerViews...); err != nil {
		log.Fatal(err)
	}

	lis, err := net.Listen("tcp", defaultServeListenAddr)
	if err != nil {
		return err
	}
	s, err := sybild.NewServer()
	if err != nil {
		return err
	}
	grpcServer := grpc.NewServer(grpc.StatsHandler(&ocgrpc.ServerHandler{}))
	pb.RegisterSybilServer(grpcServer, s)
	return grpcServer.Serve(lis)
}
