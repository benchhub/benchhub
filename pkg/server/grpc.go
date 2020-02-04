package server

import (
	"context"
	"net"

	"github.com/benchhub/benchhub/bhpb"
	"github.com/benchhub/benchhub/pkg/store"
	"github.com/dyweb/gommon/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

// grpc.go is the grpc server implementation

var _ bhpb.BenchHubServer = (*BenchHubGRPCServer)(nil)

type BenchHubGRPCServer struct {
	cfg  Config
	meta store.Meta
}

func New(cfg Config) (*BenchHubGRPCServer, error) {
	return &BenchHubGRPCServer{
		cfg:  cfg,
		meta: store.NewMetaMem(),
	}, nil
}

func (s *BenchHubGRPCServer) Run(ctx context.Context) error {
	log.Infof("listen on addr %s", s.cfg.Addr)
	lis, err := net.Listen("tcp", s.cfg.Addr)
	if err != nil {
		return errors.Wrapf(err, "error listen on: %s", s.cfg.Addr)
	}
	srv := grpc.NewServer()
	bhpb.RegisterBenchHubServer(srv, s)
	if err := srv.Serve(lis); err != nil {
		return errors.Wrap(err, "error calling grpc server Serve")
	}
	return nil
}

func (s *BenchHubGRPCServer) Ping(ctx context.Context, req *bhpb.PingRequest) (*bhpb.PingResponse, error) {
	p, ok := peer.FromContext(ctx)
	content := req.Content + " pong "
	if ok {
		content += p.Addr.String()
	}
	return &bhpb.PingResponse{
		Content: content,
	}, nil
}

func (s *BenchHubGRPCServer) GoBenchmarkRegisterJob(ctx context.Context, spec *bhpb.GoBenchmarkSpec) (*bhpb.JobRegisterResponse, error) {
	return s.meta.GoBenchmarkRegister(ctx, spec)
}

func (s *BenchHubGRPCServer) GoBenchmarkReportResult(ctx context.Context, result *bhpb.GoBenchmarkReportResultRequest) (*bhpb.ResultReportResponse, error) {
	return s.meta.GoBenchmarkReportResult(ctx, result)
}
