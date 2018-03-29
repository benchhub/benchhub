package server

import (
	"context"
	"fmt"
	"net/http"
	"os"

	ihttp "github.com/at15/go.ice/ice/transport/http"
	"github.com/dyweb/gommon/errors"
	dlog "github.com/dyweb/gommon/log"

	pb "github.com/benchhub/benchhub/pkg/bhpb"
	"github.com/benchhub/benchhub/pkg/central/store/meta"
	"github.com/benchhub/benchhub/pkg/config"
)

type HttpServer struct {
	meta         meta.Provider
	registry     *Registry
	globalConfig config.CentralServerConfig
	log          *dlog.Logger
}

func NewHttpServer(meta meta.Provider, r *Registry) (*HttpServer, error) {
	s := &HttpServer{
		meta:         meta,
		registry:     r,
		globalConfig: r.Config,
	}
	dlog.NewStructLogger(log, s)
	return s, nil
}

func (srv *HttpServer) Ping(ctx context.Context, ping *pb.Ping) (*pb.Pong, error) {
	if host, err := os.Hostname(); err != nil {
		return &pb.Pong{Message: "pong from unknown"}, errors.Wrap(err, "can't get hostname")
	} else {
		res := fmt.Sprintf("pong from central %s your message is %s", host, ping.Message)
		return &pb.Pong{Message: res}, nil
	}
}

func (srv *HttpServer) NodeInfo(ctx context.Context) (*pb.NodeInfoRes, error) {
	node, err := NodeInfo(srv.globalConfig)
	if err != nil {
		log.Warnf("failed to get central node info %v", err)
		return nil, err
	}
	return &pb.NodeInfoRes{
		Node: node,
	}, nil
}

func (srv *HttpServer) ListAgent(ctx context.Context) (*pb.ListAgentRes, error) {
	node, err := srv.meta.ListNodes()
	if err != nil {
		return nil, err
	}
	return &pb.ListAgentRes{
		Agents: node,
	}, nil
}

func (srv *HttpServer) Handler() http.Handler {
	mux := http.NewServeMux()
	jMux := ihttp.NewJsonHandlerMux()
	srv.RegisterHandler(jMux)
	jMux.MountToStd(mux)
	// FIXME: need to figure out a way to mount api in both / and /api
	//mux.Handle("/api/", http.StripPrefix("/api/", mux))
	return mux
}

func (srv *HttpServer) RegisterHandler(mux *ihttp.JsonHandlerMux) {
	mux.AddHandlerFunc("/api/ping", func() interface{} {
		return &pb.Ping{}
	}, func(ctx context.Context, req interface{}) (res interface{}, err error) {
		if ping, ok := req.(*pb.Ping); !ok {
			return nil, errors.New("invalid type, can't cast to *pb.Ping")
		} else {
			return srv.Ping(ctx, ping)
		}
	})
	mux.AddHandlerFunc("/api/node", nil, func(ctx context.Context, req interface{}) (res interface{}, err error) {
		return srv.NodeInfo(ctx)
	})
	mux.AddHandlerFunc("/api/agent/list", nil, func(ctx context.Context, req interface{}) (res interface{}, err error) {
		return srv.ListAgent(ctx)
	})
}
