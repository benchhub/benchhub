package server

import (
	"context"
	"fmt"
	"net/http"

	igrpc "github.com/at15/go.ice/ice/transport/grpc"
	ihttp "github.com/at15/go.ice/ice/transport/http"
	"github.com/dyweb/gommon/errors"
	dlog "github.com/dyweb/gommon/log"

	pbc "github.com/benchhub/benchhub/pkg/bhpb"
	"github.com/benchhub/benchhub/pkg/config"
)

// HttpServer is mainly used to communicate with browser, routes are mounted in transport http package
type HttpServer struct {
	registry     *Registry
	globalConfig config.AgentServerConfig
	log          *dlog.Logger
}

func NewHttpServer(r *Registry) (*HttpServer, error) {
	s := &HttpServer{
		registry:     r,
		globalConfig: r.Config,
	}
	dlog.NewStructLogger(log, s)
	return s, nil
}

func (srv *HttpServer) Ping(ctx context.Context, ping *pbc.Ping) (*pbc.Pong, error) {
	res := fmt.Sprintf("pong from agent %s your message is %s", igrpc.Hostname(), ping.Message)
	return &pbc.Pong{Message: res}, nil

}

func (srv *HttpServer) NodeInfo(ctx context.Context) (*pbc.NodeInfoRes, error) {
	node, err := NodeInfo(srv.globalConfig)
	if err != nil {
		log.Warnf("failed to get central node info %v", err)
		return nil, err
	}
	return &pbc.NodeInfoRes{
		Node: node,
	}, nil
}

func (srv *HttpServer) Handler() http.Handler {
	mux := http.NewServeMux()
	jMux := ihttp.NewJsonHandlerMux()
	srv.RegisterHandler(jMux)
	jMux.MountToStd(mux)
	return mux
}

func (srv *HttpServer) RegisterHandler(mux *ihttp.JsonHandlerMux) {
	mux.AddHandlerFunc("/api/ping", func() interface{} {
		return &pbc.Ping{}
	}, func(ctx context.Context, req interface{}) (res interface{}, err error) {
		if ping, ok := req.(*pbc.Ping); !ok {
			return nil, errors.New("invalid type, can't cast to *pbc.Ping")
		} else {
			return srv.Ping(ctx, ping)
		}
	})
	mux.AddHandlerFunc("/api/node", nil, func(ctx context.Context, req interface{}) (res interface{}, err error) {
		return srv.NodeInfo(ctx)
	})
}
