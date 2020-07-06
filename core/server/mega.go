package server

import (
	"context"
	"net"

	"github.com/benchhub/benchhub/bhpb"
	"google.golang.org/grpc"
)

// Mega combines all the services.
// We have multiple services that looks like microservice outside.
// But it's a single process inside. Mega (R) Inside
type Mega struct {
	addr string
	srv  *grpc.Server
}

func newMega(addr string) (*Mega, error) {
	srv := grpc.NewServer()
	userSvc, err := newUserService()
	if err != nil {
		return nil, err
	}
	bhpb.RegisterUserServiceServer(srv, userSvc)
	return &Mega{
		addr: addr,
		srv:  srv,
	}, nil
}

func (m *Mega) Run(ctx context.Context) error {
	lis, err := net.Listen("tcp", m.addr)
	if err != nil {
		return err
	}
	// TODO: make use of context for graceful shutdown
	return m.srv.Serve(lis)
}
